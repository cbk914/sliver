// Package device contains device emulation definitions for use with chromedp's
// Emulate action.
//
// See: https://raw.githubusercontent.com/puppeteer/puppeteer/main/src/common/DeviceDescriptors.ts
package device

// Generated by gen.go. DO NOT EDIT.

//go:generate go run gen.go

// Info holds device information for use with chromedp.Emulate.
type Info struct {
	// Name is the device name.
	Name string

	// UserAgent is the device user agent string.
	UserAgent string

	// Width is the viewport width.
	Width int64

	// Height is the viewport height.
	Height int64

	// Scale is the device viewport scale factor.
	Scale float64

	// Landscape indicates whether or not the device is in landscape mode or
	// not.
	Landscape bool

	// Mobile indicates whether it is a mobile device or not.
	Mobile bool

	// Touch indicates whether the device has touch enabled.
	Touch bool
}

// String satisfies fmt.Stringer.
func (i Info) String() string {
	return i.Name
}

// Device satisfies chromedp.Device.
func (i Info) Device() Info {
	return i
}

// infoType provides the enumerated device type.
type infoType int

// String satisfies fmt.Stringer.
func (i infoType) String() string {
	return devices[i].String()
}

// Device satisfies chromedp.Device.
func (i infoType) Device() Info {
	return devices[i]
}

// Devices.
const (
	// Reset is the reset device.
	Reset infoType = iota

	// BlackberryPlayBook is the "Blackberry PlayBook" device.
	BlackberryPlayBook

	// BlackberryPlayBooklandscape is the "Blackberry PlayBook landscape" device.
	BlackberryPlayBooklandscape

	// BlackBerryZ30 is the "BlackBerry Z30" device.
	BlackBerryZ30

	// BlackBerryZ30landscape is the "BlackBerry Z30 landscape" device.
	BlackBerryZ30landscape

	// GalaxyNote3 is the "Galaxy Note 3" device.
	GalaxyNote3

	// GalaxyNote3landscape is the "Galaxy Note 3 landscape" device.
	GalaxyNote3landscape

	// GalaxyNoteII is the "Galaxy Note II" device.
	GalaxyNoteII

	// GalaxyNoteIIlandscape is the "Galaxy Note II landscape" device.
	GalaxyNoteIIlandscape

	// GalaxySIII is the "Galaxy S III" device.
	GalaxySIII

	// GalaxySIIIlandscape is the "Galaxy S III landscape" device.
	GalaxySIIIlandscape

	// GalaxyS5 is the "Galaxy S5" device.
	GalaxyS5

	// GalaxyS5landscape is the "Galaxy S5 landscape" device.
	GalaxyS5landscape

	// GalaxyS8 is the "Galaxy S8" device.
	GalaxyS8

	// GalaxyS8landscape is the "Galaxy S8 landscape" device.
	GalaxyS8landscape

	// GalaxyS9 is the "Galaxy S9+" device.
	GalaxyS9

	// GalaxyS9landscape is the "Galaxy S9+ landscape" device.
	GalaxyS9landscape

	// GalaxyTabS4 is the "Galaxy Tab S4" device.
	GalaxyTabS4

	// GalaxyTabS4landscape is the "Galaxy Tab S4 landscape" device.
	GalaxyTabS4landscape

	// IPad is the "iPad" device.
	IPad

	// IPadlandscape is the "iPad landscape" device.
	IPadlandscape

	// IPadgen6 is the "iPad (gen 6)" device.
	IPadgen6

	// IPadgen6landscape is the "iPad (gen 6) landscape" device.
	IPadgen6landscape

	// IPadgen7 is the "iPad (gen 7)" device.
	IPadgen7

	// IPadgen7landscape is the "iPad (gen 7) landscape" device.
	IPadgen7landscape

	// IPadMini is the "iPad Mini" device.
	IPadMini

	// IPadMinilandscape is the "iPad Mini landscape" device.
	IPadMinilandscape

	// IPadPro is the "iPad Pro" device.
	IPadPro

	// IPadProlandscape is the "iPad Pro landscape" device.
	IPadProlandscape

	// IPadPro11 is the "iPad Pro 11" device.
	IPadPro11

	// IPadPro11landscape is the "iPad Pro 11 landscape" device.
	IPadPro11landscape

	// IPhone4 is the "iPhone 4" device.
	IPhone4

	// IPhone4landscape is the "iPhone 4 landscape" device.
	IPhone4landscape

	// IPhone5 is the "iPhone 5" device.
	IPhone5

	// IPhone5landscape is the "iPhone 5 landscape" device.
	IPhone5landscape

	// IPhone6 is the "iPhone 6" device.
	IPhone6

	// IPhone6landscape is the "iPhone 6 landscape" device.
	IPhone6landscape

	// IPhone6Plus is the "iPhone 6 Plus" device.
	IPhone6Plus

	// IPhone6Pluslandscape is the "iPhone 6 Plus landscape" device.
	IPhone6Pluslandscape

	// IPhone7 is the "iPhone 7" device.
	IPhone7

	// IPhone7landscape is the "iPhone 7 landscape" device.
	IPhone7landscape

	// IPhone7Plus is the "iPhone 7 Plus" device.
	IPhone7Plus

	// IPhone7Pluslandscape is the "iPhone 7 Plus landscape" device.
	IPhone7Pluslandscape

	// IPhone8 is the "iPhone 8" device.
	IPhone8

	// IPhone8landscape is the "iPhone 8 landscape" device.
	IPhone8landscape

	// IPhone8Plus is the "iPhone 8 Plus" device.
	IPhone8Plus

	// IPhone8Pluslandscape is the "iPhone 8 Plus landscape" device.
	IPhone8Pluslandscape

	// IPhoneSE is the "iPhone SE" device.
	IPhoneSE

	// IPhoneSElandscape is the "iPhone SE landscape" device.
	IPhoneSElandscape

	// IPhoneX is the "iPhone X" device.
	IPhoneX

	// IPhoneXlandscape is the "iPhone X landscape" device.
	IPhoneXlandscape

	// IPhoneXR is the "iPhone XR" device.
	IPhoneXR

	// IPhoneXRlandscape is the "iPhone XR landscape" device.
	IPhoneXRlandscape

	// IPhone11 is the "iPhone 11" device.
	IPhone11

	// IPhone11landscape is the "iPhone 11 landscape" device.
	IPhone11landscape

	// IPhone11Pro is the "iPhone 11 Pro" device.
	IPhone11Pro

	// IPhone11Prolandscape is the "iPhone 11 Pro landscape" device.
	IPhone11Prolandscape

	// IPhone11ProMax is the "iPhone 11 Pro Max" device.
	IPhone11ProMax

	// IPhone11ProMaxlandscape is the "iPhone 11 Pro Max landscape" device.
	IPhone11ProMaxlandscape

	// IPhone12 is the "iPhone 12" device.
	IPhone12

	// IPhone12landscape is the "iPhone 12 landscape" device.
	IPhone12landscape

	// IPhone12Pro is the "iPhone 12 Pro" device.
	IPhone12Pro

	// IPhone12Prolandscape is the "iPhone 12 Pro landscape" device.
	IPhone12Prolandscape

	// IPhone12ProMax is the "iPhone 12 Pro Max" device.
	IPhone12ProMax

	// IPhone12ProMaxlandscape is the "iPhone 12 Pro Max landscape" device.
	IPhone12ProMaxlandscape

	// IPhone12Mini is the "iPhone 12 Mini" device.
	IPhone12Mini

	// IPhone12Minilandscape is the "iPhone 12 Mini landscape" device.
	IPhone12Minilandscape

	// IPhone13 is the "iPhone 13" device.
	IPhone13

	// IPhone13landscape is the "iPhone 13 landscape" device.
	IPhone13landscape

	// IPhone13Pro is the "iPhone 13 Pro" device.
	IPhone13Pro

	// IPhone13Prolandscape is the "iPhone 13 Pro landscape" device.
	IPhone13Prolandscape

	// IPhone13ProMax is the "iPhone 13 Pro Max" device.
	IPhone13ProMax

	// IPhone13ProMaxlandscape is the "iPhone 13 Pro Max landscape" device.
	IPhone13ProMaxlandscape

	// IPhone13Mini is the "iPhone 13 Mini" device.
	IPhone13Mini

	// IPhone13Minilandscape is the "iPhone 13 Mini landscape" device.
	IPhone13Minilandscape

	// JioPhone2 is the "JioPhone 2" device.
	JioPhone2

	// JioPhone2landscape is the "JioPhone 2 landscape" device.
	JioPhone2landscape

	// KindleFireHDX is the "Kindle Fire HDX" device.
	KindleFireHDX

	// KindleFireHDXlandscape is the "Kindle Fire HDX landscape" device.
	KindleFireHDXlandscape

	// LGOptimusL70 is the "LG Optimus L70" device.
	LGOptimusL70

	// LGOptimusL70landscape is the "LG Optimus L70 landscape" device.
	LGOptimusL70landscape

	// MicrosoftLumia550 is the "Microsoft Lumia 550" device.
	MicrosoftLumia550

	// MicrosoftLumia950 is the "Microsoft Lumia 950" device.
	MicrosoftLumia950

	// MicrosoftLumia950landscape is the "Microsoft Lumia 950 landscape" device.
	MicrosoftLumia950landscape

	// Nexus10 is the "Nexus 10" device.
	Nexus10

	// Nexus10landscape is the "Nexus 10 landscape" device.
	Nexus10landscape

	// Nexus4 is the "Nexus 4" device.
	Nexus4

	// Nexus4landscape is the "Nexus 4 landscape" device.
	Nexus4landscape

	// Nexus5 is the "Nexus 5" device.
	Nexus5

	// Nexus5landscape is the "Nexus 5 landscape" device.
	Nexus5landscape

	// Nexus5X is the "Nexus 5X" device.
	Nexus5X

	// Nexus5Xlandscape is the "Nexus 5X landscape" device.
	Nexus5Xlandscape

	// Nexus6 is the "Nexus 6" device.
	Nexus6

	// Nexus6landscape is the "Nexus 6 landscape" device.
	Nexus6landscape

	// Nexus6P is the "Nexus 6P" device.
	Nexus6P

	// Nexus6Plandscape is the "Nexus 6P landscape" device.
	Nexus6Plandscape

	// Nexus7 is the "Nexus 7" device.
	Nexus7

	// Nexus7landscape is the "Nexus 7 landscape" device.
	Nexus7landscape

	// NokiaLumia520 is the "Nokia Lumia 520" device.
	NokiaLumia520

	// NokiaLumia520landscape is the "Nokia Lumia 520 landscape" device.
	NokiaLumia520landscape

	// NokiaN9 is the "Nokia N9" device.
	NokiaN9

	// NokiaN9landscape is the "Nokia N9 landscape" device.
	NokiaN9landscape

	// Pixel2 is the "Pixel 2" device.
	Pixel2

	// Pixel2landscape is the "Pixel 2 landscape" device.
	Pixel2landscape

	// Pixel2XL is the "Pixel 2 XL" device.
	Pixel2XL

	// Pixel2XLlandscape is the "Pixel 2 XL landscape" device.
	Pixel2XLlandscape

	// Pixel3 is the "Pixel 3" device.
	Pixel3

	// Pixel3landscape is the "Pixel 3 landscape" device.
	Pixel3landscape

	// Pixel4 is the "Pixel 4" device.
	Pixel4

	// Pixel4landscape is the "Pixel 4 landscape" device.
	Pixel4landscape

	// Pixel4a5G is the "Pixel 4a (5G)" device.
	Pixel4a5G

	// Pixel4a5Glandscape is the "Pixel 4a (5G) landscape" device.
	Pixel4a5Glandscape

	// Pixel5 is the "Pixel 5" device.
	Pixel5

	// Pixel5landscape is the "Pixel 5 landscape" device.
	Pixel5landscape

	// MotoG4 is the "Moto G4" device.
	MotoG4

	// MotoG4landscape is the "Moto G4 landscape" device.
	MotoG4landscape
)

// devices is the list of devices.
var devices = [...]Info{
	{"", "", 0, 0, 0.000000, false, false, false},
	{"Blackberry PlayBook", "Mozilla/5.0 (PlayBook; U; RIM Tablet OS 2.1.0; en-US) AppleWebKit/536.2+ (KHTML like Gecko) Version/7.2.1.0 Safari/536.2+", 600, 1024, 1.000000, false, true, true},
	{"Blackberry PlayBook landscape", "Mozilla/5.0 (PlayBook; U; RIM Tablet OS 2.1.0; en-US) AppleWebKit/536.2+ (KHTML like Gecko) Version/7.2.1.0 Safari/536.2+", 1024, 600, 1.000000, true, true, true},
	{"BlackBerry Z30", "Mozilla/5.0 (BB10; Touch) AppleWebKit/537.10+ (KHTML, like Gecko) Version/10.0.9.2372 Mobile Safari/537.10+", 360, 640, 2.000000, false, true, true},
	{"BlackBerry Z30 landscape", "Mozilla/5.0 (BB10; Touch) AppleWebKit/537.10+ (KHTML, like Gecko) Version/10.0.9.2372 Mobile Safari/537.10+", 640, 360, 2.000000, true, true, true},
	{"Galaxy Note 3", "Mozilla/5.0 (Linux; U; Android 4.3; en-us; SM-N900T Build/JSS15J) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30", 360, 640, 3.000000, false, true, true},
	{"Galaxy Note 3 landscape", "Mozilla/5.0 (Linux; U; Android 4.3; en-us; SM-N900T Build/JSS15J) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30", 640, 360, 3.000000, true, true, true},
	{"Galaxy Note II", "Mozilla/5.0 (Linux; U; Android 4.1; en-us; GT-N7100 Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30", 360, 640, 2.000000, false, true, true},
	{"Galaxy Note II landscape", "Mozilla/5.0 (Linux; U; Android 4.1; en-us; GT-N7100 Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30", 640, 360, 2.000000, true, true, true},
	{"Galaxy S III", "Mozilla/5.0 (Linux; U; Android 4.0; en-us; GT-I9300 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30", 360, 640, 2.000000, false, true, true},
	{"Galaxy S III landscape", "Mozilla/5.0 (Linux; U; Android 4.0; en-us; GT-I9300 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30", 640, 360, 2.000000, true, true, true},
	{"Galaxy S5", "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36", 360, 640, 3.000000, false, true, true},
	{"Galaxy S5 landscape", "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36", 640, 360, 3.000000, true, true, true},
	{"Galaxy S8", "Mozilla/5.0 (Linux; Android 7.0; SM-G950U Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36", 360, 740, 3.000000, false, true, true},
	{"Galaxy S8 landscape", "Mozilla/5.0 (Linux; Android 7.0; SM-G950U Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36", 740, 360, 3.000000, true, true, true},
	{"Galaxy S9+", "Mozilla/5.0 (Linux; Android 8.0.0; SM-G965U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.111 Mobile Safari/537.36", 320, 658, 4.500000, false, true, true},
	{"Galaxy S9+ landscape", "Mozilla/5.0 (Linux; Android 8.0.0; SM-G965U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.111 Mobile Safari/537.36", 658, 320, 4.500000, true, true, true},
	{"Galaxy Tab S4", "Mozilla/5.0 (Linux; Android 8.1.0; SM-T837A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.80 Safari/537.36", 712, 1138, 2.250000, false, true, true},
	{"Galaxy Tab S4 landscape", "Mozilla/5.0 (Linux; Android 8.1.0; SM-T837A) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.80 Safari/537.36", 1138, 712, 2.250000, true, true, true},
	{"iPad", "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1", 768, 1024, 2.000000, false, true, true},
	{"iPad landscape", "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1", 1024, 768, 2.000000, true, true, true},
	{"iPad (gen 6)", "Mozilla/5.0 (iPad; CPU OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 768, 1024, 2.000000, false, true, true},
	{"iPad (gen 6) landscape", "Mozilla/5.0 (iPad; CPU OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 1024, 768, 2.000000, true, true, true},
	{"iPad (gen 7)", "Mozilla/5.0 (iPad; CPU OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 810, 1080, 2.000000, false, true, true},
	{"iPad (gen 7) landscape", "Mozilla/5.0 (iPad; CPU OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 1080, 810, 2.000000, true, true, true},
	{"iPad Mini", "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1", 768, 1024, 2.000000, false, true, true},
	{"iPad Mini landscape", "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1", 1024, 768, 2.000000, true, true, true},
	{"iPad Pro", "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1", 1024, 1366, 2.000000, false, true, true},
	{"iPad Pro landscape", "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1", 1366, 1024, 2.000000, true, true, true},
	{"iPad Pro 11", "Mozilla/5.0 (iPad; CPU OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 834, 1194, 2.000000, false, true, true},
	{"iPad Pro 11 landscape", "Mozilla/5.0 (iPad; CPU OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 1194, 834, 2.000000, true, true, true},
	{"iPhone 4", "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D257 Safari/9537.53", 320, 480, 2.000000, false, true, true},
	{"iPhone 4 landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D257 Safari/9537.53", 480, 320, 2.000000, true, true, true},
	{"iPhone 5", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1", 320, 568, 2.000000, false, true, true},
	{"iPhone 5 landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1", 568, 320, 2.000000, true, true, true},
	{"iPhone 6", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1", 375, 667, 2.000000, false, true, true},
	{"iPhone 6 landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1", 667, 375, 2.000000, true, true, true},
	{"iPhone 6 Plus", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1", 414, 736, 3.000000, false, true, true},
	{"iPhone 6 Plus landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1", 736, 414, 3.000000, true, true, true},
	{"iPhone 7", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1", 375, 667, 2.000000, false, true, true},
	{"iPhone 7 landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1", 667, 375, 2.000000, true, true, true},
	{"iPhone 7 Plus", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1", 414, 736, 3.000000, false, true, true},
	{"iPhone 7 Plus landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1", 736, 414, 3.000000, true, true, true},
	{"iPhone 8", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1", 375, 667, 2.000000, false, true, true},
	{"iPhone 8 landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1", 667, 375, 2.000000, true, true, true},
	{"iPhone 8 Plus", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1", 414, 736, 3.000000, false, true, true},
	{"iPhone 8 Plus landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1", 736, 414, 3.000000, true, true, true},
	{"iPhone SE", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1", 320, 568, 2.000000, false, true, true},
	{"iPhone SE landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1", 568, 320, 2.000000, true, true, true},
	{"iPhone X", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1", 375, 812, 3.000000, false, true, true},
	{"iPhone X landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1", 812, 375, 3.000000, true, true, true},
	{"iPhone XR", "Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Mobile/15E148 Safari/604.1", 414, 896, 3.000000, false, true, true},
	{"iPhone XR landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Mobile/15E148 Safari/604.1", 896, 414, 3.000000, true, true, true},
	{"iPhone 11", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1 Mobile/15E148 Safari/604.1", 414, 828, 2.000000, false, true, true},
	{"iPhone 11 landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1 Mobile/15E148 Safari/604.1", 828, 414, 2.000000, true, true, true},
	{"iPhone 11 Pro", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1 Mobile/15E148 Safari/604.1", 375, 812, 3.000000, false, true, true},
	{"iPhone 11 Pro landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1 Mobile/15E148 Safari/604.1", 812, 375, 3.000000, true, true, true},
	{"iPhone 11 Pro Max", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1 Mobile/15E148 Safari/604.1", 414, 896, 3.000000, false, true, true},
	{"iPhone 11 Pro Max landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1 Mobile/15E148 Safari/604.1", 896, 414, 3.000000, true, true, true},
	{"iPhone 12", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 390, 844, 3.000000, false, true, true},
	{"iPhone 12 landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 844, 390, 3.000000, true, true, true},
	{"iPhone 12 Pro", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 390, 844, 3.000000, false, true, true},
	{"iPhone 12 Pro landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 844, 390, 3.000000, true, true, true},
	{"iPhone 12 Pro Max", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 428, 926, 3.000000, false, true, true},
	{"iPhone 12 Pro Max landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 926, 428, 3.000000, true, true, true},
	{"iPhone 12 Mini", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 375, 812, 3.000000, false, true, true},
	{"iPhone 12 Mini landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 812, 375, 3.000000, true, true, true},
	{"iPhone 13", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 390, 844, 3.000000, false, true, true},
	{"iPhone 13 landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 844, 390, 3.000000, true, true, true},
	{"iPhone 13 Pro", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 390, 844, 3.000000, false, true, true},
	{"iPhone 13 Pro landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 844, 390, 3.000000, true, true, true},
	{"iPhone 13 Pro Max", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 428, 926, 3.000000, false, true, true},
	{"iPhone 13 Pro Max landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 926, 428, 3.000000, true, true, true},
	{"iPhone 13 Mini", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 375, 812, 3.000000, false, true, true},
	{"iPhone 13 Mini landscape", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1", 812, 375, 3.000000, true, true, true},
	{"JioPhone 2", "Mozilla/5.0 (Mobile; LYF/F300B/LYF-F300B-001-01-15-130718-i;Android; rv:48.0) Gecko/48.0 Firefox/48.0 KAIOS/2.5", 240, 320, 1.000000, false, true, true},
	{"JioPhone 2 landscape", "Mozilla/5.0 (Mobile; LYF/F300B/LYF-F300B-001-01-15-130718-i;Android; rv:48.0) Gecko/48.0 Firefox/48.0 KAIOS/2.5", 320, 240, 1.000000, true, true, true},
	{"Kindle Fire HDX", "Mozilla/5.0 (Linux; U; en-us; KFAPWI Build/JDQ39) AppleWebKit/535.19 (KHTML, like Gecko) Silk/3.13 Safari/535.19 Silk-Accelerated=true", 800, 1280, 2.000000, false, true, true},
	{"Kindle Fire HDX landscape", "Mozilla/5.0 (Linux; U; en-us; KFAPWI Build/JDQ39) AppleWebKit/535.19 (KHTML, like Gecko) Silk/3.13 Safari/535.19 Silk-Accelerated=true", 1280, 800, 2.000000, true, true, true},
	{"LG Optimus L70", "Mozilla/5.0 (Linux; U; Android 4.4.2; en-us; LGMS323 Build/KOT49I.MS32310c) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/75.0.3765.0 Mobile Safari/537.36", 384, 640, 1.250000, false, true, true},
	{"LG Optimus L70 landscape", "Mozilla/5.0 (Linux; U; Android 4.4.2; en-us; LGMS323 Build/KOT49I.MS32310c) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/75.0.3765.0 Mobile Safari/537.36", 640, 384, 1.250000, true, true, true},
	{"Microsoft Lumia 550", "Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 550) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/14.14263", 640, 360, 2.000000, false, true, true},
	{"Microsoft Lumia 950", "Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 950) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/14.14263", 360, 640, 4.000000, false, true, true},
	{"Microsoft Lumia 950 landscape", "Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 950) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/14.14263", 640, 360, 4.000000, true, true, true},
	{"Nexus 10", "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 10 Build/MOB31T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Safari/537.36", 800, 1280, 2.000000, false, true, true},
	{"Nexus 10 landscape", "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 10 Build/MOB31T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Safari/537.36", 1280, 800, 2.000000, true, true, true},
	{"Nexus 4", "Mozilla/5.0 (Linux; Android 4.4.2; Nexus 4 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36", 384, 640, 2.000000, false, true, true},
	{"Nexus 4 landscape", "Mozilla/5.0 (Linux; Android 4.4.2; Nexus 4 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36", 640, 384, 2.000000, true, true, true},
	{"Nexus 5", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36", 360, 640, 3.000000, false, true, true},
	{"Nexus 5 landscape", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36", 640, 360, 3.000000, true, true, true},
	{"Nexus 5X", "Mozilla/5.0 (Linux; Android 8.0.0; Nexus 5X Build/OPR4.170623.006) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36", 412, 732, 2.625000, false, true, true},
	{"Nexus 5X landscape", "Mozilla/5.0 (Linux; Android 8.0.0; Nexus 5X Build/OPR4.170623.006) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36", 732, 412, 2.625000, true, true, true},
	{"Nexus 6", "Mozilla/5.0 (Linux; Android 7.1.1; Nexus 6 Build/N6F26U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36", 412, 732, 3.500000, false, true, true},
	{"Nexus 6 landscape", "Mozilla/5.0 (Linux; Android 7.1.1; Nexus 6 Build/N6F26U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36", 732, 412, 3.500000, true, true, true},
	{"Nexus 6P", "Mozilla/5.0 (Linux; Android 8.0.0; Nexus 6P Build/OPP3.170518.006) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36", 412, 732, 3.500000, false, true, true},
	{"Nexus 6P landscape", "Mozilla/5.0 (Linux; Android 8.0.0; Nexus 6P Build/OPP3.170518.006) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36", 732, 412, 3.500000, true, true, true},
	{"Nexus 7", "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 7 Build/MOB30X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Safari/537.36", 600, 960, 2.000000, false, true, true},
	{"Nexus 7 landscape", "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 7 Build/MOB30X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Safari/537.36", 960, 600, 2.000000, true, true, true},
	{"Nokia Lumia 520", "Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; NOKIA; Lumia 520)", 320, 533, 1.500000, false, true, true},
	{"Nokia Lumia 520 landscape", "Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; NOKIA; Lumia 520)", 533, 320, 1.500000, true, true, true},
	{"Nokia N9", "Mozilla/5.0 (MeeGo; NokiaN9) AppleWebKit/534.13 (KHTML, like Gecko) NokiaBrowser/8.5.0 Mobile Safari/534.13", 480, 854, 1.000000, false, true, true},
	{"Nokia N9 landscape", "Mozilla/5.0 (MeeGo; NokiaN9) AppleWebKit/534.13 (KHTML, like Gecko) NokiaBrowser/8.5.0 Mobile Safari/534.13", 854, 480, 1.000000, true, true, true},
	{"Pixel 2", "Mozilla/5.0 (Linux; Android 8.0; Pixel 2 Build/OPD3.170816.012) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36", 411, 731, 2.625000, false, true, true},
	{"Pixel 2 landscape", "Mozilla/5.0 (Linux; Android 8.0; Pixel 2 Build/OPD3.170816.012) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36", 731, 411, 2.625000, true, true, true},
	{"Pixel 2 XL", "Mozilla/5.0 (Linux; Android 8.0.0; Pixel 2 XL Build/OPD1.170816.004) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36", 411, 823, 3.500000, false, true, true},
	{"Pixel 2 XL landscape", "Mozilla/5.0 (Linux; Android 8.0.0; Pixel 2 XL Build/OPD1.170816.004) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3765.0 Mobile Safari/537.36", 823, 411, 3.500000, true, true, true},
	{"Pixel 3", "Mozilla/5.0 (Linux; Android 9; Pixel 3 Build/PQ1A.181105.017.A1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.158 Mobile Safari/537.36", 393, 786, 2.750000, false, true, true},
	{"Pixel 3 landscape", "Mozilla/5.0 (Linux; Android 9; Pixel 3 Build/PQ1A.181105.017.A1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.158 Mobile Safari/537.36", 786, 393, 2.750000, true, true, true},
	{"Pixel 4", "Mozilla/5.0 (Linux; Android 10; Pixel 4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Mobile Safari/537.36", 353, 745, 3.000000, false, true, true},
	{"Pixel 4 landscape", "Mozilla/5.0 (Linux; Android 10; Pixel 4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Mobile Safari/537.36", 745, 353, 3.000000, true, true, true},
	{"Pixel 4a (5G)", "Mozilla/5.0 (Linux; Android 11; Pixel 4a (5G)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4812.0 Mobile Safari/537.36", 353, 745, 3.000000, false, true, true},
	{"Pixel 4a (5G) landscape", "Mozilla/5.0 (Linux; Android 11; Pixel 4a (5G)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4812.0 Mobile Safari/537.36", 745, 353, 3.000000, true, true, true},
	{"Pixel 5", "Mozilla/5.0 (Linux; Android 11; Pixel 5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4812.0 Mobile Safari/537.36", 393, 851, 3.000000, false, true, true},
	{"Pixel 5 landscape", "Mozilla/5.0 (Linux; Android 11; Pixel 5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4812.0 Mobile Safari/537.36", 851, 393, 3.000000, true, true, true},
	{"Moto G4", "Mozilla/5.0 (Linux; Android 7.0; Moto G (4)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4812.0 Mobile Safari/537.36", 360, 640, 3.000000, false, true, true},
	{"Moto G4 landscape", "Mozilla/5.0 (Linux; Android 7.0; Moto G (4)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4812.0 Mobile Safari/537.36", 640, 360, 3.000000, true, true, true},
}
