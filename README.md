<p align="center">
    <H1>Samsung Smart TV API</H1>
</p>

This project is a library for remote controlling Samsung televisions via a TCP/IP connection.

It currently supports modern (post-2016) TVs with Ethernet or Wi-Fi connectivity. They should be all models with
TizenOs.

## Install

## Usage

### Basic Setup & Usage

```go
config := helpers.LoadConfiguration()

c := samsung_tv_api.NewSamsungTvWebSocket(
	"IP_ADDRESS_OF_TV",
	// configuration token used to authorise with the
	// TV. If not set the device will need to be allowed
	// via the TV and the token stored for later. 
	config.Token,
	// TLS Port
	8002,
	// Number of milliseconds between each key press
	// when sending large amounts of keys at once.
	2,
	"REGISTERED_DEVICE_NAME",
	// automatically attempt to connect to the TV via
	// the rest and websocket api's to set the token.
	true)

deviceInfo, deviceInfoErr := c.Rest.GetDeviceInfo()

if deviceInfoErr != nil {
	log.Fatal(deviceInfoErr)
}

updatedToken := c.GetToken()

// Use your own provided implementation to store the auth
// token for later use, this stops the TV asking the user
// to confirm the device.
if updatedToken != "" && updatedToken != config.Token {
	config.Mac = device.Device.WifiMac
	config.Token = c.GetToken()
	_ = helpers.SaveConfiguration(&config)
}
```

### Toggle Power

```go 
if err := c.Websocket.Power(); err != nil {
	log.Fatalln(err)
}
```

### Open Browser

```go
if err := c.Websocket.OpenBrowser("https://www.google.com"); err != nil {
	log.Fatalln(err)
}
```

### Get Applications

```golang
if apps, err := c.Websocket.GetApplicationsList(); err == nil {
	firstApp := apps.Data.Applications[0]
	fmt.Println(fmt.Sprintf("application name: %s", firstApp.Name))
}
```

### Get Application Details

```golang
if apps, err := c.Websocket.GetApplicationsList(); err == nil {
	firstApp := apps.Data.Applications[0]

	appDetails, _ := c.Rest.GetApplicationStatus(firstApp.AppID)
	byteData, _ := json.MarshalIndent(appDetails, " ", "\t")

	fmt.Println(string(byteData))
}
```

### Close Application

```golang
if _, err := c.Rest.CloseApplication(AppID); err != nil {
	fmt.Println(err)
}
```

### Install Application
```golang
if _, err := c.Rest.InstallApplication(AppID); err != nil {
	fmt.Println(err)
}
```

### Get TV Information
```golang
if deviceDetails, err := c.Rest.GetDeviceInfo(); err == nil {
	byteData, _ := json.MarshalIndent(deviceDetails, "", "\t")
	fmt.Println(string(byteData))
}
```

### Wake on Lan
```go
if err := samsung_tv_api.WakeOnLan(config.Mac); err == nil {
	c.Websocket.Power()
}
```
## Full API Listings

### Rest
	* GetDeviceInfo() (DeviceResponse, error)
	* GetApplicationStatus(appId string) (ApplicationResponse, error)
	* RunApplication(appId string) (interface{}, error)
	* CloseApplication(appId string) (interface{}, error)
	* InstallApplication(appId string) (interface{}, error)

### Websocket
	* GetApplicationList() (ApplicationResopnse, error)
	* RunApplication(appId, appType, metaTag string) error
	* SendClick(key string) error
	* SendKey(key string, times int, cmd string) error
	* HoldKey(key string, seconds int) error
	* ChangeChannel(channel string) error
	* MoveCursor(x, y, duration int) error
	* OpenBrowser(url string) error
	* Power() error
	* PowerOff() error
	* PowerOn() error
	* Disconnect() error

### UpNp
	* GetCurrentVolume() (int, error)
	* SetVolume(volume int) error 
	* GetCurrentMuteStatus() (bool, error) 
	* SetCurrentMedia(url string) error 
	* PlayCurrentMedia() error 

### Other
	* WakeOnLan(mac string) error

## Supported TVs

List of support TV
models. https://developer.samsung.com/smarttv/develop/extension-libraries/smart-view-sdk/supported-device/supported-tvs.html

```
2017 : M5500 and above
2016 : K4300, K5300 and above
2015 : J5500 and above (except J6203)
2014 : H4500, H5500 and above (except H6003/H6103/H6153/H6201/H6203)
Supported TV models may vary by region.
```

For complete list https://developer.samsung.com/smarttv/develop/specifications/tv-model-groups.html

## License

[MIT](./LICENSE.md)
