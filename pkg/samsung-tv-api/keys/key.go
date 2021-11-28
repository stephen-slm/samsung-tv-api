package keys

//goland:noinspection ALL

const (
	//
	//  POWER
	//

	PowerOff = "KEY_POWEROFF"
	PowerOn  = "KEY_POWERON"
	Power    = "KEY_POWER"

	//
	// Input Keys
	//

	Source      = "KEY_SOURCE"
	Component1  = "KEY_COMPONENT1"
	Component2  = "KEY_COMPONENT2"
	AV1         = "KEY_AV1"
	AV2         = "KEY_AV2"
	AV3         = "KEY_AV3"
	SVideo1     = "KEY_SVIDEO1"
	SVideo2     = "KEY_SVIDEO2"
	SVideo3     = "KEY_SVIDEO3"
	HDMI        = "KEY_HDMI"
	FMRadio     = "KEY_FM_RADIO"
	DVI         = "KEY_DVI"
	DVR         = "KEY_DVR"
	TV          = "KEY_TV"
	AnalogTV    = "KEY_ANTENA"
	DigitalTV   = "KEY_DTV"
	AmbientMode = "KEY_AMBIENT"

	//
	// Number Keys
	//

	Key1 = "KEY_1"
	Key2 = "KEY_2"
	Key3 = "KEY_3"
	Key4 = "KEY_4"
	Key5 = "KEY_5"
	Key6 = "KEY_6"
	Key7 = "KEY_7"
	Key8 = "KEY_8"
	Key9 = "KEY_9"
	Key0 = "KEY_0"

	//
	// Misc Keys
	//

	ThreeD       = "KEY_PANNEL_CHDOWN"
	AnyNetPlus   = "KEY_ANYNET"
	EnergySaving = "KEY_ESAVING"
	SleepTimer   = "KEY_SLEEP"
	DTVSignal    = "KEY_DTV_SIGNAL"

	//
	// Channel Keys
	//

	ChannelUp        = "KEY_CHUP"
	ChannelDown      = "KEY_CHDOWN"
	PreviousChannel  = "KEY_PRECH"
	FavoriteChannels = "KEY_FAVCH"
	ChannelList      = "KEY_CH_LIST"
	AutoProgram      = "KEY_AUTO_PROGRAM"
	MagicChannel     = "KEY_MAGIC_CHANNEL"

	//
	// Volume Keys
	//

	VolumeUp   = "KEY_VOLUP"
	VolumeDown = "KEY_VOLDOWN"
	Mute       = "KEY_MUTE"

	//
	// Direction Keys
	//

	NavigationUp     = "KEY_UP"
	NavigationDown   = "KEY_DOWN"
	NavigationLeft   = "KEY_LEFT"
	NavigationRight  = "KEY_RIGHT"
	NavigationReturn = "KEY_RETURN"
	NavigationEnter  = "KEY_ENTER"

	//
	// Media Keys
	//

	Rewind               = "KEY_REWIND"
	Stop                 = "KEY_STOP"
	Play                 = "KEY_PLAY"
	FastForward          = "KEY_FF"
	Record               = "KEY_REC"
	Pause                = "KEY_PAUSE"
	Live                 = "KEY_LIVE"
	fnKEY_QUICK_REPLAY   = "KEY_QUICK_REPLAY"
	fnKEY_STILL_PICTURE  = "KEY_STILL_PICTURE"
	fnKEY_INSTANT_REPLAY = "KEY_INSTANT_REPLAY"

	//
	// Picture in Picture Keys
	//

	PIPOnOff        = "KEY_PIP_ONOFF"
	PIPSwap         = "KEY_PIP_SWAP"
	PIPSize         = "KEY_PIP_SIZE"
	PIPChannelUp    = "KEY_PIP_CHUP"
	PIPChannelDown  = "KEY_PIP_CHDOWN"
	PIPSmall        = "KEY_AUTO_ARC_PIP_SMALL"
	PIPWide         = "KEY_AUTO_ARC_PIP_WIDE"
	PIPBottomRight  = "KEY_AUTO_ARC_PIP_RIGHT_BOTTOM"
	PIPSourceChange = "KEY_AUTO_ARC_PIP_SOURCE_CHANGE"
	PIPScan         = "KEY_PIP_SCAN"

	//
	// Modes Keys
	//

	VCRMode  = "KEY_VCR_MODE"
	CATVMode = "KEY_CATV_MODE"
	DSSMode  = "KEY_DSS_MODE"
	TVMode   = "KEY_TV_MODE"
	DVDMode  = "KEY_DVD_MODE"
	STBMode  = "KEY_STB_MODE"
	PCMode   = "KEY_PCMODE"

	//
	// Color Keys
	//

	Green  = "KEY_GREEN"
	Yellow = "KEY_YELLOW"
	Cyan   = "KEY_CYAN"
	Red    = "KEY_RED"

	//
	// Teletext
	//

	TeletextMix     = "KEY_TTX_MIX"
	TeletextSubface = "KEY_TTX_SUBFACE"

	//
	// Aspect Ratio Keys
	//

	AspectRatio       = "KEY_ASPECT"
	PictureSize       = "KEY_PICTURE_SIZE"
	AspectRatio43     = "KEY_4_3"
	AspectRatio169    = "KEY_16_9"
	AspectRatio34Alt  = "KEY_EXT14"
	AspectRatio169Alt = "KEY_EXT15"

	//
	// Picture Mode Keys
	//

	KEY_PMODE    = "PictureMode"
	KEY_PANORAMA = "PictureModePanorama"
	KEY_DYNAMIC  = "PictureModeDynamic"
	KEY_STANDARD = "PictureModeStandard"
	KEY_MOVIE1   = "PictureModeMovie"
	KEY_GAME     = "PictureModeGame"
	KEY_CUSTOM   = "PictureModeCustom"
	KEY_EXT9     = "PictureModeMovieAlt"
	KEY_EXT10    = "PictureModeStandardAlt"

	//
	// Menus Keys
	//

	Menu     = "KEY_MENU"
	TopMenu  = "KEY_TOPMENU"
	Tools    = "KEY_TOOLS"
	Home     = "KEY_HOME"
	Contents = "KEY_CONTENTS"
	Guide    = "KEY_GUIDE"
	DiscMenu = "KEY_DISC_MENU"
	DVRMenu  = "KEY_DVR_MENU"
	Help     = "KEY_HELP"

	//
	// OSD Keys
	//

	Info         = "KEY_INFO"
	Caption      = "KEY_CAPTION"
	ClockDisplay = "KEY_CLOCK_DISPLAY"
	SetupClock   = "KEY_SETUP_CLOCK_TIMER"
	Subtitle     = "KEY_SUB_TITLE"

	//
	// Zoom Keys
	//

	ZoomMove = "KEY_ZOOM_MOVE"
	ZoomIn   = "KEY_ZOOM_IN"
	ZoomOut  = "KEY_ZOOM_OUT"
	Zoom1    = "KEY_ZOOM1"
	Zoom2    = "KEY_ZOOM2"

	//
	// Other Keys
	//

	WheelLeft             = "KEY_WHEEL_LEFT"
	WheelRight            = "KEY_WHEEL_RIGHT"
	AddOrDel              = "KEY_ADDDEL"
	Plus100               = "KEY_PLUS100"
	AD                    = "KEY_AD"
	Link                  = "KEY_LINK"
	Turbo                 = "KEY_TURBO"
	Convergence           = "KEY_CONVERGENCE"
	DeviceConnect         = "KEY_DEVICE_CONNECT"
	Key11                 = "KEY_11"
	Key12                 = "KEY_12"
	KeyFactory            = "KEY_FACTORY"
	Key3SPEED             = "KEY_3SPEED"
	KeyRSURF              = "KEY_RSURF"
	FF_                   = "KEY_FF_"
	REWIND_               = "KEY_REWIND_"
	Angle                 = "KEY_ANGLE"
	Reserved1             = "KEY_RESERVED1"
	Program               = "KEY_PROGRAM"
	Bookmark              = "KEY_BOOKMARK"
	Print                 = "KEY_PRINT"
	Clear                 = "KEY_CLEAR"
	VChip                 = "KEY_VCHIP"
	Repeat                = "KEY_REPEAT"
	Door                  = "KEY_DOOR"
	Open                  = "KEY_OPEN"
	DMA                   = "KEY_DMA"
	MTS                   = "KEY_MTS"
	DNIe                  = "KEY_DNIe"
	SRS                   = "KEY_SRS"
	ConvertAudioMainOrSub = "KEY_CONVERT_AUDIO_MAINSUB"
	MDC                   = "KEY_MDC"
	SoundEffect           = "KEY_SEFFECT"
	PERPECTFocus          = "KEY_PERPECT_FOCUS"
	CallerID              = "KEY_CALLER_ID"
	Scale                 = "KEY_SCALE"
	MagicBright           = "KEY_MAGIC_BRIGHT"
	WLink                 = "KEY_W_LINK"
	DTVLink               = "KEY_DTV_LINK"
	ApplicationList       = "KEY_APP_LIST"
	BackMHP               = "KEY_BACK_MHP"
	AlternateMHP          = "KEY_ALT_MHP"
	DNSe                  = "KEY_DNSe"
	RSS                   = "KEY_RSS"
	Entertainment         = "KEY_ENTERTAINMENT"
	IDInput               = "KEY_ID_INPUT"
	IDSetup               = "KEY_ID_SETUP"
	AnyView               = "KEY_ANYVIEW"
	MS                    = "KEY_MS"
	More                  = "KEY_MORE"
	Mic                   = "KEY_MIC"
	NineSeparate          = "KEY_NINE_SEPERATE"
	AutoFormat            = "KEY_AUTO_FORMAT"
	DNET                  = "KEY_DNET"

	//
	// Auto Arc Keys
	//

	AUTO_ARC_C_FORCE_AGING     = "KEY_AUTO_ARC_C_FORCE_AGING"
	AUTO_ARC_CAPTION_ENG       = "KEY_AUTO_ARC_CAPTION_ENG"
	AUTO_ARC_USBJACK_INSPECT   = "KEY_AUTO_ARC_USBJACK_INSPECT"
	AUTO_ARC_RESET             = "KEY_AUTO_ARC_RESET"
	AUTO_ARC_LNA_ON            = "KEY_AUTO_ARC_LNA_ON"
	AUTO_ARC_LNA_OFF           = "KEY_AUTO_ARC_LNA_OFF"
	AUTO_ARC_ANYNET_MODE_OK    = "KEY_AUTO_ARC_ANYNET_MODE_OK"
	AUTO_ARC_ANYNET_AUTO_START = "KEY_AUTO_ARC_ANYNET_AUTO_START"
	AUTO_ARC_CAPTION_ON        = "KEY_AUTO_ARC_CAPTION_ON"
	AUTO_ARC_CAPTION_OFF       = "KEY_AUTO_ARC_CAPTION_OFF"
	AUTO_ARC_PIP_DOUBLE        = "KEY_AUTO_ARC_PIP_DOUBLE"
	AUTO_ARC_PIP_LARGE         = "KEY_AUTO_ARC_PIP_LARGE"
	AUTO_ARC_PIP_LEFT_TOP      = "KEY_AUTO_ARC_PIP_LEFT_TOP"
	AUTO_ARC_PIP_RIGHT_TOP     = "KEY_AUTO_ARC_PIP_RIGHT_TOP"
	AUTO_ARC_PIP_LEFT_BOTTOM   = "KEY_AUTO_ARC_PIP_LEFT_BOTTOM"
	AUTO_ARC_PIP_CH_CHANGE     = "KEY_AUTO_ARC_PIP_CH_CHANGE"
	AUTO_ARC_AUTOCOLOR_SUCCESS = "KEY_AUTO_ARC_AUTOCOLOR_SUCCESS"
	AUTO_ARC_AUTOCOLOR_FAIL    = "KEY_AUTO_ARC_AUTOCOLOR_FAIL"
	AUTO_ARC_JACK_IDENT        = "KEY_AUTO_ARC_JACK_IDENT"
	AUTO_ARC_CAPTION_KOR       = "KEY_AUTO_ARC_CAPTION_KOR"
	AUTO_ARC_ANTENNA_AIR       = "KEY_AUTO_ARC_ANTENNA_AIR"
	AUTO_ARC_ANTENNA_CABLE     = "KEY_AUTO_ARC_ANTENNA_CABLE"
	AUTO_ARC_ANTENNA_SATELLITE = "KEY_AUTO_ARC_ANTENNA_SATELLITE"

	//
	// Panel Keys
	//

	PANNEL_POWER  = "KEY_PANNEL_POWER"
	PANNEL_CHUP   = "KEY_PANNEL_CHUP"
	PANNEL_VOLUP  = "KEY_PANNEL_VOLUP"
	PANNEL_VOLDOW = "KEY_PANNEL_VOLDOW"
	PANNEL_MENU   = "KEY_PANNEL_MENU"
	PANNEL_SOURCE = "KEY_PANNEL_SOURCE"
	PANNEL_ENTER  = "KEY_PANNEL_ENTER"
)
