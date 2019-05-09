package main

import (
    "fmt"
    "github.com/gotk3/gotk3/cairo"
    "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gtk"
)

func main() {
    x := 0.0
    y := 0.0
    
    gtk.Init(nil)
    win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
    if err != nil {
        fmt.Println("Unable to create window:", err)
    }
    win.SetTitle("Go GTK3 Test")
    win.Connect("destroy", func() {
        gtk.MainQuit()
    })
    
    //~ lbl, err := gtk.LabelNew("Hello, gotk3!")
    //~ if err != nil {
        //~ fmt.Println("Unable to create label:", err)
    //~ }
    //~ win.Add(lbl)
    
    da, _ := gtk.DrawingAreaNew()
    win.Add(da)
    da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
        cr.SetSourceRGB(0.05, 0.2, 1.0)
        cr.Rectangle(x, y, 64, 64)
        cr.Fill()
        
        cr.CurveTo(0.0, 0.0, 50.0, 0.0, 100.0, 100.0)
        cr.CurveTo(100.0, 100.0, 150.0, 0.0, 200.0, 0.0)
        cr.CurveTo(200.0, 0.0, 250.0, 0.0, 300.0, 100.0)
        cr.Stroke()
    })
    
    glib.TimeoutAdd(100, func() bool {
        //~ fmt.Println("callback triggered")
        x++
        y++
        win.QueueDraw()
        return true
    }, nil)
    
    win.SetDefaultSize(800, 600)
    win.ShowAll()
    gtk.Main()
}
