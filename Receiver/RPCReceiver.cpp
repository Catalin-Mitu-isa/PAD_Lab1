#include "RPCReceiver.h"

RPCReceiver::RPCReceiver()
    : m_listeningForMessages(false)
{
    std::stringstream channelAddress;
    channelAddress << BROKER_NAME << ':' << BROKER_GRPC_PORT;
    auto channel = grpc::CreateChannel(channelAddress.str(), grpc::InsecureChannelCredentials());
    m_stub = SubscriberService::NewStub(std::move(channel));
}

RPCReceiver::~RPCReceiver()
{
    m_listeningForMessages = false;
}

bool RPCReceiver::subscribeToTopic(std::string topic)
{
    grpc::ClientContext ctx;
    SubscribeToTopicRequest request;
    SubscribeToTopicResponse response;
    request.set_name(topic);
    m_stub->SubscribeToTopic(&ctx, request, &response);
    return true;
}

void RPCReceiver::listenForMessages(std::function<void(const std::string &)> handler)
{
    m_listeningForMessages = true;
    m_listenerThread = std::thread([this, handler=handler]() ->void {
        while (m_listeningForMessages)
        {
            grpc::ClientContext ctx;
            ReceiveMessageRequest request;
            ReceiveMessageResponse response;
            m_stub->ReceiveMessage(&ctx, request, &response);
            std::cout << "Message received: " << response.message() << std::endl;
            handler(response.message());
        }
    });
}