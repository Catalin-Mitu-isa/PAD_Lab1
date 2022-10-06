package messages

const CREATE_TOPIC_ACTION string = "CREATE_TOPIC"
const PUBLISH_MESSAGE_ACTION string = "PUBLISH_MESSAGE"
const SUBSCRIBE_ACTION string = "SUBSCRIBE"

type SenderMessage struct {
	Action string `json:"action"`
}

type SenderCreateRequest struct {
	SenderMessage
	TopicName string `json:"topic_name"`
}

type SenderPublishRequest struct {
	SenderMessage
	Message string `json:"message"`
}

type SenderResponse struct {
	SenderMessage
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type SenderCreateResponse struct {
	SenderResponse
}

type SenderPublishResponse struct {
	SenderResponse
}
