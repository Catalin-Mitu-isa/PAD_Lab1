#ifndef SENDER_GRPCSENDER_H
#define SENDER_GRPCSENDER_H

#include "IReceiver.h"
#include "Receiver.pb.h"
#include "Receiver.grpc.pb.h"
#include <sstream>
#include <memory>
#include <thread>
#include <grpcpp/create_channel.h>

#define BROKER_GRPC_PORT 8081
#define BROKER_NAME "broker.pad.utm.md"

class RPCReceiver : public IReceiver
{
public:
    RPCReceiver();
    ~RPCReceiver();
    bool subscribeToTopic(std::string topic) override;
    void listenForMessages(std::function<void(const std::string &)> handler) override;

private:
    std::unique_ptr<SubscriberService::Stub> m_stub;
    std::thread m_listenerThread;
    bool m_listeningForMessages;
};


#endif //SENDER_GRPCSENDER_H
