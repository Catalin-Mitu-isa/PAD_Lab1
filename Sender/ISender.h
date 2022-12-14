#ifndef SENDER_ISENDER_H
#define SENDER_ISENDER_H

#include <string>

class ISender {
public:
    virtual bool createTopic(std::string topic) = 0;
    virtual bool publishMessage(std::string message) = 0;
};

#endif //SENDER_ISENDER_H
