package myconverts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type MapDynamic = map[string]interface{}

// convert struct to map[string]interface{}
func StructToMap(obj interface{}) (MapDynamic, error) {
	var result MapDynamic

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// convert struct to []byte{}
func StructToByte(obj interface{}) ([]byte, error) {

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("Marshal StructToByte.Err : ", err)

		return nil, err
	}

	return jsonBytes, err
}

// convert struct + id (string) to []byte{}
func StructAndIdToByte(obj interface{}, id string) ([]byte, error) {
	/// ini untuk put biasanya
	/// pergunakan sesuai kebutuhan ges
	mapS := make(MapDynamic)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("Marshal StructToByte.Err : ", err)

		return nil, err
	}
	merr := json.Unmarshal(jsonBytes, &mapS)
	if merr != nil {
		return nil, merr
	}

	mapS["id"] = id
	finalBytes, err := json.Marshal(mapS)

	return finalBytes, err
}

// convert map[string]interface{} to []byte{}
func MapToByte(theMap MapDynamic) ([]byte, error) {
	finalBytes, err := json.Marshal(theMap)
	if err != nil {
		fmt.Println("Marshal MapToByte Err : ", err)
	}

	return finalBytes, err
}

// convert url.values to map[string]interface{}
func UrlValuesToMap(urlV url.Values) MapDynamic {
	resultMap := make(map[string]interface{})
	for key, vals := range urlV {
		if len(vals) == 1 {
			resultMap[key] = vals[0]
		} else {
			var interfaceSlice []interface{}
			for _, v := range vals {
				interfaceSlice = append(interfaceSlice, v)
			}
			resultMap[key] = interfaceSlice
		}
	}
	return resultMap
}

// convert http.Response to map[string]interface
func BodyResToMap(resp *http.Response) (MapDynamic, error) {
	defer resp.Body.Close()
	mapReturn := make(MapDynamic)
	parsedBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("io.ReadAll err : ", err)

		return nil, err
	}
	mErr := json.Unmarshal(parsedBody, &mapReturn)
	if mErr != nil {
		fmt.Println("json.Unmarsal err : ", err)
	}
	return mapReturn, mErr
}
