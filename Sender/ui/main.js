import {$} from "@sciter";

document.on("click", "button#btn_create_topic", function() {
    var topic_input_field = $("#topic_input_field");
    const topicName = topic_input_field.value;
    if (topicName == "")
        return;

    const topicCreated = Window.this.MainFrame.createTopic(topicName);
    if (topicCreated)
    {
        $("#create_topic_section").style.display = "none";
        $("#message_input_section").style.display = "block";
    }
    topic_input_field.value = "";
});

document.on("click", "button#btn_publish_message", function() {
    var message_input_field = $("#message_textarea");
    const message = message_input_field.value;
    if (message == "")
        return;

    const messageSent = Window.this.MainFrame.publishMessage(message);
    if (messageSent)
    {
    }
    message_input_field.value = "";
});

document.on("click", "#connection_section > label", function() {
    Window.this.MainFrame.toggleConnection();
});