#include "SocketSender.h"

SocketSender::SocketSender()
    : m_socket(m_io_context)
{
    try {
        connectToBroker();
    } catch (const std::exception & exc) {
        std::cout << exc.what() << std::endl;
    }
}

void SocketSender::connectToBroker()
{
    asio::ip::tcp::resolver resolver(m_io_context);
    const auto endpoints = resolver.resolve(BROKER_NAME
            , std::to_string(BROKER_SOCKET_PORT));

    m_socket.connect(*(endpoints.begin()));
}

std::size_t SocketSender::sendStr(const std::string & str)
{
    if (m_socket.is_open())
    {
        std::cout << __FUNCTION__ << ": socket is open. Send: " << str << std::endl;
        return m_socket.write_some(asio::buffer(str.data(), str.size()));
    }
    else
    {
        std::cout << __FUNCTION__
            << ": socket is closed. attempted to send: " << str << std::endl;
        return 0;
    }
}

std::string SocketSender::receiveStr()
{
    if (!m_socket.is_open())
    {
        std::cout << __FUNCTION__ << ": socket closed" << std::endl;
        return {};
    }

    std::stringstream stringStream;
    m_socket.wait(m_socket.wait_read);

    while (true)
    {
        const std::size_t availableBytes = m_socket.available();
        if (availableBytes == 0)
            break;

        std::vector<char> buffer(availableBytes);
        std::size_t bytesRead = m_socket.read_some(
                asio::buffer(buffer.data(), buffer.size()));
        stringStream.write(buffer.data(), bytesRead);
    }

    std::cout << __FUNCTION__ << ": received string: " << stringStream.str() << std::endl;

    return stringStream.str();
}

bool SocketSender::createTopic(std::string topic)
{
    std::stringstream jsonMessage;
    jsonMessage << R"({"action": "CREATE_TOPIC", "topic_name": ")"
        << topic
        << "\"}\r\n\r\n";

    if (m_socket.is_open())
        if (sendStr(jsonMessage.str()))
        {
            const std::string receivedStr = receiveStr();
            if (!receivedStr.empty())
            {
                if (m_topicName.empty())
                    m_topicName = topic;
                return true;
            }
        }

    return false;
}

bool SocketSender::publishMessage(std::string message)
{
    std::cout << __FUNCTION__ << std::endl;
    std::stringstream jsonMessage;
    jsonMessage << R"({"action": "PUBLISH_MESSAGE", "message": ")"
                << message
                << R"(", "topic_name": ")"
                << m_topicName
                << "\"}\r\n\r\n";

    if (m_socket.is_open())
        if (sendStr(jsonMessage.str()))
        {
            const std::string ceAmPrimit = receiveStr();
            std::cout << __FUNCTION__ << ". ce am primit: " << ceAmPrimit << std::endl;
            if (!ceAmPrimit.empty())
                return true;
        }

    return false;
}
