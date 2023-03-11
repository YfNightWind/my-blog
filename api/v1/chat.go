package v1

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"my-blog/utils"
	"my-blog/utils/errormsg"
	"net/http"
)

func Chat(ctx *gin.Context) {
	// 设置 API 访问地址
	url := "https://api.openai.com/v1/chat/completions"

	question, _ := ctx.GetPostForm("question")
	message := []interface{}{
		map[string]interface{}{
			"role":    "user",
			"content": question,
		},
	}

	// 设置 POST 请求的请求体
	requestData := map[string]interface{}{
		"model":    "gpt-3.5-turbo",
		"messages": message,
		"n":        1,
		"stop":     "",
	}
	requestBody, _ := json.Marshal(requestData)
	body := bytes.NewBuffer(requestBody)

	// 创建 POST 请求
	req, _ := http.NewRequest("POST", url, body)

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+utils.ApiKey) // 设置您的ApiKey

	// 发送请求
	client := &http.Client{}
	resp, _ := client.Do(req)

	// 解析响应
	responseBody, _ := io.ReadAll(resp.Body)
	var responseData map[string]interface{}
	err := json.Unmarshal(responseBody, &responseData)
	if err != nil {
		return
	}

	// 响应结果 案例
	//{
	//	"id": "chatcmpl-123",
	//	"object": "chat.completion",
	//	"created": 1677652288,
	//	"choices": [{
	//		"index": 0,
	//		"message": {
	//			"role": "assistant",
	//			"content": "\n\nHello there, how may I assist you today?",
	//		},
	//		"finish_reason": "stop"
	//	}],
	//	"usage": {
	//		"prompt_tokens": 9,
	//		"completion_tokens": 12,
	//		"total_tokens": 21
	//	}
	//}

	// 获取 "content" 值
	choices, ok := responseData["choices"].([]interface{})
	if !ok {
		return
	}
	choice, ok := choices[0].(map[string]interface{})
	if !ok {
		return
	}
	messages, ok := choice["message"].(map[string]interface{})
	if !ok {
		return
	}
	content, ok := messages["content"].(string)
	if !ok {
		return
	}

	// 输出响应结果
	ctx.JSON(http.StatusOK, gin.H{
		"code": errormsg.SUCCESS,
		"data": content,
		"msg":  errormsg.GetErrorMsg(errormsg.SUCCESS),
	})
}
