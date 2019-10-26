package libs

import (
	"fmt"

	configs "catdogs-be/configs/common"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

const (
	method   = "POST"
	scheme   = "https"
	domain   = "dysmsapi.aliyuncs.com"
	version  = "2017-05-25"
	apiName  = "SendSms"
	regionId = "cn-hangzhou"
)

func SendSms(phone, code string) error {
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", configs.C.AccessKeyId, configs.C.AccessSecret)
	if err != nil {
		fmt.Println(err)
		return err
	}

	request := requests.NewCommonRequest()
	request.Method = method
	request.Scheme = scheme
	request.Domain = domain
	request.Version = version
	request.ApiName = apiName
	request.QueryParams["RegionId"] = regionId
	request.QueryParams["SignName"] = configs.C.SignName
	request.QueryParams["TemplateCode"] = configs.C.TemplateCode
	request.QueryParams["PhoneNumbers"] = phone
	request.QueryParams["TemplateParam"] = "{\"code\":\"" + code + "\"}"

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Print(response.GetHttpContentString())
	return nil
}
