#ifndef SENDER_GRPCSENDER_H
#define SENDER_GRPCSENDER_H

#include "IReceiver.h"

class RPCReceiver : public IReceiver
{
public:
    bool subscribeToTopic(std::string topic) override;
    void listenForMessages(std::function<void(const std::string &)> handler) override;
};


#endif //SENDER_GRPCSENDER_H
