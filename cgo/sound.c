#include "checkOS.c"

#if defined (UNIX)
    void Play() {
        // Nothing here, for now :)
    }
#endif

#if defined (WINDOWS)
    #include <windows.h>
    #include <mmsystem.h>

    void Play() {
        PlaySound((LPCSTR)SND_ALIAS_SYSTEMEXCLAMATION, NULL, SND_ALIAS_ID);

    }
#endif
