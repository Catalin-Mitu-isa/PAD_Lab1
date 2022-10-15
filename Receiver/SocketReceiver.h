#ifndef SENDER_SOCKETRECEIVER_H
#define SENDER_SOCKETRECEIVER_H

#include "IReceiver.h"
#include <asio.hpp>
#include <sstream>
#include <array>
#include <iostream>
#include <thread>

#define BROKER_SOCKET_PORT 43200
#define BROKER_NAME "broker.pad.utm.md"

class SocketReceiver : public IReceiver
{
public:
    SocketReceiver();
    ~SocketReceiver();
    bool subscribeToTopic(std::string topic) override;
    void listenForMessages(std::function<void(const std::string &)> handler) override;

private:
    void connectToBroker();
    std::size_t sendStrSync(const std::string & str);
    std::string receiveStrSync();

    asio::io_context m_io_context;
    asio::ip::tcp::socket m_socket;
    std::thread m_listenerThread;
};

#endif //SENDER_SOCKETRECEIVER_H
