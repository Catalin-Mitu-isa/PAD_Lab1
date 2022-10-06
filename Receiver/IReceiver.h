#ifndef SENDER_IRECEIVER_H
#define SENDER_IRECEIVER_H

#include <string>
#include <functional>

class IReceiver {
public:
    virtual bool subscribeToTopic(std::string topic) = 0;
    virtual void listenForMessages(std::function<void(const std::string &)> handler) = 0;
};

#endif //SENDER_IRECEIVER_H
