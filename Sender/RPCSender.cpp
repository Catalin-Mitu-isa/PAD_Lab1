#include "RPCSender.h"

RPCSender::RPCSender()
{
    std::stringstream channelAddress;
    channelAddress << BROKER_NAME << ':' << BROKER_GRPC_PORT;
    auto channel = grpc::CreateChannel(channelAddress.str(), grpc::InsecureChannelCredentials());
    m_stub = sender::SenderService::NewStub(std::move(channel));
}

bool RPCSender::createTopic(const std::string topic)
{
    grpc::ClientContext ctx;
    sender::CreateTopicRequest request;
    sender::CreateTopicResponse response;
    request.set_name(topic);
    m_stub->CreateTopic(&ctx, request, &response);
    const bool createSuccessful = response.success();
    if (createSuccessful)
    {
        std::cout << "Topic prin rpc creat" << std::endl;
        m_topicName = topic;
    }
    return createSuccessful;
}

bool RPCSender::publishMessage(const std::string message)
{
    grpc::ClientContext ctx;
    sender::PublishMessageRequest request;
    sender::PublishMessageResponse response;
    request.set_message(message);
    request.set_topicname(m_topicName);
    m_stub->PublishMessage(&ctx, request, &response);
    return response.success();
}