<p align="center">
    <H1>Samsung Smart TV API</H1>
</p>

This project is a library for remote controlling Samsung televisions via a TCP/IP connection.

It currently supports modern (post-2016) TVs with Ethernet or Wi-Fi connectivity. They should be all models with TizenOs.

## Install

## Usage

### Basic Setup
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
if err := c.Websocket.PowerOn(); err != nil {
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

### Get Application

### Close Application

### Install Application

### Get TV Information

## Full API Listings

## Supported TVs

List of support TV models. https://developer.samsung.com/smarttv/develop/extension-libraries/smart-view-sdk/supported-device/supported-tvs.html

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
