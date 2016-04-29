package x11ui

import (
	// "fmt"
	"image"
	"image/color"
	"log"

	"github.com/llgcode/draw2d/draw2dkit"

	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/lucasb-eyer/go-colorful"
)

type TextSettingBox struct {
	p  *Window //ParentWindow
	me *Window
	// Custom properties
	bgColor   color.Color
	txtColor  color.Color
	fmtString string
	textSize  float64
}

func NewTextSettingBox(title string, p *Window, dims ...int) *TextSettingBox {
	if p == nil {
		log.Fatal("Cannot Create Widget without Application")
	}
	pbox := new(TextSettingBox)
	pbox.me = newWindow(p.Window.X, p, title, dims...)

	// pbox.SetFmtString(title)
	pbox.ResetFmtString()
	pbox.loadTheme()
	pbox.drawBackground()
	// pbar.SetValue(0.5)
	return pbox
}
func (p *TextSettingBox) loadTheme() {
	p.bgColor = colorful.LinearRgb(.4, .6, .1)
	p.txtColor = color.RGBA{20, 200, 30, 200}
}
func (p *TextSettingBox) SetBgColor(bc color.Color) {
	p.bgColor = bc
}
func (p *TextSettingBox) SetTextColor(tc color.Color) {
	p.txtColor = tc
}
func (p *TextSettingBox) SetTextSize(ts float64) {
	p.textSize = ts
}

func (p *TextSettingBox) SetFmtString(s string) {
	p.fmtString = s
	r := p.me.Rect
	r.MoveTo(0, 0)

	dest := image.NewRGBA(r.ImageRect())

	gc := draw2dimg.NewGraphicContext(dest)

	// Set text color
	gc.SetFillColor(p.txtColor)
	gc.SetFontSize(p.textSize)
	gc.SetLineWidth(1)
	cx, cy := r.Center()
	str := p.fmtString
	// fmt.Sprintf(p.fmtString, (p.val * 100.0))
	x0, y0, w0, h0 := gc.GetStringBounds(str)
	// log.Println("Required dimension for string ", x0, y0, w0, h0, cx, cy)
	tx, ty := float64(cx)-w0/2.0-x0, float64(cy)-h0/2.0-y0/2.0
	// gc.StrokeStringAt(str, tx, ty)
	gc.FillStringAt(str, tx, ty)
	// gc.FillStroke()
	gc.Close()

	g := xgraphics.NewConvert(p.me.X(), dest)

	// w.drawLabel(g, w.title)
	g.XSurfaceSet(p.me.Id)
	g.XDraw()
	g.XPaintRects(p.me.Id, r.ImageRect())
}

func (p *TextSettingBox) FmtString() string {
	return p.fmtString
}
func (p *TextSettingBox) ResetFmtString() {
	p.fmtString = "%4.2f%%"
}

func (p *TextSettingBox) drawBackground() {
	r := p.me.Rect
	r.MoveTo(0, 0)

	dest := image.NewRGBA(r.ImageRect())

	gc := draw2dimg.NewGraphicContext(dest)

	// bg := colorful.LinearRgb(.025, .025, .025)
	// switch s {
	// case StateNormal, StateReleased:

	// gc.SetFillColor(color.RGBA{0x20, 0x20, 0x20, 20})
	gc.SetFillColor(p.bgColor)
	log.Println(p.bgColor)
	// gc.SetStrokeColor(systemFG)
	gc.SetStrokeColor(p.bgColor)
	// case StateHovered:
	// gc.SetFillColor(color.RGsBA{0x65, 0x20, 0x20, 20})
	// 	gc.SetStrokeColor(systemFG)
	// case StatePressed:
	// 	gc.SetFillColor(color.RGBA{0x20, 0x30, 0x20, 20})
	// 	gc.SetStrokeColor(systemFG)
	// case StateSpecial:
	// 	gc.SetFillColor(color.RGBA{0x20, 0x80, 0x20, 0x80})
	// 	gc.SetStrokeColor(systemFG)
	// }

	// // gc.SetLineJoin(draw2d.RoundJoin)
	// // gc.Rotate(math.Pi / 4.0)
	WW := float64(r.Width)
	HH := float64(r.Height)

	// Draw Background
	margin := 2.0
	gc.SetLineWidth(1)
	draw2dkit.Rectangle(gc, margin, margin, WW-margin, HH-margin)
	gc.FillStroke()

	gc.Close()

	g := xgraphics.NewConvert(p.me.X(), dest)

	// w.drawLabel(g, w.title)
	g.XSurfaceSet(p.me.Id)
	g.XDraw()
	g.XPaintRects(p.me.Id, r.ImageRect())
	// return g
	// return g
}
