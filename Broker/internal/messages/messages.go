package messages

const CreateTopicAction string = "CREATE_TOPIC"
const PublishMessageAction string = "PUBLISH_MESSAGE"
const SubscribeAction string = "SUBSCRIBE"

type SenderRequest struct {
	Action    string `json:"action"`
	TopicName string `json:"topic_name"`
	Message   string `json:"message"`
}

type SenderResponse struct {
	Action  string `json:"action"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
