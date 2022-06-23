#include <stdio.h>
#include <X11/Xlib.h>
#include <cairo/cairo-xlib.h>
#include <cairo/cairo.h>

void captureScreen(char * filename) {
    Display * display = XOpenDisplay(NULL);
    Window root = DefaultRootWindow(display);
    int screen = DefaultScreen(display);
    Visual * visual = DefaultVisual(display, screen);
    XWindowAttributes windowAtts;

    XGetWindowAttributes(display, root, &windowAtts);

    cairo_surface_t * surface = cairo_xlib_surface_create (
        display,
        root,
        visual,
        windowAtts.width,
        windowAtts.height
    );

    cairo_surface_write_to_png(surface, filename);

    cairo_surface_destroy(surface);
}

void captureRect(double x, double y, double width, double height, char * filename) {
    Display * display = XOpenDisplay(NULL);
    Window root = DefaultRootWindow(display);
    int screen = DefaultScreen(display);
    Visual * visual = DefaultVisual(display, screen);
    XWindowAttributes windowAtts;
    
    XGetWindowAttributes(display, root, &windowAtts);

    cairo_surface_t * surface = cairo_xlib_surface_create (
        display,
        root,
        visual,
        windowAtts.width,
        windowAtts.height
    );

    cairo_surface_t * rectangle = cairo_surface_create_for_rectangle (
        surface,
        x,
        y,
        width,
        height
    );

    cairo_surface_write_to_png(rectangle, filename);
    
    cairo_surface_destroy(surface);
}
