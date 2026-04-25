package main

type Device interface {
	On()
	Off()
	SetVolume(int)
}

type TV struct{}
type Radio struct{}

type Remote struct {
	device Device // bridge
}

func (r *Remote) TogglePower() {
	r.device.On()
}

type AdvancedRemote struct {
	Remote
}

func (a *AdvancedRemote) Mute() {
	a.device.SetVolume(0)
}

// Remote ---------> Device(interface)
//    |                 |
//    |                 ├── TV
//    |                 ├── Radio
//    |                 └── Speaker

// ----> is bridge

//AdvancedRemote is not the bridge itself, but it is part of the Bridge Pattern.
// AdvancedRemote is a refined abstraction in the Bridge Pattern, not the bridge itself.
// It inherits and uses the existing bridge (device Device) from Remote.

// Remote = Abstraction
// AdvancedRemote = Refined Abstraction
// Device = Implementor
// TV, Radio = Concrete Implementors
// device Device field = the bridge/reference/composition link

// Device
// ├── TV
// ├── Radio
// └── Speaker

// TV + Radio + Speaker --> Device = hierarchy/family.
