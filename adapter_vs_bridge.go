package main

// =========================
// Adapter
// =========================

package main

import "fmt"

// interface yang dibutuhkan client
type Printer interface {
	Print()
}

// kode lama
type OldPrinter struct{}

func (o *OldPrinter) PrintOld() {
	fmt.Println("Printing from old printer")
}

// adapter
type PrinterAdapter struct {
	oldPrinter *OldPrinter
}

func (p *PrinterAdapter) Print() {
	p.oldPrinter.PrintOld()
}

func main() {
	old := &OldPrinter{}
	adapter := &PrinterAdapter{oldPrinter: old}

	var printer Printer = adapter
	printer.Print()
}

// =========================
// Bridge
// =========================

package main

import "fmt"

// implementation
type Device interface {
	On()
	Off()
}

type TV struct{}

func (t *TV) On() {
	fmt.Println("TV ON")
}

func (t *TV) Off() {
	fmt.Println("TV OFF")
}

type Radio struct{}

func (r *Radio) On() {
	fmt.Println("Radio ON")
}

func (r *Radio) Off() {
	fmt.Println("Radio OFF")
}

// abstraction
type Remote struct {
	device Device
}

func (r *Remote) PowerOn() {
	r.device.On()
}

func (r *Remote) PowerOff() {
	r.device.Off()
}

// refined abstraction
type AdvancedRemote struct {
	Remote
}

func (a *AdvancedRemote) Mute() {
	fmt.Println("Volume muted")
}

func main() {
	tv := &TV{}
	radio := &Radio{}

	remoteTV := Remote{device: tv}
	remoteTV.PowerOn()

	remoteRadio := Remote{device: radio}
	remoteRadio.PowerOn()

	adv := AdvancedRemote{
		Remote{device: tv},
	}
	adv.Mute()
}