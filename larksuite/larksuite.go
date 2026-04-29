package larksuite

import (
	"context"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func SendString(appId, appSecret, msg, receiveId string) {
	// 创建 Client
	client := lark.NewClient(appId, appSecret)
	// 创建请求对象

	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(`user_id`).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(receiveId).
			MsgType(`text`).
			Content(fmt.Sprintf(`{"text":"%s"}`, msg)).
			Build()).
		Build()

	// 发起请求
	resp, err := client.Im.Message.Create(context.Background(), req)

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Printf("logId: %s, error response: \n%s", resp.RequestId(), larkcore.Prettify(resp.CodeError))
		return
	}

	// 业务处理
	//fmt.Println(larkcore.Prettify(resp))
}
