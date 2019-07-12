package content

import (
	"math"
	"sudare_contents/lib"
	"sudare_contents/lib/util"
	"time"
)

type ContentCirWave struct {
	data  *lib.Cylinder
	count float64
	start time.Time
}

func NewContentCirWave() CylinderContent {
	c := new(ContentCirWave)
	c.data = lib.NewCylinder()
	return c
}

func (c *ContentCirWave) Begin() {
	c.start = time.Now()
}

func (c *ContentCirWave) GetFrame() []byte {
	c.data.Clear()

	duration := time.Now().Sub(c.start)
	if duration > 10*time.Millisecond {
		c.start = time.Now()
		c.count += 0.1
	}

	c.data.RenderEachCylinder(func(i int) {

		wavedis := []float64{5, 10, 2}
		waveh := []float64{1, 0.5, 0.8}

		for y := 0; y < lib.CylinderHeight; y++ {
			for wi, wd := range wavedis {
				colordepth := c.count + float64(y)/wd
				colordepth = (math.Sin(colordepth) + 1) / 2
				waveheight := colordepth * lib.CylinderRadius * waveh[wi]

				color := util.GetRainbow(colordepth)
				c.data.SetAt(int(waveheight), y, i, color)
			}
		}
	})

	return c.data.GetData()
}
