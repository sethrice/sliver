package configs

/*
	Sliver Implant Framework
	Copyright (C) 2021  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"errors"
	"fmt"
	insecureRand "math/rand"
	"path"
	"regexp"
	"strings"

	"github.com/bishopfox/sliver/protobuf/clientpb"
	"github.com/bishopfox/sliver/server/assets"
	"github.com/bishopfox/sliver/server/log"
)

const (
	httpC2ConfigFileName = "http-c2.json"
	DefaultChromeBaseVer = 106
	DefaultMacOSVer      = "10_15_7"
)

// HTTPC2Config - Parent config file struct for implant/server
type HTTPC2Config struct {
	ImplantConfig *HTTPC2ImplantConfig `json:"implant_config"`
	ServerConfig  *HTTPC2ServerConfig  `json:"server_config"`
}

// RandomImplantConfig - Randomly generate a new implant config from the parent config,
// this is the primary configuration used by the implant generation.
func (h *HTTPC2Config) RandomImplantConfig() *HTTPC2ImplantConfig {
	return &HTTPC2ImplantConfig{

		NonceQueryArgs: h.ImplantConfig.NonceQueryArgs,
		URLParameters:  h.ImplantConfig.URLParameters,
		Headers:        h.ImplantConfig.Headers,

		PollFileExt: h.ImplantConfig.PollFileExt,
		PollFiles:   h.ImplantConfig.RandomPollFiles(),
		PollPaths:   h.ImplantConfig.RandomPollPaths(),

		StartSessionFileExt: h.ImplantConfig.StartSessionFileExt,
		SessionFileExt:      h.ImplantConfig.SessionFileExt,
		SessionFiles:        h.ImplantConfig.RandomSessionFiles(),
		SessionPaths:        h.ImplantConfig.RandomSessionPaths(),

		CloseFileExt: h.ImplantConfig.CloseFileExt,
		CloseFiles:   h.ImplantConfig.RandomCloseFiles(),
		ClosePaths:   h.ImplantConfig.RandomClosePaths(),
	}
}

// GenerateUserAgent - Generate a user-agent depending on OS/Arch
func (h *HTTPC2Config) GenerateUserAgent(goos string, goarch string) string {
	return h.generateChromeUserAgent(goos, goarch)
}

func (h *HTTPC2Config) generateChromeUserAgent(goos string, goarch string) string {
	if h.ImplantConfig.UserAgent == "" {
		switch goos {
		case "windows":
			switch goarch {
			case "amd64":
				return fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", h.ChromeVer())
			}

		case "linux":
			switch goarch {
			case "amd64":
				return fmt.Sprintf("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", h.ChromeVer())
			}

		case "darwin":
			switch goarch {
			case "arm64":
				fallthrough // https://source.chromium.org/chromium/chromium/src/+/master:third_party/blink/renderer/core/frame/navigator_id.cc;l=76
			case "amd64":
				return fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X %s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", h.MacOSVer(), h.ChromeVer())
			}

		}
	} else {
		return h.ImplantConfig.UserAgent
	}

	// Default is a generic Windows/Chrome
	return fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", h.ChromeVer())
}

// ChromeVer - Generate a random Chrome user-agent
func (h *HTTPC2Config) ChromeVer() string {
	chromeVer := h.ImplantConfig.ChromeBaseVersion
	if chromeVer == 0 {
		chromeVer = DefaultChromeBaseVer
	}
	return fmt.Sprintf("%d.0.%d.%d", chromeVer+insecureRand.Intn(3), 1000+insecureRand.Intn(8999), insecureRand.Intn(999))
}

func (h *HTTPC2Config) MacOSVer() string {
	macosVer := h.ImplantConfig.MacOSVersion
	if macosVer == "" {
		macosVer = DefaultMacOSVer
	}
	return macosVer
}

// HTTPC2ServerConfig - Server configuration options
type HTTPC2ServerConfig struct {
	RandomVersionHeaders bool                   `json:"random_version_headers"`
	Headers              []NameValueProbability `json:"headers"`
	Cookies              []string               `json:"cookies"`
}

type NameValueProbability struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Probability int    `json:"probability"`
	Methods     []string
}

// HTTPC2ImplantConfig - Implant configuration options
// Procedural C2
// ===============
// .txt = rsakey
// .css = start
// .php = session
//
//	.js = poll
//
// .png = stop
// .woff = sliver shellcode
type HTTPC2ImplantConfig struct {
	UserAgent         string `json:"user_agent"`
	ChromeBaseVersion int    `json:"chrome_base_version"`
	MacOSVersion      string `json:"macos_version"`

	NonceQueryArgs string                 `json:"nonce_query_args"`
	URLParameters  []NameValueProbability `json:"url_parameters"`
	Headers        []NameValueProbability `json:"headers"`

	MaxFiles int `json:"max_files"`
	MinFiles int `json:"min_files"`
	MaxPaths int `json:"max_paths"`
	MinPaths int `json:"min_paths"`

	// Stager File Extension
	StagerFileExt string `json:"stager_file_ext"`

	// Poll files and paths
	PollFileExt string   `json:"poll_file_ext"`
	PollFiles   []string `json:"poll_files"`
	PollPaths   []string `json:"poll_paths"`

	// Session files and paths
	StartSessionFileExt string   `json:"start_session_file_ext"`
	SessionFileExt      string   `json:"session_file_ext"`
	SessionFiles        []string `json:"session_files"`
	SessionPaths        []string `json:"session_paths"`

	// Close session files and paths
	CloseFileExt string   `json:"close_file_ext"`
	CloseFiles   []string `json:"close_files"`
	ClosePaths   []string `json:"close_paths"`
}

func (h *HTTPC2ImplantConfig) RandomPollFiles() []string {
	min := h.MinFiles
	if min < 1 {
		min = 1
	}
	return h.randomSample(h.PollFiles, h.PollFileExt, min, h.MaxFiles)
}

func (h *HTTPC2ImplantConfig) RandomPollPaths() []string {
	return h.randomSample(h.PollPaths, "", h.MinPaths, h.MaxPaths)
}

func (h *HTTPC2ImplantConfig) RandomCloseFiles() []string {
	min := h.MinFiles
	if min < 1 {
		min = 1
	}
	return h.randomSample(h.CloseFiles, h.CloseFileExt, min, h.MaxFiles)
}

func (h *HTTPC2ImplantConfig) RandomClosePaths() []string {
	return h.randomSample(h.ClosePaths, "", h.MinPaths, h.MaxPaths)
}

func (h *HTTPC2ImplantConfig) RandomSessionFiles() []string {
	min := h.MinFiles
	if min < 1 {
		min = 1
	}
	return h.randomSample(h.SessionFiles, h.SessionFileExt, min, h.MaxFiles)
}

func (h *HTTPC2ImplantConfig) RandomSessionPaths() []string {
	return h.randomSample(h.SessionPaths, "", h.MinPaths, h.MaxPaths)
}

func (h *HTTPC2ImplantConfig) randomSample(values []string, ext string, min int, max int) []string {
	count := insecureRand.Intn(len(values))
	if count < min {
		count = min
	}
	if max < count {
		count = max
	}
	sample := []string{}
	for i := 0; len(sample) < count; i++ {
		index := (count + i) % len(values)
		sample = append(sample, values[index])
	}
	return sample
}

// GetHTTPC2ConfigPath - File path to http-c2.json
func GetHTTPC2ConfigPath() string {
	appDir := assets.GetRootAppDir()
	httpC2ConfigPath := path.Join(appDir, "configs", httpC2ConfigFileName)
	return httpC2ConfigPath
}

// CheckHTTPC2ConfigErrors - Get the current HTTP C2 config
func CheckHTTPC2ConfigErrors(config *clientpb.HTTPC2Config) error {
	err := checkHTTPC2Config(config)
	if err != nil {
		httpC2ConfigLog.Errorf("Invalid http c2 config: %s", err)
		return err
	}
	return nil
}

var (
	ErrMissingCookies             = errors.New("server config must specify at least one cookie")
	ErrMissingStagerFileExt       = errors.New("implant config must specify a stager_file_ext")
	ErrMissingPollFileExt         = errors.New("implant config must specify a poll_file_ext")
	ErrTooFewPollFiles            = errors.New("implant config must specify at least one poll_files value")
	ErrMissingKeyExchangeFileExt  = errors.New("implant config must specify a key_exchange_file_ext")
	ErrTooFewKeyExchangeFiles     = errors.New("implant config must specify at least one key_exchange_files value")
	ErrMissingCloseFileExt        = errors.New("implant config must specify a close_file_ext")
	ErrTooFewCloseFiles           = errors.New("implant config must specify at least one close_files value")
	ErrMissingStartSessionFileExt = errors.New("implant config must specify a start_session_file_ext")
	ErrMissingSessionFileExt      = errors.New("implant config must specify a session_file_ext")
	ErrTooFewSessionFiles         = errors.New("implant config must specify at least one session_files value")
	ErrNonUniqueFileExt           = errors.New("implant config must specify unique file extensions")
	ErrQueryParamNameLen          = errors.New("implant config url query parameter names must be 3 or more characters")

	fileNameExp = regexp.MustCompile(`[^a-zA-Z0-9\\._-]+`)
)

// checkHTTPC2Config - Validate the HTTP C2 config, coerces common mistakes
func checkHTTPC2Config(config *clientpb.HTTPC2Config) error {
	err := checkServerConfig(config.ServerConfig)
	if err != nil {
		return err
	}
	return checkImplantConfig(config.ImplantConfig)
}

func coerceFileExt(value string) string {
	value = fileNameExp.ReplaceAllString(value, "")
	for strings.HasPrefix(value, ".") {
		value = strings.TrimPrefix(value, ".")
	}
	return value
}

func coerceFiles(values []string, ext string) []string {
	values = uniqueFileName(values)
	coerced := []string{}
	for _, value := range values {
		if strings.HasSuffix(value, fmt.Sprintf(".%s", ext)) {
			value = strings.TrimSuffix(value, fmt.Sprintf(".%s", ext))
		}
		coerced = append(coerced, value)
	}
	return coerced
}

func uniqueFileName(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		item = fileNameExp.ReplaceAllString(item, "")
		if len(item) < 1 {
			continue
		}
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func checkServerConfig(config *clientpb.HTTPC2ServerConfig) error {
	if len(config.Cookies) < 1 {
		return ErrMissingCookies
	}
	return nil
}

func checkImplantConfig(config *clientpb.HTTPC2ImplantConfig) error {

	// MinFiles and MaxFiles
	if config.MinFiles < 1 {
		config.MinFiles = 1
	}
	if config.MaxFiles < config.MinFiles {
		config.MaxFiles = config.MinFiles
	}

	// MinPaths and MaxPaths
	if config.MinPaths < 0 {
		config.MinPaths = 0
	}
	if config.MaxPaths < config.MinPaths {
		config.MaxPaths = config.MinPaths
	}

	// Stager
	config.StagerFileExtension = coerceFileExt(config.StagerFileExtension)
	if config.StagerFileExtension == "" {
		return ErrMissingStagerFileExt
	}

	// File Extensions
	config.PollFileExtension = coerceFileExt(config.PollFileExtension)
	if config.PollFileExtension == "" {
		return ErrMissingPollFileExt
	}
	config.StartSessionFileExtension = coerceFileExt(config.StartSessionFileExtension)
	if config.StartSessionFileExtension == "" {
		return ErrMissingStartSessionFileExt
	}
	config.SessionFileExtension = coerceFileExt(config.SessionFileExtension)
	if config.SessionFileExtension == "" {
		return ErrMissingSessionFileExt
	}
	config.CloseFileExtension = coerceFileExt(config.CloseFileExtension)
	if config.CloseFileExtension == "" {
		return ErrMissingCloseFileExt
	}

	return nil
}

func GenerateDefaultHTTPC2Config() *clientpb.HTTPC2Config {

	// Implant Config
	httpC2UrlParameters := []*clientpb.HTTPC2URLParameter{}
	httpC2Headers := []*clientpb.HTTPC2Header{}
	pathSegments := GenerateHTTPC2DefaultPathSegment()

	implantConfig := clientpb.HTTPC2ImplantConfig{
		UserAgent:                 "",
		ChromeBaseVersion:         DefaultChromeBaseVer,
		MacOSVersion:              DefaultMacOSVer,
		NonceQueryArgChars:        "abcdefghijklmnopqrstuvwxyz",
		ExtraURLParameters:        httpC2UrlParameters,
		Headers:                   httpC2Headers,
		MaxFiles:                  8,
		MinFiles:                  2,
		MaxPaths:                  8,
		MinPaths:                  2,
		StagerFileExtension:       ".woff",
		PollFileExtension:         ".js",
		StartSessionFileExtension: ".html",
		SessionFileExtension:      ".php",
		CloseFileExtension:        ".png",
		PathSegments:              pathSegments,
	}

	// Server Config
	serverHeaders := []*clientpb.HTTPC2Header{
		{
			Method:      "GET",
			Name:        "Cache-Control",
			Value:       "no-store, no-cache, must-revalidate",
			Probability: 100,
		},
	}
	cookies := GenerateDefaultHTTPC2Cookies()
	serverConfig := clientpb.HTTPC2ServerConfig{
		RandomVersionHeaders: false,
		Headers:              serverHeaders,
		Cookies:              cookies,
	}

	// HTTPC2Config

	defaultConfig := clientpb.HTTPC2Config{
		ServerConfig:  &serverConfig,
		ImplantConfig: &implantConfig,
		Name:          "default",
	}

	return &defaultConfig
}

func GenerateDefaultHTTPC2Cookies() []*clientpb.HTTPC2Cookie {
	cookies := []*clientpb.HTTPC2Cookie{}
	for _, cookie := range Cookies {
		cookies = append(cookies, &clientpb.HTTPC2Cookie{
			Name: cookie,
		})
	}
	return cookies
}

func GenerateHTTPC2DefaultPathSegment() []*clientpb.HTTPC2PathSegment {
	pathSegments := []*clientpb.HTTPC2PathSegment{}

	/*
		IsFile      bool
		SegmentType int32 // Poll, Session, Close
		Value       string
	*/

	// files
	for _, poll := range PollFiles {
		pathSegments = append(pathSegments, &clientpb.HTTPC2PathSegment{
			IsFile:      true,
			SegmentType: 0,
			Value:       poll,
		})
	}

	for _, session := range SessionFiles {
		pathSegments = append(pathSegments, &clientpb.HTTPC2PathSegment{
			IsFile:      true,
			SegmentType: 1,
			Value:       session,
		})
	}

	for _, close := range CloseFiles {
		pathSegments = append(pathSegments, &clientpb.HTTPC2PathSegment{
			IsFile:      true,
			SegmentType: 2,
			Value:       close,
		})
	}

	// paths
	for _, poll := range PollPaths {
		pathSegments = append(pathSegments, &clientpb.HTTPC2PathSegment{
			IsFile:      true,
			SegmentType: 0,
			Value:       poll,
		})
	}

	for _, session := range SessionPaths {
		pathSegments = append(pathSegments, &clientpb.HTTPC2PathSegment{
			IsFile:      true,
			SegmentType: 1,
			Value:       session,
		})
	}

	for _, close := range ClosePaths {
		pathSegments = append(pathSegments, &clientpb.HTTPC2PathSegment{
			IsFile:      true,
			SegmentType: 2,
			Value:       close,
		})
	}

	return pathSegments
}

var (
	httpC2ConfigLog = log.NamedLogger("config", "http-c2")

	Cookies = []string{
		"PHPSESSID",
		"SID",
		"SSID",
		"APISID",
		"csrf-state",
		"AWSALBCORS",
	}
	PollFiles = []string{
		"bootstrap",
		"bootstrap.min",
		"jquery.min",
		"jquery",
		"route",
		"app",
		"app.min",
		"array",
		"backbone",
		"script",
		"email",
	}
	PollPaths = []string{
		"js",
		"umd",
		"assets",
		"bundle",
		"bundles",
		"scripts",
		"script",
		"javascripts",
		"javascript",
		"jscript",
	}
	SessionFiles = []string{
		"login",
		"signin",
		"api",
		"samples",
		"rpc",
		"index",
		"admin",
		"register",
		"sign-up",
	}
	SessionPaths = []string{
		"php",
		"api",
		"upload",
		"actions",
		"rest",
		"v1",
		"auth",
		"authenticate",
		"oauth",
		"oauth2",
		"oauth2callback",
		"database",
		"db",
		"namespaces",
	}
	CloseFiles = []string{
		"favicon",
		"sample",
		"example",
	}
	ClosePaths = []string{
		"static",
		"www",
		"assets",
		"images",
		"icons",
		"image",
		"icon",
		"png",
	}
)
