package content

import (
	"math"
	"sudare_contents/lib"
	"sudare_contents/lib/util"

	"github.com/llgcode/draw2d/draw2dimg"
)

type ContentExLine struct {
	data  *lib.Cylinder
	count float64
}

func NewContentExLine() CylinderContent {
	c := new(ContentExLine)
	c.data = lib.NewCylinder()
	return c
}

func (c *ContentExLine) Begin() {

}

func (c *ContentExLine) GetFrame() []byte {
	c.count += 0.1
	c.data.Render(func(i int, gc *draw2dimg.GraphicContext) {

		unit := (2 * math.Pi) / lib.CylinderCount
		center := float64(lib.CylinderHeight / 2)
		sin := math.Sin((float64(i)*unit + c.count) * 1.2)
		sin2 := math.Sin(-(float64(i)*unit + c.count) * 2)

		wave := sin * 20

		color := util.GetRainbow(math.Abs(sin2))

		gc.SetStrokeColor(color)
		gc.SetLineWidth(1)

		gc.MoveTo(0, center) // should always be called first for a new path
		gc.LineTo(15*math.Abs(sin2)+7, center+wave)
		gc.Close()
		gc.FillStroke()
	})
	return c.data.GetData()
}
