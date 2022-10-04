#include "SocketSender.h"

SocketSender::SocketSender()
    : m_socket(m_io_context)
{ }

void SocketSender::connectToBroker()
{
    asio::ip::tcp::resolver resolver(m_io_context);
    const auto endpoints = resolver.resolve(BROKER_NAME
            , std::to_string(BROKER_SOCKET_PORT));

//    m_tcpStream.connect(BROKER_NAME, std::to_string(BROKER_SOCKET_PORT));
    m_socket.connect(*(endpoints.begin()));
}

bool SocketSender::createTopic(const std::string topic)
{
    connectToBroker();

    asio::write(m_socket, asio::buffer(topic));

    asio::async_read(m_socket, m_streambuf, &SocketSender::asyncReader);

//    asio::buffer ceva;
//    std::stringstream dataStream;
//    m_tcpStream
//    m_socket.write_some(asio::buffer(topic, topic.size()));


//    readFromSocket(dataStream);

//    std::cout << "Data read: " << dataStream.str() << std::endl;
    return true;
}

bool SocketSender::readFromSocket(std::iostream & dataStream)
{
    std::array<char, 1024> buffer{};
    std::uint16_t bytesRead; // uint16 because max nr is 1024
    asio::error_code errorCode;


    while (true)
    {
//        std::cout << "reading again" << std::endl;
//        bytesRead = m_socket.receive(asio::buffer(buffer), 0, errorCode);
        bytesRead = m_socket.read_some(asio::buffer(buffer), errorCode);
//        dataStream.write(buffer.data(), bytesRead);
        std::cout.write(buffer.data(), bytesRead);
        std::cout << std::flush;
        if (errorCode == asio::error::eof)
            break;
    }

    return true;
}

bool SocketSender::publishMessage(const std::string message)
{

    return true;
}
