#include "checkOS.c"
#include <stdio.h>

#if defined (UNIX)
    #include <X11/Xlib.h>
    #include <X11/extensions/Xrandr.h>
    #include <X11/extensions/randr.h>

    struct DisplayInfo {
        int x;
        int y;
        int w;
        int h;
    };

    struct DisplayInfo * getDisplays() {
        static struct DisplayInfo displayInfos[10];

        Display * dpy = XOpenDisplay(NULL);
        XRRScreenResources * screen = XRRGetScreenResources(dpy, DefaultRootWindow(dpy));
        XRRCrtcInfo * crtcInfo;
        int ncrtc = screen -> ncrtc;
        int turns = 0;

        for (int i = 0; i < ncrtc; i++) {
            crtcInfo = XRRGetCrtcInfo(dpy, screen, screen -> crtcs[i]);

            if (crtcInfo -> width != 0) {
                struct DisplayInfo displayInfo;

                displayInfo.x = crtcInfo -> x;
                displayInfo.y = crtcInfo -> y;
                displayInfo.w = crtcInfo -> width;
                displayInfo.h = crtcInfo -> height;

                displayInfos[turns] = displayInfo;
                turns += 1;
            }

            XRRFreeCrtcInfo(crtcInfo);
        }

        XRRFreeScreenResources(screen);

        return displayInfos;
    }
#endif

#if defined (WINDOWS)
    #include <winuser.h>
    #include <stdio.h>

    struct DisplayInfo {
        int x;
        int y;
        int w;
        int h;
    };

    static struct DisplayInfo displayInfos[10];
    int turns = 0;

    BOOL CALLBACK MonitorEnum(HMONITOR hMon, HDC hdc, LPRECT lprc, LPARAM dwData) {
        MONITORINFO info;
        
        info.cbSize = sizeof(info);
        
        if (GetMonitorInfo(hMon, &info)) {
            struct DisplayInfo displayInfo;
            
            displayInfo.x = info.rcMonitor.left;
            displayInfo.y = info.rcMonitor.top;
            displayInfo.w = info.rcMonitor.right;
            displayInfo.h = info.rcMonitor.bottom;
            
            displayInfos[turns] = displayInfo;
            turns += 1;
        }
        
        return TRUE;
    }

    struct DisplayInfo * getDisplays() {
        EnumDisplayMonitors(NULL, NULL, MonitorEnum, 0);
        
        return displayInfos;
    }
#endif
