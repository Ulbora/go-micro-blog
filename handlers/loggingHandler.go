// Package handlers ...
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	lg "github.com/GolangToolKits/go-level-logger"
)

// LogLevel LogLevel
type LogLevel struct {
	Level string `json:"logLevel"`
}

// LogResponse LogResponse
type LogResponse struct {
	Success  bool   `json:"success"`
	LogLevel string `json:"logLevel"`
}

const (
	defaultLoggingKey = "45sdbb2345"

	debugLevel = "DEBUG"
	infoLevel  = "INFO"
	allLevel   = "ALL"
	offLevel   = "OFF"
)

// SetLogLevel SetLogLevel
func (h *MCHandler) SetLogLevel(w http.ResponseWriter, r *http.Request) {
	var logRes LogResponse
	h.setContentType(w)
	logContOk := h.checkContent(r)

	//fmt.Println("conOk: ", logContOk)

	if !logContOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var loggingKey string
		if os.Getenv("LOGGING_KEY") != "" {
			loggingKey = os.Getenv("LOGGING_KEY")
		} else {
			loggingKey = defaultLoggingKey
		}
		loggingKeyHdr := r.Header.Get("Logging_KEY")
		if loggingKey == loggingKeyHdr {
			var lv LogLevel
			lgsuc, lgerr := h.processBody(r, &lv)
			//fmt.Println("lgsuc: ", lgsuc)
			//fmt.Println("LogLevel: ", lv)
			//fmt.Println("lgerr: ", lgerr)
			if !lgsuc && lgerr != nil {
				http.Error(w, lgerr.Error(), http.StatusBadRequest)
			} else {
				switch strings.ToUpper(lv.Level) {
				case debugLevel:
					h.Log.SetLogLevel(lg.DebugLevel)
					logRes.Success = true
					logRes.LogLevel = debugLevel
				case infoLevel:
					h.Log.SetLogLevel(lg.InfoLevel)
					logRes.Success = true
					logRes.LogLevel = infoLevel
				case allLevel:
					h.Log.SetLogLevel(lg.AllLevel)
					logRes.Success = true
					logRes.LogLevel = allLevel
				case offLevel:
					h.Log.SetLogLevel(lg.OffLevel)
					logRes.Success = true
					logRes.LogLevel = offLevel
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
		resJSON, _ := json.Marshal(logRes)
		fmt.Fprint(w, string(resJSON))
	}
}
