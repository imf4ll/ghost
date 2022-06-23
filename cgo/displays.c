#include <stdio.h>
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
