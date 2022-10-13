#include "RPCReceiver_Listener.h"

grpc::Status RPCReceiver_Listener::SendMessage(
        grpc::ServerContext *
        , const broker::SendMessageRequest * request
        , broker::SendMessageResponse * response)
{
    m_msgHandler(request->message());
    return grpc::Status::OK;
}

void RPCReceiver_Listener::setMsgHandler(
        const std::function<void(const std::string &)> & handler)
{
    m_msgHandler = handler;
}