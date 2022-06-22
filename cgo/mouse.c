#include <stdio.h>
#include <stdlib.h>
#include "checkOS.c"

struct MousePosition {
    int x1,
        y1,
        x2,
        y2;
};

static struct MousePosition mousePosition[1];

#if defined (UNIX)
    #include <X11/Xlib.h>
    #include <X11/Xutil.h>
    #include <X11/X.h>
    #include <X11/cursorfont.h>

    struct MousePosition *getMouse(void) {
        Display *display = XOpenDisplay(NULL);
        Window root = XRootWindow(display, 0);
        XEvent event;

        Cursor cursor = XCreateFontCursor(display, XC_crosshair);
        XGrabPointer(display, root, False, ButtonPressMask, GrabModeAsync, GrabModeAsync, None, cursor, CurrentTime);

        for (int i = 0; i < 2; i++) {
            XWindowEvent(display, root, ButtonPressMask | KeyPressMask, &event);

            switch (event.type) {
                case ButtonPress:
                    switch (event.xbutton.button) {
                        case Button1:
                            if (i == 0) {
                                mousePosition->x1 = event.xbutton.x_root;
                                mousePosition->y1 = event.xbutton.y_root;

                            } else {
                                mousePosition->x2 = event.xbutton.x_root;
                                mousePosition->y2 = event.xbutton.y_root;

                            }

                            break;
                    }

                    break;

                default:
                    i--;
            }
        }

        XUngrabPointer(display, CurrentTime);

        return mousePosition;
    }

    struct MousePosition *getMouseAndKeyboard(void) {
        Display *display = XOpenDisplay(NULL);
        Window root = XRootWindow(display, 0);
        XEvent event;

        Cursor cursor = XCreateFontCursor(display, XC_crosshair);
        XGrabPointer(display, root, False, ButtonPressMask, GrabModeAsync, GrabModeAsync, None, cursor, CurrentTime);
        XGrabKeyboard(display, root, KeyPressMask, GrabModeAsync, GrabModeAsync, CurrentTime);

        for (int i = 0; i < 2; i++) {
            XWindowEvent(display, root, ButtonPressMask | KeyPressMask, &event);

            switch (event.type) {
                case ButtonPress:
                    switch (event.xbutton.button) {
                        case Button1:
                            if (i == 0) {
                                mousePosition->x1 = event.xbutton.x_root;
                                mousePosition->y1 = event.xbutton.y_root;

                            } else {
                                mousePosition->x2 = event.xbutton.x_root;
                                mousePosition->y2 = event.xbutton.y_root;

                            }

                            break;
                    }

                    break;

                case KeyPress:
                    switch (event.xkey.keycode) {
                        case 9:
                            exit(3);

                            break;
                    }
            
                default:
                    i--;
            }
        }

        XUngrabPointer(display, CurrentTime);
        XUngrabKeyboard(display, CurrentTime);

        return mousePosition;
    }
#endif
