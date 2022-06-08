#include "checkOS.c"

#if defined (WINDOWS)
    #include <windows.h>

    void setClipboard(char * screenshot) {
        const size_t screenshotSize = strlen(screenshot) + 1;

        HGLOBAL hMem = GlobalAlloc(GMEM_MOVEABLE, screenshotSize);
        memcpy(GlobalLock(hMem), screenshot, screenshotSize);
        GlobalUnlock(hMem);

        OpenClipboard(NULL);
        EmptyClipboard();
        SetClipboardData(CF_DIB, hMem);

        CloseClipboard();
    }
#endif
