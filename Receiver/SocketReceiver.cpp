#include "SocketReceiver.h"

SocketReceiver::SocketReceiver()
    : m_socket(m_io_context)
{
    connectToBroker();
}

SocketReceiver::~SocketReceiver()
{
    m_socket.close();
}

void SocketReceiver::connectToBroker()
{
    asio::ip::tcp::resolver resolver(m_io_context);
    const auto endpoints = resolver.resolve(BROKER_NAME
            , std::to_string(BROKER_SOCKET_PORT));

    m_socket.connect(*(endpoints.begin()));
}

std::size_t SocketReceiver::sendStrSync(const std::string & str)
{
    if (m_socket.is_open())
        return m_socket.write_some(asio::buffer(str.data(), str.size()));
    else
        return 0;
}

std::string SocketReceiver::receiveStrSync()
{
    if (!m_socket.is_open())
        return {};

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

    return stringStream.str();
}

bool SocketReceiver::subscribeToTopic(std::string topic)
{
    std::stringstream jsonMessage;
    jsonMessage << R"({"action": "SUBSCRIBE", "topic_to_subscribe": ")"
        << topic
        << "\"}";

    if (m_socket.is_open())
        if (sendStrSync(jsonMessage.str()))
            if (!receiveStrSync().empty())
                return true;

    return false;
}

void SocketReceiver::listenForMessages(std::function<void(const std::string &)> handler)
{
    m_listenerThread = std::thread([this, handler]() -> void {
       while (m_socket.is_open())
           handler(receiveStrSync());
    });
}