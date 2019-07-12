package content

import (
	"container/list"
	"math"
	"math/rand"
	"sudare_contents/lib"
	"sudare_contents/lib/util"
	"time"

	"github.com/llgcode/draw2d/draw2dimg"
)

type ContentSinWideLine struct {
	data      *lib.Cylinder
	count     float64
	lineQueue *list.List
	start     time.Time
}

func NewContentSinWideLine() CylinderContent {
	c := new(ContentSinWideLine)
	c.data = lib.NewCylinder()
	c.lineQueue = list.New()
	return c
}

func (c *ContentSinWideLine) Begin() {
	c.start = time.Now()
}

func (c *ContentSinWideLine) GetFrame() []byte {
	c.data.Clear()

	duration := time.Now().Sub(c.start)
	if duration > 10*time.Millisecond {
		c.start = time.Now()
		c.count += 0.1
	}

	if rand.Intn(10) == 1 {
		c.lineQueue.PushBack(&lineData{
			y:      0,
			count:  0,
			len:    lib.CylinderRadius/2 + rand.Float64()*lib.CylinderRadius/2,
			height: rand.Float64() * 15,
			yspeed: rand.Float64()*2 + 0.5,
			cspeed: rand.Float64()*0.5 + 0.1,
			colorh: rand.Float64(),
		})
	}

	c.data.Render(func(i int, gc *draw2dimg.GraphicContext) {

		for e := c.lineQueue.Front(); e != nil; e = e.Next() {

			lined := e.Value.(*lineData)
			unit := (2 * math.Pi) / lib.CylinderCount

			wave := math.Sin(float64(i)*unit+lined.count) * lined.height

			//			color := util.GetRainbow((colordepth + 1) / 2)
			color := util.GetRainbow(lined.colorh)
			gc.SetStrokeColor(color)
			gc.SetLineWidth(5)

			gc.MoveTo(15, lined.y-wave) // should always be called first for a new path
			gc.LineTo(15, lined.y+wave)
			gc.Close()
			gc.FillStroke()

		}

	})
	for e := c.lineQueue.Front(); e != nil; e = e.Next() {
		lined := e.Value.(*lineData)
		if lined.y >= lib.CylinderHeight {
			c.lineQueue.Remove(e)
		}
		lined.y += lined.yspeed
		lined.count += lined.cspeed
	}
	return c.data.GetData()
}
