#include "RPCReceiver.h"

RPCReceiver::RPCReceiver()
    : m_listeningForMessages(false)
{
    createStub();
    createListenerServer();
}

RPCReceiver::~RPCReceiver()
{
    m_listeningForMessages = false;
    m_server->Shutdown();
    m_server->Wait();
}

void RPCReceiver::createStub()
{
    std::stringstream channelAddress;
    channelAddress << BROKER_NAME << ':' << BROKER_GRPC_PORT;
    auto channel = grpc::CreateChannel(channelAddress.str(), grpc::InsecureChannelCredentials());
    m_stub = receiver::ReceiverService::NewStub(std::move(channel));
}

void RPCReceiver::createListenerServer()
{
    grpc::ServerBuilder serverBuilder;
    {
        std::stringstream ss;
        const int randomPort = rand() % 30000 + 10000 // interval 10000; 40000
        ss << RECEIVER_HOSTNAME << ':' << randomPort;
        serverBuilder.AddListeningPort(ss.str(), grpc::InsecureServerCredentials());
    }
    serverBuilder.RegisterService(&m_listenerService);
    m_server = serverBuilder.BuildAndStart();
}

bool RPCReceiver::subscribeToTopic(std::string topic)
{
    grpc::ClientContext ctx;
    receiver::SubscribeRequest request;
    receiver::SubscribeResponse response;
    request.set_topicname(topic);
    request.set_hostname(RECEIVER_HOSTNAME);
    request.set_port(RECEIVER_GRPC_PORT);
    m_stub->Subscribe(&ctx, request, &response);
    return response.success();
}

void RPCReceiver::listenForMessages(std::function<void(const std::string &)> handler)
{
    m_listenerService.setMsgHandler(handler);
}