package juniperUDP

import (
	"os"
	"log"
	"fmt"
	"time"
	"reflect"
	"github.com/golang/protobuf/jsonpb"
	"encoding/json"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/metric"
	"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/telemetry_top"
	_"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/qmon"  		// blank import as it is only used to Unmarshal proto message
	_"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/lsp_mon"
	_"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/cmerror_data"
	_"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/cmerror"
	_"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/cpu_memory_utilization"
	_"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/fabric"
	_"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/firewall"
	_"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/inline_jflow"
	_"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/logical_port"
	_"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/lsp_stats"
	_"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/npu_memory_utilization"
	_"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/npu_utilization"
	_"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/optics"
	_"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/packet_stats"
	_"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/port_exp"
	_"github.com/influxdata/telegraf/plugins/parsers/juniperUDP/port"
	"github.com/golang/protobuf/proto"
	"strings"
)

// JuniperUDPParser is an object for Parsing incoming metrics.
type JuniperUDPParser struct {
	// DefaultTags will be added to every parsed metric
	DefaultTags map[string]string
}

/*
func (p *JuniperUDPParser) ParseWithDefaultTimePrecision(buf []byte, t time.Time, precision string) ([]telegraf.Metric, error) {
	if !bytes.HasSuffix(buf, []byte("\n")) {
		buf = append(buf, '\n')
	}
	// parse even if the buffer begins with a newline
	buf = bytes.TrimPrefix(buf, []byte("\n"))
	metrics, err := metric.ParseWithDefaultTimePrecision(buf, t, precision)
	if len(p.DefaultTags) > 0 {
		for _, m := range metrics {
			for k, v := range p.DefaultTags {
				// only set the default tag if it doesn't already exist:
				if !m.HasTag(k) {
					m.AddTag(k, v)
				}
			}
		}
	}
	return metrics, err
}
*/

func parseArray(data []interface{}, masterKey string) []interface{} {
        var arrData []interface{}
        for _,val := range data{
                valType := reflect.ValueOf(val).Kind()
                if valType == reflect.Map{
			mapData := parseMap(val.(map[string]interface{}), masterKey)
			for _,tmpData := range(mapData){
				arrData = append(arrData, tmpData)
			}
                } else {fmt.Println("Error!!!! Leaf elements in array are not coded. Please open a issue.")} 
        }

        return arrData
}


func parseMap(data map[string]interface{}, masterKey string) []interface{} {
        var leafData map[string]interface{}
        var arrData []interface{}
        var arrKey []string
        var finalData []interface{}
	var newMasterKey string
	leafData = make(map[string]interface{})
	
        for key,val := range data{
                if masterKey == ""{
                        newMasterKey = key
                } else {
                        newMasterKey = masterKey + "." + key
                }

                valType := reflect.ValueOf(val).Kind()
                if valType == reflect.Map{
                        mapData := parseMap(val.(map[string]interface{}), newMasterKey)
       			if reflect.TypeOf(mapData).Kind() == reflect.Map{
				var tmpArr []interface{}
				tmpArr = append(tmpArr, mapData)
				arrData = append(arrData, tmpArr)
			} else if reflect.TypeOf(mapData).Kind() == reflect.Slice{
				arrData = append(arrData, mapData)
			}
                        arrKey = append(arrKey, newMasterKey)
                } else if valType == reflect.Slice{
                        arrData = append(arrData, parseArray(val.([]interface{}), newMasterKey))
                        arrKey = append(arrKey, newMasterKey)

                } else { leafData[newMasterKey] = val}
	}
	if len(leafData) != 0 {
	        for i,key := range arrKey{
			_ = key
			for _,data_aa := range arrData[i].([]interface{}){
                                leafTmp := leafData
				if data_aa != nil {
                                        if reflect.ValueOf(data_aa).Kind() == reflect.Map{
                                                for key_aa, value_aa := range data_aa.(map[string]interface{}){
                                                        leafTmp[key_aa] = value_aa
                                                }
				                finalData = append(finalData, make(map[string]interface{}))
						for k,v := range leafTmp{
                                                                finalData[len(finalData)-1].(map[string]interface{})[k] = v
                                                }	
                                        } else {
                                                 for _,data_ha := range data_aa.([]interface{}){
                                                        leafTmp = leafData
							for key_aa,value_aa := range data_ha.(map[string]interface{}){
                                                        	leafTmp[key_aa] = value_aa
                                                        }
							finalData = append(finalData, make(map[string]interface{}))
							for k,v := range leafTmp{
								finalData[len(finalData)-1].(map[string]interface{})[k] = v
							}	
						}
                                        }
                                }
                        }
	       }
        } else {finalData = arrData}
	arrData = arrData[:0]
        if (len(finalData) == 0) && (len(leafData)!= 0) {
	        finalData = append(finalData, leafData)
        }
        return finalData
}


// Parse returns a slice of Metrics from a text representation of a
// metric (in line-protocol format)
// with each metric separated by newlines. If any metrics fail to parse,
// a non-nil error will be returned in addition to the metrics that parsed
// successfully.
func (p *JuniperUDPParser) Parse(buf []byte) ([]telegraf.Metric, error) {
	ts := &telemetry_top.TelemetryStream{}
	if err := proto.Unmarshal(buf, ts); err != nil {
        	log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Println(ts)	
	host,errHost := os.Hostname()
	if errHost != nil {
		fmt.Println("Error!! Host name not found")
	}
	deviceName := ts.GetSystemId()
	gpbTime := ts.GetTimestamp()
	measurementPrefix := "enterprise.juniperNetworks"
	jnprSensorName := ts.GetSensorName()
	tmpSlice := strings.Split(jnprSensorName, ":")
	sensorName := tmpSlice[1]
	_ = gpbTime
	_ = measurementPrefix
	_ = jnprSensorName
	m := &jsonpb.Marshaler{}
	tsJSON,err := m.MarshalToString(ts)	
	if err!= nil{
		fmt.Println("Error")
	}
	var data map[string]interface{}
        errU := json.Unmarshal([]byte(tsJSON), &data)
        if errU != nil {
                panic(errU)	
	}
	enterpriseSensorData := data["enterprise"]
	sensorData, ok := enterpriseSensorData.(map[string]interface{})
	jnprSensorData := sensorData["[juniperNetworks]"]
	if !ok {
    		panic("inner map is not a map!")
	}
	metrics := make([]telegraf.Metric, 0)
	sensorNum := 0
	for _,sensorData := range jnprSensorData.(map[string]interface{}){
		var fields map[string]interface{}
		if reflect.ValueOf(sensorData).Kind() == reflect.Map {
			_ = sensorName
			parsedData := parseMap(sensorData.(map[string]interface{}), "")
			for _,finalData := range(parsedData){
				sequenceNum := 0
				for _,fin := range(finalData.([]interface{})){
					//fin = fin.(map[string] interface{})
					fin.(map[string]interface{})["device"] = deviceName
					fin.(map[string]interface{})["host"] = host
					fin.(map[string]interface{})["sensor_name"] = sensorName
					fin.(map[string]interface{})["_seq"] = sequenceNum
					fields = fin.(map[string]interface{})
					tags := make(map[string]string)
					for k, v := range p.DefaultTags {
						tags[k] = v
					}
					mtrc,err := metric.New(sensorName, tags, fields,time.Now().UTC())
					metrics = append(metrics, mtrc)
					if err!=nil {fmt.Println(err)}
					sensorNum++
					sequenceNum++
				}
			}		
		}
	}		
//	fmt.Printf("\nData (JSON) = \n%s\n", data)
//	fmt.Println("\nJuniper Sensor Name: \n%s\n", jnprSensorName)
//	fmt.Println("\nDevice name: \n%s\n", deviceName)
//	fmt.Println("\nGPB time: \n%s\n", gpbTime)
//	fmt.Println(measurementPrefix)
//	fmt.Println("\nMetrics: \n")
//	fmt.Println(metrics)
	return metrics, err	
//	return p.ParseWithDefaultTimePrecision(buf, time.Now(), "")
}



func (p *JuniperUDPParser) ParseLine(line string) (telegraf.Metric, error) {
	metrics, err := p.Parse([]byte(line + "\n"))

	if err != nil {
		return nil, err
	}

	if len(metrics) < 1 {
		return nil, fmt.Errorf(
			"Can not parse the line: %s, for data format: influx ", line)
	}

	return metrics[0], nil
}

func (p *JuniperUDPParser) SetDefaultTags(tags map[string]string) {
	p.DefaultTags = tags
}

