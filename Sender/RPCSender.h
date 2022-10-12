#ifndef SENDER_GRPCSENDER_H
#define SENDER_GRPCSENDER_H

#define BROKER_GRPC_PORT 8081
#define BROKER_NAME "broker.pad.utm.md"

#include <sstream>
#include <memory>
#include "ISender.h"
#include "Sender.pb.h"
#include "Sender.grpc.pb.h"
#include <grpcpp/create_channel.h>

class RPCSender : public ISender
{
public:
    RPCSender();
    bool createTopic(const std::string topic) override;
    bool publishMessage(const std::string message) override;

private:
    std::unique_ptr<SenderService::Stub> m_stub;
};


#endif //SENDER_GRPCSENDER_H
