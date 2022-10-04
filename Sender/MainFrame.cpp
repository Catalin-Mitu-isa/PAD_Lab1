#include "MainFrame.h"

MainFrame::MainFrame()
    : window(SW_TITLEBAR
        | SW_RESIZEABLE
        | SW_CONTROLS
        | SW_MAIN
        | SW_ENABLE_DEBUG)
{
    m_sender = &m_socketSender;
}

bool MainFrame::createTopic(sciter::value s_value)
{
    std::string topicName = s_value.get<sciter::astring>();
    if (topicName.empty())
        return false;

    std::cout << "Create topic. Topic name: " << topicName << std::endl;
    return m_sender->createTopic(topicName);
}

bool MainFrame::publishMessage(sciter::value s_value)
{
    std::string message = s_value.get<sciter::astring>();
    if (message.empty())
        return false;

    std::cout << "Publish message. Message: " << message << std::endl;
    return m_sender->publishMessage(message);
}

void MainFrame::toggleConnection()
{
    std::cout << "Toggle connection" << std::endl;
    if (m_sender == &m_socketSender)
        m_sender = &m_rpcSender;
    else
        m_sender = &m_socketSender;
}