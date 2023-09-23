package main

import (
	"fmt"
	"unsafe"

	"github.com/valyala/fastjson"
)

//export go_filter
func go_filter(tag *uint8, tag_len uint, time_sec uint, time_nsec uint, record *uint8, record_len uint) *uint8 {
	// btag := unsafe.Slice(tag, tag_len) // Note, requires Go 1.17 (tinygo 0.20)
	brecord := unsafe.Slice(record, record_len)
	// now := time.Unix(int64(time_sec), int64(time_nsec))

	br := string(brecord)

	var p fastjson.Parser
	value, err := p.Parse(br)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	obj, err := value.Object()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	kubernetesObj := obj.Get("kubernetes")
	if kubernetesObj == nil {
		s := obj.String()
		s += string(rune(0)) // Note: explicit null terminator.
		rv := []byte(s)

		return &rv[0]
	}

	labelsObj := kubernetesObj.GetObject("labels")
	if labelsObj == nil {
		s := obj.String()
		s += string(rune(0)) // Note: explicit null terminator.
		rv := []byte(s)

		return &rv[0]
	}

	var ar fastjson.Arena
	var region, tenant, hasSensitiveData, newTag string

	if labelsObj.Get("region") != nil {
		region = labelsObj.Get("region").String()
	}
	if labelsObj.Get("tenant") != nil {
		tenant = labelsObj.Get("tenant").String()
	}
	if labelsObj.Get("has-sensitive-data") != nil {
		hasSensitiveData = labelsObj.Get("has-sensitive-data").String()
	}

	if region == "\"ap-south-1\"" && tenant == "\"private-a1\"" {
		newTag = "private-a1-elastic-ap-south-1"
		if hasSensitiveData == "\"true\"" {
			newTag = "private-a1-s3-ap-south-1"
		}
	} else if region == "\"ap-south-1\"" {
		newTag = "general-elastic-ap-south-1"
		if hasSensitiveData == "\"true\"" {
			newTag = "general-s3-ap-south-1"
		}
	} else {
		newTag = "general-elastic"
		if hasSensitiveData == "\"true\"" {
			newTag = "general-s3"
		}
	}

	obj.Set("new_tag", ar.NewString(newTag))
	s := obj.String()
	s += string(rune(0)) // Note: explicit null terminator.
	rv := []byte(s)

	return &rv[0]
}

func main() {}
