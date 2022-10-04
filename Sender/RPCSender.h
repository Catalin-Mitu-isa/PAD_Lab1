#ifndef SENDER_GRPCSENDER_H
#define SENDER_GRPCSENDER_H

#include "ISender.h"

class RPCSender : public ISender
{
public:
    bool createTopic(const std::string topic) override;
    bool publishMessage(const std::string message) override;
};


#endif //SENDER_GRPCSENDER_H
