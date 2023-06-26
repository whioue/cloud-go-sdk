package utils

import "gitlab.inspir.work/HanYiming/cloud-go-sdk/pkg/json"

type MapStringInterface map[string]string

func ConvertStruct2Map(v interface{}) ([]byte, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
