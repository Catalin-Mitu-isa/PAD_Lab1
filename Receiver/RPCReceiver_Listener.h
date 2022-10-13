#ifndef RECEIVER_RPCRECEIVER_LISTENER_H
#define RECEIVER_RPCRECEIVER_LISTENER_H

#include "Broker.pb.h"
#include "Broker.grpc.pb.h"
#include <functional>

class RPCReceiver_Listener : public broker::BrokerService::Service
{
public:
    grpc::Status SendMessage(grpc::ServerContext * , const broker::SendMessageRequest * request, broker::SendMessageResponse * response) override;
    void setMsgHandler(const std::function<void(const std::string &)> & handler);

private:
    std::function<void(const std::string)> m_msgHandler;
};

#endif //RECEIVER_RPCRECEIVER_LISTENER_H
