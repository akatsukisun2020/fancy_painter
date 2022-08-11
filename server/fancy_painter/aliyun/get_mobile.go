package aliyun

import (
	"log"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dypnsapi20170525 "github.com/alibabacloud-go/dypnsapi-20170525/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

// 参考：  https://next.api.aliyun.com/api/Dypnsapi/2017-05-25/GetMobile?lang=GO&params={}&accounttraceid=af2afc363891431e8763974be052b8cccrvu

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dypnsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dypnsapi.aliyuncs.com")
	_result = &dypnsapi20170525.Client{}
	_result, _err = dypnsapi20170525.NewClient(config)
	return _result, _err
}

// 通过token获取号码，返回值参考： https://help.aliyun.com/document_detail/85198.htm?spm=a2c4g.11186623.0.0.5260698dIZ0TpZ
func GetMobile(accessKeyId, accessKeySecret *string, accessToken *string) (_mobile string, _err error) {
	client, _err := CreateClient(accessKeyId, accessKeySecret)
	if _err != nil {
		return "", _err
	}

	getMobileRequest := &dypnsapi20170525.GetMobileRequest{
		AccessToken: accessToken,
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值

		resp, _errTmp := client.GetMobileWithOptions(getMobileRequest, runtime)
		if _errTmp != nil || *resp.Body.Code != "OK" {
			log.Printf("GetMobileWithOptions error, resp:%v, _err:%v", resp, _errTmp)
			_err = _errTmp
			return _errTmp
		}
		_mobile = *resp.Body.GetMobileResultDTO.Mobile
		log.Printf("GetMobileWithOptions resp:%v", resp)

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		util.AssertAsString(error.Message)
	}
	return _mobile, _err
}
