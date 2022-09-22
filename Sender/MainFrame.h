#ifndef SENDER_MAINFRAME_H
#define SENDER_MAINFRAME_H

#include <sciter-x.h>
#include <sciter-x-window.hpp>
#include "Sender.h"
#include <iostream>

class MainFrame: public sciter::window
{
public:
    MainFrame();

    bool createTopic(sciter::value s_value);
    bool publishMessage(sciter::value s_value);

    SOM_PASSPORT_BEGIN(MainFrame)
        SOM_FUNCS(
            SOM_FUNC(createTopic),
            SOM_FUNC(publishMessage)
        )
    SOM_PASSPORT_END

private:
    Sender m_sender;
};

#endif //SENDER_MAINFRAME_H
