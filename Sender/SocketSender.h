#ifndef SENDER_SOCKETSENDER_H
#define SENDER_SOCKETSENDER_H

#include "ISender.h"
#include <asio.hpp>
#include <sstream>
#include <array>
#include <iostream>

#define BROKER_SOCKET_PORT 8080
#define BROKER_NAME "broker.pad.utm.md"

class SocketSender : public ISender
{
public:
    SocketSender();
    bool createTopic(const std::string topic) override;
    bool publishMessage(const std::string message) override;

private:
    void connectToBroker();
    bool readFromSocket(std::iostream & dataStream);
    void asyncReader(const asio::error_code & error, std::size_t bytesTransfered);

    asio::io_context m_io_context;
    asio::ip::tcp::socket m_socket;
    asio::streambuf m_streambuf;
};

#endif //SENDER_SOCKETSENDER_H
