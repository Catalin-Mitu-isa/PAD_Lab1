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
    bool createTopic(std::string topic) override;
    bool publishMessage(std::string message) override;

private:
    void connectToBroker();
    std::size_t sendStr(const std::string & str);
    std::string receiveStr();

    asio::io_context m_io_context;
    asio::ip::tcp::socket m_socket;
    std::string m_topicName;
};

#endif //SENDER_SOCKETSENDER_H
