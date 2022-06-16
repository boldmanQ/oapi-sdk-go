package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/feishu/oapi-sdk-go"
	"github.com/feishu/oapi-sdk-go/core"
	"github.com/feishu/oapi-sdk-go/service/im/v1"
)

func uploadImage(client *client.Client) {
	pdf, err := os.Open("/Users/bytedance/Downloads/a.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pdf.Close()

	resp, err := client.Im.Images.Create(context.Background(),
		im.NewCreateImageReqBuilder().
			Body(im.NewCreateImageReqBodyBuilder().
				ImageType(im.IMAGE_TYPE_MESSAGE).
				Image(pdf).
				Build()).
			Build())

	if err != nil {
		fmt.Println(core.Prettify(err))
		return
	}
	fmt.Println(core.Prettify(resp))
	fmt.Println(resp.RequestId())

}

func uploadFile(client *client.Client) {
	pdf, err := os.Open("/Users/bytedance/Downloads/redis.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pdf.Close()

	resp, err := client.Im.Files.Create(context.Background(),
		im.NewCreateFileReqBuilder().
			Body(im.NewCreateFileReqBodyBuilder().
				FileType(im.FILE_TYPE_PDF).
				FileName("redis.pdf").
				File(pdf).
				Build()).
			Build())

	if err != nil {
		fmt.Println(core.Prettify(err))
		return
	}
	fmt.Println(core.Prettify(resp))
	fmt.Println(resp.RequestId())

}

func uploadImage2(client *client.Client) {
	body, err := im.NewCreateImagePathReqBodyBuilder().ImagePath("/Users/bytedance/Downloads/a.jpg").ImageType(im.IMAGE_TYPE_MESSAGE).Build()
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Im.Images.Create(context.Background(), im.NewCreateImageReqBuilder().Body(body).Build())

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(core.Prettify(resp))
	fmt.Println(resp.RequestId())

}

func downLoadImage(client *client.Client) {
	resp, err := client.Im.Images.Get(context.Background(), im.NewGetImageReqBuilder().ImageKey("img_v2_cd2657c7-ad1e-410a-8e76-942c89203bfg").Build())

	if err != nil {
		fmt.Println(core.Prettify(err))
		return
	}

	if resp.Code != 0 {
		fmt.Println(core.Prettify(resp))
		return
	}
	fmt.Println(resp.FileName)
	fmt.Println(resp.RequestId())

	bs, err := ioutil.ReadAll(resp.File)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile("test_download_v2.jpg", bs, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func downLoadImageV2(client *client.Client) {
	resp, err := client.Im.Images.Get(context.Background(), im.NewGetImageReqBuilder().ImageKey("img_v2_cd2657c7-ad1e-410a-8e76-942c89203bfg").Build())

	if err != nil {
		fmt.Println(core.Prettify(err))
		return
	}

	if resp.Code != 0 {
		fmt.Println(core.Prettify(resp))
		return
	}
	fmt.Println(resp.FileName)
	fmt.Println(resp.RequestId())

	resp.WriteFile("a.jpg")

}

func sendTextMsg(client *client.Client) {
	content, err := im.NewTextMsgBuilder().
		Text("加多").
		Line().
		TextLine("hello").
		TextLine("world").
		AtUser("ou_c245b0a7dff2725cfa2fb104f8b48b9d", "陆续").
		//AtAll().
		Build()

	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Im.Messages.Create(context.Background(), im.NewCreateMessageReqBuilder().
		ReceiveIdType(im.RECEIVE_ID_TYPE_OPEN_ID).
		Body(im.NewCreateMessageReqBodyBuilder().
			MsgType(im.MSG_TYPE_TEXT).
			ReceiveId("ou_c245b0a7dff2725cfa2fb104f8b48b9d").
			Content(content).
			Build()).
		Build())

	if err != nil {
		fmt.Println(core.Prettify(err))
		return
	}
	fmt.Println(core.Prettify(resp))
	fmt.Println(resp.RequestId())

}

func sendImageMsg(client *client.Client) {
	msgImage := im.MessageImage{ImageKey: "img_v2_0db0c471-fff0-460a-883f-e3523c478c4g"}
	content, err := msgImage.String()
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Im.Messages.Create(context.Background(), im.NewCreateMessageReqBuilder().
		ReceiveIdType(im.RECEIVE_ID_TYPE_OPEN_ID).
		Body(im.NewCreateMessageReqBodyBuilder().
			MsgType(im.MSG_TYPE_IMAGE).
			ReceiveId("ou_c245b0a7dff2725cfa2fb104f8b48b9d").
			Content(content).
			Build()).
		Build())

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(core.Prettify(resp))
	fmt.Println(resp.RequestId())

}

func sendShardChatMsg(client *client.Client) {
	msgImage := im.MessageShareChat{ChatId: "oc_4ea14cc15e39ef80a579ca74895a33ff"}
	content, err := msgImage.String()
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Im.Messages.Create(context.Background(), im.NewCreateMessageReqBuilder().
		ReceiveIdType(im.RECEIVE_ID_TYPE_OPEN_ID).
		Body(im.NewCreateMessageReqBodyBuilder().
			MsgType(im.MSG_TYPE_SHARE_CHAT).
			ReceiveId("ou_c245b0a7dff2725cfa2fb104f8b48b9d").
			Content(content).
			Build()).
		Build())

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(core.Prettify(resp))
	fmt.Println(resp.RequestId())

}

func sendShardUserMsg(client *client.Client) {
	msgImage := im.MessageShareUser{UserId: "ou_487f709a942d16edafe57fd6fbc4bcf5"}
	content, err := msgImage.String()
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Im.Messages.Create(context.Background(), im.NewCreateMessageReqBuilder().
		ReceiveIdType(im.RECEIVE_ID_TYPE_OPEN_ID).
		Body(im.NewCreateMessageReqBodyBuilder().
			MsgType(im.MSG_TYPE_SHARE_USER).
			ReceiveId("ou_c245b0a7dff2725cfa2fb104f8b48b9d").
			Content(content).
			Build()).
		Build())

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(core.Prettify(resp))
	fmt.Println(resp.RequestId())

}

func sendAudioMsg(client *client.Client) {
	msgImage := im.MessageAudio{FileKey: "file_v2_0c7f5b4b-64ec-408a-a9eb-09aec7954a4g"}
	content, err := msgImage.String()
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Im.Messages.Create(context.Background(), im.NewCreateMessageReqBuilder().
		ReceiveIdType(im.RECEIVE_ID_TYPE_OPEN_ID).
		Body(im.NewCreateMessageReqBodyBuilder().
			MsgType(im.MSG_TYPE_AUDIO).
			ReceiveId("ou_c245b0a7dff2725cfa2fb104f8b48b9d").
			Content(content).
			Build()).
		Build())

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(core.Prettify(resp))
	fmt.Println(resp.RequestId())

}

func sendMediaMsg(client *client.Client) {
	msgImage := im.MessageMedia{FileKey: "file_v2_0c7f5b4b-64ec-408a-a9eb-09aec7954a4g", ImageKey: "ssss"}
	content, err := msgImage.String()
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Im.Messages.Create(context.Background(), im.NewCreateMessageReqBuilder().
		ReceiveIdType(im.RECEIVE_ID_TYPE_OPEN_ID).
		Body(im.NewCreateMessageReqBodyBuilder().
			MsgType(im.MSG_TYPE_MEDIA).
			ReceiveId("ou_c245b0a7dff2725cfa2fb104f8b48b9d").
			Content(content).
			Build()).
		Build())

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(core.Prettify(resp))
	fmt.Println(resp.RequestId())

}

func sendFileMsg(client *client.Client) {
	msgImage := im.MessageFile{FileKey: "file_v2_0c7f5b4b-64ec-408a-a9eb-09aec7954a4g"}
	content, err := msgImage.String()
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Im.Messages.Create(context.Background(), im.NewCreateMessageReqBuilder().
		ReceiveIdType(im.RECEIVE_ID_TYPE_OPEN_ID).
		Body(im.NewCreateMessageReqBodyBuilder().
			MsgType(im.MSG_TYPE_FILE).
			ReceiveId("ou_c245b0a7dff2725cfa2fb104f8b48b9d").
			Content(content).
			Build()).
		Build())

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(core.Prettify(resp))
	fmt.Println(resp.RequestId())

}

func sendStickerMsg(client *client.Client) {
	msgImage := im.MessageSticker{FileKey: "img_v2_2a0372ea-dc03-44d7-b052-255bede4d42g"}
	content, err := msgImage.String()
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Im.Messages.Create(context.Background(), im.NewCreateMessageReqBuilder().
		ReceiveIdType(im.RECEIVE_ID_TYPE_CHAT_ID).
		Body(im.NewCreateMessageReqBodyBuilder().
			MsgType(im.MSG_TYPE_STICKER).
			ReceiveId("121212").
			Content(content).
			Build()).
		Build())

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(core.Prettify(resp))
	fmt.Println(resp.RequestId())

}

func sendPostMsg(client *client.Client) {
	// 2.1 创建text与href元素
	zhCnPostText := &im.MessagePostText{Text: "中文内容", UnEscape: false}
	zhCnPostA := &im.MessagePostA{Text: "test content", Href: "http://www.baidu.com", UnEscape: false}
	enUsPostText := &im.MessagePostText{Text: "英文内容", UnEscape: false}
	enUsPostA := &im.MessagePostA{Text: "test content", Href: "http://www.baidu.com", UnEscape: false}

	// 2.2 构建消息content
	zhCnMessagePostContent := &im.MessagePostContent{Title: "title1", Content: [][]im.MessagePostElement{{zhCnPostText, zhCnPostA}}}
	enUsMessagePostContent := &im.MessagePostContent{Title: "title2", Content: [][]im.MessagePostElement{{enUsPostText, enUsPostA}}}
	messagePostText := &im.MessagePost{ZhCN: zhCnMessagePostContent, EnUS: enUsMessagePostContent}
	content, err := messagePostText.String()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Im.Messages.Create(context.Background(), im.NewCreateMessageReqBuilder().
		ReceiveIdType(im.RECEIVE_ID_TYPE_OPEN_ID).
		Body(im.NewCreateMessageReqBodyBuilder().
			MsgType(im.MSG_TYPE_POST).
			ReceiveId("ou_c245b0a7dff2725cfa2fb104f8b48b9d").
			Content(content).
			Build()).
		Build(), core.WithRequestId("jiaduo_id"))

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(core.Prettify(resp))
	fmt.Println(resp.RequestId())
}

func sendPostMsgUseBuilder(client *client.Client) {
	zhCnPostText := &im.MessagePostText{Text: "中文内容", UnEscape: false}
	zhCnPostA := &im.MessagePostA{Text: "test content", Href: "http://www.baidu.com", UnEscape: false}
	enUsPostText := &im.MessagePostText{Text: "英文内容", UnEscape: false}
	enUsPostA := &im.MessagePostA{Text: "test content", Href: "http://www.baidu.com", UnEscape: false}

	zhCn := im.NewMessagePostContent().
		ContentTitle("title1").
		AppendContent([]im.MessagePostElement{zhCnPostText, zhCnPostA}).
		Build()

	enUs := im.NewMessagePostContent().
		ContentTitle("title1").
		AppendContent([]im.MessagePostElement{enUsPostA, enUsPostText}).
		Build()
	postText, err := im.NewMessagePost().ZhCn(zhCn).EnUs(enUs).Build()
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Im.Messages.Create(context.Background(), im.NewCreateMessageReqBuilder().
		ReceiveIdType(im.RECEIVE_ID_TYPE_OPEN_ID).
		Body(im.NewCreateMessageReqBodyBuilder().
			MsgType(im.MSG_TYPE_POST).
			ReceiveId("ou_c245b0a7dff2725cfa2fb104f8b48b9d").
			Content(postText).
			Build()).
		Build())

	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.Success() {
		fmt.Println(core.Prettify(resp))
	} else {
		fmt.Println(resp.RequestId(), resp.Msg, resp.Code)
	}

}
func main() {
	var appID, appSecret = os.Getenv("APP_ID"), os.Getenv("APP_SECRET")
	var feishu_client = client.NewClient(appID, appSecret, client.WithLogLevel(core.LogLevelDebug))
	//downLoadImageV2(client)
	//uploadImage(client)
	//uploadImage(client)
	//downLoadImage(client)
	//uploadImage2(client)
	//sendTextMsg(client)
	//sendImageMsg(client)
	//uploadFile(client)
	//sendFileMsg(client)
	//sendAudioMsg(client)
	//sendMediaMsg(client)
	//sendShardChatMsg(client)
	//sendShardUserMsg(client)
	sendPostMsg(feishu_client)
	//sendPostMsgUseBuilder(client)
}
