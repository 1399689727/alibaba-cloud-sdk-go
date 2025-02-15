package multimediaai

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

// CreateGifTask invokes the multimediaai.CreateGifTask API synchronously
// api document: https://help.aliyun.com/api/multimediaai/creategiftask.html
func (client *Client) CreateGifTask(request *CreateGifTaskRequest) (response *CreateGifTaskResponse, err error) {
	response = CreateCreateGifTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateGifTaskWithChan invokes the multimediaai.CreateGifTask API asynchronously
// api document: https://help.aliyun.com/api/multimediaai/creategiftask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGifTaskWithChan(request *CreateGifTaskRequest) (<-chan *CreateGifTaskResponse, <-chan error) {
	responseChan := make(chan *CreateGifTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateGifTask(request)
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

// CreateGifTaskWithCallback invokes the multimediaai.CreateGifTask API asynchronously
// api document: https://help.aliyun.com/api/multimediaai/creategiftask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGifTaskWithCallback(request *CreateGifTaskRequest, callback func(response *CreateGifTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateGifTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateGifTask(request)
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

// CreateGifTaskRequest is the request struct for api CreateGifTask
type CreateGifTaskRequest struct {
	*requests.RpcRequest
	VideoUrl      string `position:"Query" name:"VideoUrl"`
	VideoName     string `position:"Query" name:"VideoName"`
	ApplicationId string `position:"Query" name:"ApplicationId"`
}

// CreateGifTaskResponse is the response struct for api CreateGifTask
type CreateGifTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    int64  `json:"TaskId" xml:"TaskId"`
}

// CreateCreateGifTaskRequest creates a request to invoke CreateGifTask API
func CreateCreateGifTaskRequest() (request *CreateGifTaskRequest) {
	request = &CreateGifTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("multimediaai", "2019-08-10", "CreateGifTask", "multimediaai", "openAPI")
	return
}

// CreateCreateGifTaskResponse creates a response to parse from CreateGifTask response
func CreateCreateGifTaskResponse() (response *CreateGifTaskResponse) {
	response = &CreateGifTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
