#ifndef SENDER_MAINFRAME_H
#define SENDER_MAINFRAME_H

#include <sciter-x.h>
#include <sciter-x-window.hpp>
#include <iostream>
#include "SocketReceiver.h"
#include "RPCReceiver.h"

class MainFrame: public sciter::window
{
public:
    MainFrame();

    void toggleConnection();
    bool subscribeToTopic(sciter::value topic);

    SOM_PASSPORT_BEGIN(MainFrame)
        SOM_FUNCS(
            SOM_FUNC(toggleConnection),
            SOM_FUNC(subscribeToTopic)
        )
    SOM_PASSPORT_END

private:
    SocketReceiver m_socketReceiver;
    RPCReceiver m_rpcReceiver;
    IReceiver * m_sender;
};

#endif //SENDER_MAINFRAME_H
