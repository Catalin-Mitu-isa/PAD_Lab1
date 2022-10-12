package messages

const CreateTopicAction string = "CREATE_TOPIC"
const PublishMessageAction string = "PUBLISH_MESSAGE"
const SubscribeAction string = "SUBSCRIBE"

type Data struct {
	Id uint `json:"id"`
}

type SenderRequest struct {
	Action    string `json:"action"`
	TopicName string `json:"topic_name"`
	Message   string `json:"message"`
}

type SenderResponse struct {
	Action  string `json:"action"`
	Success bool   `json:"success"`
	Data    Data   `json:"data"`
	Error   string `json:"error"`
}
