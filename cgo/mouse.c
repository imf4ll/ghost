#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <unistd.h>
#include <X11/Xlib.h>
#include <X11/Xutil.h>
#include <X11/X.h>
#include <X11/cursorfont.h>
#include <cairo/cairo.h>
#include <cairo/cairo-xlib.h>

struct MousePosition {
    int x1,
        y1,
        x2,
        y2;
};

static struct MousePosition mousePosition[1];

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

    XVisualInfo vinfo;
    if (!XMatchVisualInfo(display, DefaultScreen(display), 32, TrueColor, &vinfo)) {
        printf("No visual found supporting 32 bit color, terminating\n");
        exit(EXIT_FAILURE);
    }

    attrs.colormap = XCreateColormap(display, root, vinfo.visual, AllocNone);
    attrs.background_pixel = 0;
    attrs.border_pixel = 0;

    Window overlay;

    if (mousePosition->x1 < mousePosition->x2) {
        overlay = XCreateWindow (
            display, root,
            mousePosition->x1, mousePosition->y1, mousePosition->x2 - mousePosition->x1, mousePosition->y2 - mousePosition->y1, 0,
            vinfo.depth, InputOutput,
            vinfo.visual,
            CWOverrideRedirect | CWColormap | CWBackPixel | CWBorderPixel, &attrs
        );
    
    } else {
        overlay = XCreateWindow (
            display, root,
            mousePosition->x2, mousePosition->y2, mousePosition->x1 - mousePosition->x2, mousePosition->y1 - mousePosition->y2, 0,
            vinfo.depth, InputOutput,
            vinfo.visual,
            CWOverrideRedirect | CWColormap | CWBackPixel | CWBorderPixel, &attrs
        );
    }

    XMapWindow(display, overlay);

    cairo_surface_t* surf = cairo_xlib_surface_create(display, overlay, vinfo.visual, 0, 0);
    cairo_t* cr = cairo_create(surf);

    XFlush(display);

    sleep(1);

    cairo_destroy(cr);
    cairo_surface_destroy(surf);

    XUnmapWindow(display, overlay);
    XCloseDisplay(display);

    XUngrabPointer(display, CurrentTime);

    return mousePosition;
}

struct MousePosition *getMouseAndKeyboard(void) {
    Display *display = XOpenDisplay(NULL);
    Window root = XRootWindow(display, 0);
    XEvent event;

    XSetWindowAttributes attrs;
    attrs.override_redirect = true;

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

    XVisualInfo vinfo;
    if (!XMatchVisualInfo(display, DefaultScreen(display), 32, TrueColor, &vinfo)) {
        printf("No visual found supporting 32 bit color, terminating\n");
        exit(EXIT_FAILURE);
    }

    attrs.colormap = XCreateColormap(display, root, vinfo.visual, AllocNone);
    attrs.background_pixel = 0;
    attrs.border_pixel = 0;

    Window overlay;

    if (mousePosition->x1 < mousePosition->x2) {
        overlay = XCreateWindow (
            display, root,
            mousePosition->x1, mousePosition->y1, mousePosition->x2 - mousePosition->x1, mousePosition->y2 - mousePosition->y1, 0,
            vinfo.depth, InputOutput,
            vinfo.visual,
            CWOverrideRedirect | CWColormap | CWBackPixel | CWBorderPixel, &attrs
        );
    
    } else {
        overlay = XCreateWindow (
            display, root,
            mousePosition->x2, mousePosition->y2, mousePosition->x1 - mousePosition->x2, mousePosition->y1 - mousePosition->y2, 0,
            vinfo.depth, InputOutput,
            vinfo.visual,
            CWOverrideRedirect | CWColormap | CWBackPixel | CWBorderPixel, &attrs
        );
    }

    XMapWindow(display, overlay);

    cairo_surface_t* surf = cairo_xlib_surface_create(display, overlay, vinfo.visual, 0, 0);
    cairo_t* cr = cairo_create(surf);

    XFlush(display);

    sleep(1);

    cairo_destroy(cr);
    cairo_surface_destroy(surf);

    XUnmapWindow(display, overlay);
    XCloseDisplay(display);

    XUngrabPointer(display, CurrentTime);
    XUngrabKeyboard(display, CurrentTime);

    return mousePosition;
}
