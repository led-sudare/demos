package main

import (
	"bufio"
	"demos/lib/util"
	"demos/lib/webapi"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	log "github.com/cihub/seelog"
)

type Configs struct {
	XProxySubBind  string `json:"XProxySubBind"`
	AdapterSubBind string `json:"AdapterSubBind"`
}

func NewConfigs() Configs {
	return Configs{
		XProxySubBind:  "0.0.0.0:5510",
		AdapterSubBind: "0.0.0.0:5520",
	}
}

type Demo struct {
	cmd  *exec.Cmd
	name string
	args []string
}

func NewDemo(name string, args []string) *Demo {
	demo := new(Demo)

	demo.name = name
	demo.args = args

	return demo
}

func (d *Demo) SetUp() {
	d.cmd = exec.Command(d.name, d.args...)
}

func (d *Demo) Run() error {
	stdout, err := d.cmd.StdoutPipe()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	stderr, err := d.cmd.StderrPipe()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	err = d.cmd.Start()
	if err != nil {
		return err
	}

	scanStdout := bufio.NewScanner(stdout)
	scanStderr := bufio.NewScanner(stderr)
	for scanStdout.Scan() || scanStderr.Scan() {

		outTest := scanStdout.Text()
		errText := scanStderr.Text()
		if len(outTest) > 0 {
			fmt.Println(outTest)
		}
		if len(errText) > 0 {
			fmt.Println(errText)
		}
	}

	return d.cmd.Wait()
}

type Demos struct {
	demos []*Demo
	abort chan struct{}
}

func NewDemos(demos []*Demo) *Demos {
	d := new(Demos)
	d.demos = demos
	d.abort = make(chan struct{})
	return d
}

func Aborter(d *Demo, enable <-chan bool, finish chan struct{}) error {

	for {
		select {
		case e, ok := <-enable:
			if e {
				continue
			}
			log.Infof("Kill Process.. %v", d.name)
			err := d.cmd.Process.Kill()
			if err != nil {
				return err
			}
			if ok {
				return nil
			}
			return errors.New("Abort")
		case <-finish:
			return nil
		}
	}
}

func (d *Demos) RunDemos(enable <-chan bool) {

	for {
		for _, demo := range d.demos {
			log.Infof("Process Start: %v %v", demo.name, demo.args)
			finish := make(chan struct{})
			demo.SetUp()
			go Aborter(demo, enable, finish)
			err := demo.Run()
			close(finish)
			if err != nil {
				log.Info("RunDemos: ", err)
				return
			}
		}
	}
}

func doDo(demos *Demos, enable <-chan bool) {
	e := true
	for {
		if e {
			demos.RunDemos(enable)
		}
		log.Info("Demo Abort.")
		e = <-enable
		log.Info("Enable: ", e)
	}
}

type WebAPICtrlImpl struct {
	enable   chan bool
	isEnable bool
}

func NewWebAPICtrlImpl(enable chan bool) *WebAPICtrlImpl {
	w := new(WebAPICtrlImpl)
	w.isEnable = true
	w.enable = enable
	return w
}

func (w *WebAPICtrlImpl) Enable(enable bool) {
	log.Info("WebAPICtrlImpl Enable: ", enable)
	w.isEnable = enable

	select {
	case w.enable <- enable:
		break
	default:
		<-w.enable
		w.enable <- enable
	}
}
func (w *WebAPICtrlImpl) IsEnable() bool {
	return w.isEnable
}

func main() {
	configs := NewConfigs()
	util.ReadConfig(&configs)

	var (
		proxyAddr   = flag.String("p", configs.XProxySubBind, "Specify IP and port of XSUB-XPUB server.")
		adapterAddr = flag.String("a", configs.AdapterSubBind, "Specify IP and port of Adapter server.")
	)

	flag.Parse()
	enable := make(chan bool)
	controller := NewWebAPICtrlImpl(enable)

	demos := NewDemos([]*Demo{
		NewDemo("./moridemo/moridemo", strings.Split(*adapterAddr, ":")),
		NewDemo("./sudare_contents/sudare_contents", []string{"-r", *proxyAddr}),
	})

	go doDo(demos, enable)
	webapi.SetUpWebAPIforCommon(controller)

	log.Info("Http Server 5003 Start")
	http.ListenAndServe(":5003", nil)

}
