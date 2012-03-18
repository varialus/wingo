package main

import "code.google.com/p/jamslam-x-go-binding/xgb"

import (
    "github.com/BurntSushi/xgbutil"
    "github.com/BurntSushi/xgbutil/xevent"
    "github.com/BurntSushi/xgbutil/xrect"
    "github.com/BurntSushi/xgbutil/xwindow"
)

type window struct {
    id xgb.Id
    geom xrect.Rect
}

const (
    DoX = xgb.ConfigWindowX
    DoY = xgb.ConfigWindowY
    DoW = xgb.ConfigWindowWidth
    DoH = xgb.ConfigWindowHeight
    DoBorder = xgb.ConfigWindowBorderWidth
    DoSibling = xgb.ConfigWindowSibling
    DoStack = xgb.ConfigWindowStackMode
)

func newWindow(id xgb.Id) *window {
    return &window{
        id: id,
        geom: xrect.Make(0, 0, 1, 1),
    }
}

func createWindow(parent xgb.Id, mask uint32, vals []uint32) *window {
    wid := X.Conn().NewId()
    scrn := X.Screen()

    X.Conn().CreateWindow(scrn.RootDepth, wid, parent, 0, 0, 1, 1, 0,
                          xgb.WindowClassInputOutput, scrn.RootVisual,
                          mask, vals)

    return newWindow(wid)
}

func (w *window) listen(masks int) {
    xwindow.Listen(X, w.id, masks)
}

func (w *window) map_() {
    X.Conn().MapWindow(w.id)
}

func (w *window) unmap() {
    X.Conn().UnmapWindow(w.id)
}

func (w *window) geometry() (xrect.Rect, error) {
    var err error
    w.geom, err = xwindow.RawGeometry(X, w.id)
    if err != nil {
        return nil, err
    }
    return w.geom, nil
}

func (w *window) kill() {
    X.Conn().KillClient(uint32(w.id))
}

func (w *window) destroy() {
    X.Conn().DestroyWindow(w.id)
    xevent.Detach(X, w.id)
}

func (w *window) focus() {
    X.Conn().SetInputFocus(xgb.InputFocusPointerRoot, w.id, 0)
}

// moveresize is a wrapper around configure that only accepts parameters
// related to size and position.
func (win *window) moveresize(flags uint16, x, y int16, w, h uint16) {
    // Kill any hopes of stacking
    flags = uint16((int16(flags) & ^DoSibling) & ^DoStack)
    win.configure(flags, x, y, w, h, xgb.Id(0), 0)
}

// configure is the method version of 'configure'.
// It is duplicated because we need to update our idea of the window's
// geometry. (We don't want another set of 'if' statements because it
// needs to be as efficient as possible.)
func (win *window) configure(flags uint16, x, y int16, w, h uint16,
                             sibling xgb.Id, stackMode byte) {
    vals := []uint32{}

    if DoX & flags > 0 {
        vals = append(vals, uint32(x))
        win.geom.XSet(x)
    }
    if DoY & flags > 0 {
        vals = append(vals, uint32(y))
        win.geom.YSet(y)
    }
    if DoW & flags > 0 {
        vals = append(vals, uint32(w))
        win.geom.WidthSet(w)
    }
    if DoH & flags > 0 {
        vals = append(vals, uint32(h))
        win.geom.HeightSet(h)
    }
    if DoSibling & flags > 0 {
        vals = append(vals, uint32(sibling))
    }
    if DoStack & flags > 0 {
        vals = append(vals, uint32(stackMode))
    }

    // Don't send anything if we have nothing to send
    if len(vals) == 0 {
        return
    }

    X.Conn().ConfigureWindow(win.id, flags, vals)
}

// configure is a nice wrapper around ConfigureWindow.
// We purposefully omit 'BorderWidth' because I don't think it's ever used
// any more.
func configure(window xgb.Id, flags uint16, x, y int16, w, h uint16,
               sibling xgb.Id, stackMode byte) {
    vals := []uint32{}

    if DoX & flags > 0 {
        vals = append(vals, uint32(x))
    }
    if DoY & flags > 0 {
        vals = append(vals, uint32(y))
    }
    if DoW & flags > 0 {
        vals = append(vals, uint32(w))
    }
    if DoH & flags > 0 {
        vals = append(vals, uint32(h))
    }
    if DoSibling & flags > 0 {
        vals = append(vals, uint32(sibling))
    }
    if DoStack & flags > 0 {
        vals = append(vals, uint32(stackMode))
    }

    X.Conn().ConfigureWindow(window, flags, vals)
}

// configureRequest responds to generic configure requests from windows that
// we don't manage.
func configureRequest(X *xgbutil.XUtil, ev xevent.ConfigureRequestEvent) {
    configure(ev.Window, ev.ValueMask, ev.X, ev.Y, ev.Width, ev.Height,
              ev.Sibling, ev.StackMode)
}
