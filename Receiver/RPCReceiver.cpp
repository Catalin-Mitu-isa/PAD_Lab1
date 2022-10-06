#include "RPCReceiver.h"

bool RPCReceiver::subscribeToTopic(std::string topic)
{
    return true;
}

void RPCReceiver::listenForMessages(std::function<void(const std::string &)> handler)
{

}