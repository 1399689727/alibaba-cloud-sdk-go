package r_kvstore

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyDBInstanceMonitor invokes the r_kvstore.ModifyDBInstanceMonitor API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifydbinstancemonitor.html
func (client *Client) ModifyDBInstanceMonitor(request *ModifyDBInstanceMonitorRequest) (response *ModifyDBInstanceMonitorResponse, err error) {
	response = CreateModifyDBInstanceMonitorResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceMonitorWithChan invokes the r_kvstore.ModifyDBInstanceMonitor API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifydbinstancemonitor.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceMonitorWithChan(request *ModifyDBInstanceMonitorRequest) (<-chan *ModifyDBInstanceMonitorResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceMonitorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceMonitor(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyDBInstanceMonitorWithCallback invokes the r_kvstore.ModifyDBInstanceMonitor API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifydbinstancemonitor.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceMonitorWithCallback(request *ModifyDBInstanceMonitorRequest, callback func(response *ModifyDBInstanceMonitorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceMonitorResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceMonitor(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyDBInstanceMonitorRequest is the request struct for api ModifyDBInstanceMonitor
type ModifyDBInstanceMonitorRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Interval             string           `position:"Query" name:"Interval"`
}

// ModifyDBInstanceMonitorResponse is the response struct for api ModifyDBInstanceMonitor
type ModifyDBInstanceMonitorResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBInstanceMonitorRequest creates a request to invoke ModifyDBInstanceMonitor API
func CreateModifyDBInstanceMonitorRequest() (request *ModifyDBInstanceMonitorRequest) {
	request = &ModifyDBInstanceMonitorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyDBInstanceMonitor", "", "")
	return
}

// CreateModifyDBInstanceMonitorResponse creates a response to parse from ModifyDBInstanceMonitor response
func CreateModifyDBInstanceMonitorResponse() (response *ModifyDBInstanceMonitorResponse) {
	response = &ModifyDBInstanceMonitorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
