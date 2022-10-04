#ifndef SENDER_MAINFRAME_H
#define SENDER_MAINFRAME_H

#include <sciter-x.h>
#include <sciter-x-window.hpp>
#include <iostream>
#include "SocketSender.h"
#include "RPCSender.h"

class MainFrame: public sciter::window
{
public:
    MainFrame();

    bool createTopic(sciter::value s_value);
    bool publishMessage(sciter::value s_value);
    void toggleConnection();

    SOM_PASSPORT_BEGIN(MainFrame)
        SOM_FUNCS(
            SOM_FUNC(createTopic),
            SOM_FUNC(publishMessage),
            SOM_FUNC(toggleConnection)
        )
    SOM_PASSPORT_END

private:
    SocketSender m_socketSender;
    RPCSender m_rpcSender;
    ISender * m_sender;
};

#endif //SENDER_MAINFRAME_H
