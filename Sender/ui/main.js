import {$} from "@sciter";

document.on("click", "button#btn_create_topic", function() {
    var topic_input_field = $("#topic_input_field");
    const topicName = topic_input_field.value;
    Window.this.MainFrame.createTopic(topicName);
    topic_input_field.value = "";
});