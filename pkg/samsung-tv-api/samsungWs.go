package samsung_tv_api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/stephenSLI/samsung-tv-ws-api/pkg/keys"
	"golang.org/x/net/websocket"
	"log"
	"net/url"
	"time"
)

type samsungWebsocket struct {
	baseUrl       *url.URL
	conn          *websocket.Conn
	keyPressDelay int
}

type WebsocketRequest struct {
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
}

func (s *samsungWebsocket) openConnection() (*WsConnectionResponse, error) {
	origin := "http://localhost/"

	config, _ := websocket.NewConfig(s.baseUrl.String(), origin)
	config.TlsConfig = &tls.Config{InsecureSkipVerify: true}

	ws, err := websocket.DialConfig(config)

	if err != nil {
		return nil, err
	}

	s.conn = ws

	var val WsConnectionResponse
	readErr := s.readJSON(&val)

	return &val, readErr
}

func (s *samsungWebsocket) sendJSONReceiveJSON(command interface{}, output interface{}) error {
	err := s.sendJSON(command)

	if err != nil {
		return err
	}

	return s.readJSON(&output)
}

func (s *samsungWebsocket) sendJSON(command interface{}) error {
	msg, err := json.Marshal(command)

	fmt.Println(string(msg))
	fmt.Println(command)

	if err != nil {
		return err
	}

	_, err = s.conn.Write(msg)
	return err
}

func (s *samsungWebsocket) read() ([]byte, error) {

	var data []byte
	err := websocket.Message.Receive(s.conn, &data)

	return data, err
}

func (s *samsungWebsocket) readJSON(val interface{}) error {
	msg, err := s.read()

	fmt.Println(string(msg))

	if err != nil {
		return err
	}

	return json.Unmarshal(msg, val)
}

func (s *samsungWebsocket) GetApplicationsList() (WsApplicationsResponse, error) {
	log.Println("Get application lists via ws api")

	var output WsApplicationsResponse

	var req = WebsocketRequest{
		Method: "ms.channel.emit",
		Params: map[string]interface{}{
			"event": "ed.installedApp.get",
			"to":    "host",
		},
	}

	err := s.sendJSONReceiveJSON(req, &output)
	return output, err
}

func (s *samsungWebsocket) RunApplication(appId, appType, metaTag string) error {
	log.Printf("Running application %s via ws api\n", appId)

	if appType == "" {
		appType = "DEEP_LINK"
	}

	var req = WebsocketRequest{
		Method: "ms.channel.emit",
		Params: map[string]interface{}{
			"event": "ed.apps.launch",
			"to":    "host",
			"data": map[string]interface{}{
				// action_type: NATIVE_LAUNCH / DEEP_LINK
				// # app_type == 2 ? "DEEP_LINK" : "NATIVE_LAUNCH",
				"action_type": appType,
				"appId":       appId,
				"metaTag":     metaTag,
			},
		},
	}

	return s.sendJSON(req)
}

func (s *samsungWebsocket) sendClick(key string) error {
	return s.SendKey(key, 1, "Click")
}

func (s *samsungWebsocket) SendKey(key string, times int, cmd string) error {

	if cmd == "" {
		cmd = "Click"
	}

	log.Printf("Sending key %s with command %s, %d times via ws api\n", key, cmd, times)

	for i := 0; i < times; i++ {
		log.Printf("Sending key %s via ws api\n", key)

		var req = WebsocketRequest{
			Method: "ms.remote.control",
			Params: map[string]interface{}{
				"Cmd":          cmd,
				"DataOfCmd":    key,
				"Option":       "false",
				"TypeOfRemote": "SendRemoteKey",
			},
		}

		err := s.sendJSON(req)

		if err != nil {
			return err
		}

		time.Sleep(time.Duration(s.keyPressDelay) * time.Millisecond)
	}

	return nil
}

func (s *samsungWebsocket) HoldKey(key string, seconds int) error {
	log.Printf("Sending hold key %s for %d seconds via ws api\n", key, seconds)

	pressErr := s.SendKey(key, 1, "Press")

	if pressErr != nil {
		return pressErr
	}

	time.Sleep(time.Duration(seconds) * time.Second)

	log.Printf("Sending release key %s via ws api\n", key)
	releaseErr := s.SendKey(key, 1, "Release")

	if releaseErr != nil {
		return releaseErr
	}

	return nil
}

func (s *samsungWebsocket) MoveCursor(x, y, duration int) error {
	log.Printf("Sending move Cursor to x: %d, y: %d for duration %d via ws api\n", x, y, duration)

	var req = WebsocketRequest{
		Method: "ms.remote.control",
		Params: map[string]interface{}{
			"Cmd":          "Move",
			"TypeOfRemote": "ProcessMouseDevice",
			"Position": map[string]interface{}{
				"x":    x,
				"y":    y,
				"Time": string(rune(duration)),
			},
		},
	}

	return s.sendJSON(req)
}

func (s *samsungWebsocket) OpenBrowser(url string) error {
	log.Printf("opening browser to url: %s via ws api\n", url)
	return s.RunApplication("org.tizen.browser", "NATIVE_LAUNCH", url)
}

func (s *samsungWebsocket) PowerOff() error {
	return s.sendClick(keys.PowerToggle)
}

func (s *samsungWebsocket) PowerOn() error {
	return s.sendClick(keys.PowerToggle)
}
