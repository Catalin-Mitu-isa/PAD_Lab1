import {$} from "@sciter";

document.on("click", "button#btn_subscribe_to_topic", function() {
    var topic_input_field = $("#topic_input_field");
    const topicName = topic_input_field.value;
    if (topicName == "")
        return;

    const subscribed = Window.this.MainFrame.subscribeToTopic(topicName);
    if (subscribed)
    {
        $("#subscribe_to_topic_section").style.display = "none";
        $("#messages_section").style.display = "block";
        $("#connection_section").style.display = "none";
    }
    topic_input_field.value = "";
});

document.on("click", "#connection_section > label", function() {
    Window.this.MainFrame.toggleConnection();
});

function displayMessage(message)
{
    var messagesSection = $("#messages_section");
    const mesLine = "<p .message>"+message+"</p>";
    messagesSection.append(mesLine);
}

globalThis.displayMessage = displayMessage;