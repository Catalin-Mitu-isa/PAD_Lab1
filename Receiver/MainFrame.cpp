#include "MainFrame.h"

MainFrame::MainFrame()
    : window(SW_TITLEBAR
        | SW_RESIZEABLE
        | SW_CONTROLS
        | SW_MAIN
        | SW_ENABLE_DEBUG)
{
    m_sender = &m_socketReceiver;
}

void MainFrame::toggleConnection()
{
    std::cout << "Toggle connection" << std::endl;
    if (m_sender == &m_socketReceiver)
        m_sender = &m_rpcReceiver;
    else
        m_sender = &m_socketReceiver;
}

bool MainFrame::subscribeToTopic(sciter::value topic)
{
    const bool subscribed = m_sender->subscribeToTopic(topic.get<std::string>());

    if (subscribed)
    {
        m_sender->listenForMessages([&](const std::string & message) -> void {
            this->call_function("displayMessage", message);
        });
        return true;
    }

    return false;
}
