#ifndef SENDER_GRPCSENDER_H
#define SENDER_GRPCSENDER_H

#include "IReceiver.h"
#include "Receiver.pb.h"
#include "Receiver.grpc.pb.h"
#include "Broker.pb.h"
#include "Broker.grpc.pb.h"
#include "RPCReceiver_Listener.h"
#include <sstream>
#include <memory>
#include <thread>
#include <grpcpp/create_channel.h>
#include <grpcpp/server_builder.h>

#define BROKER_GRPC_PORT 43201
#define BROKER_NAME "broker.pad.utm.md"

#define RECEIVER_GRPC_PORT 8081
#define RECEIVER_HOSTNAME "UFO-L"

class RPCReceiver : public IReceiver
{
public:
    RPCReceiver();
    ~RPCReceiver();
    bool subscribeToTopic(std::string topic) override;
    void listenForMessages(std::function<void(const std::string &)> handler) override;

private:
    void createStub();
    void createListenerServer();

    std::unique_ptr<receiver::ReceiverService::Stub> m_stub;
    std::unique_ptr<grpc::Server> m_server;
    RPCReceiver_Listener m_listenerService;
    std::thread m_listenerThread;
    bool m_listeningForMessages;
};

#endif //SENDER_GRPCSENDER_H
