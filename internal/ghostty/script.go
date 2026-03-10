package ghostty

import (
	"bytes"
	"embed"
	"fmt"
	"strings"
	"text/template"
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

const BundleID = "com.mitchellh.ghostty"

//go:embed scripts/*.applescript
var scriptFS embed.FS

var templates = template.Must(
	template.New("").Funcs(template.FuncMap{
		"escape": escapeAppleScript,
	}).ParseFS(scriptFS, "scripts/*.applescript"),
)

func escapeAppleScript(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	s = strings.ReplaceAll(s, "\n", `\n`)
	s = strings.ReplaceAll(s, "\r", `\r`)
	s = strings.ReplaceAll(s, "\t", `\t`)
	return s
}

func renderScript(name string, data any) (string, error) {
	var buf bytes.Buffer
	if err := templates.ExecuteTemplate(&buf, name, data); err != nil {
		return "", fmt.Errorf("render %s: %w", name, err)
	}

	return buf.String(), nil
}

func execScript(source string) (result foundation.AppleEventDescriptor, err error) {
	objc.WithAutoreleasePool(func() {
		var errDict objc.Object
		script := foundation.NewAppleScriptWithSource(source)
		result = script.ExecuteAndReturnError(unsafe.Pointer(&errDict))
		if result.IsNil() {
			if !errDict.IsNil() {
				msg := objc.Call[string](errDict, objc.Sel("objectForKey:"), foundation.NewStringWithString("NSAppleScriptErrorMessage"))
				if msg != "" {
					err = fmt.Errorf("applescript: %s", msg)
					return
				}
			}

			err = fmt.Errorf("applescript execution failed")
			return
		}

		result.Retain()
	})
	return
}

func execTemplate(name string, data any) (foundation.AppleEventDescriptor, error) {
	running, err := isRunning()
	if err != nil {
		return foundation.AppleEventDescriptor{}, err
	}

	if !running {
		return foundation.AppleEventDescriptor{}, fmt.Errorf("ghostty is not running")
	}

	src, err := renderScript(name, data)
	if err != nil {
		return foundation.AppleEventDescriptor{}, err
	}

	return execScript(src)
}

func isRunning() (running bool, err error) {
	src, err := renderScript("app_is_running.applescript", baseParams())
	if err != nil {
		return false, err
	}

	objc.WithAutoreleasePool(func() {
		script := foundation.NewAppleScriptWithSource(src)
		result := script.ExecuteAndReturnError(nil)
		if result.IsNil() {
			err = fmt.Errorf("failed to check if Ghostty is running")
			return
		}

		running = result.BooleanValue()
	})
	return
}

type baseData struct {
	BundleID string
}

func baseParams() baseData {
	return baseData{BundleID: BundleID}
}

type surfaceConfigData struct {
	baseData
	HasConfig        bool
	Command          string
	WorkingDir       string
	InitialInput     string
	FontSize         float64
	WaitAfterCommand bool
	EnvVarsList      string
}

func surfaceConfigParams(cfg SurfaceConfig) surfaceConfigData {
	d := surfaceConfigData{baseData: baseParams()}
	d.Command = cfg.Command
	d.WorkingDir = cfg.WorkingDir
	d.InitialInput = cfg.InitialInput
	d.FontSize = cfg.FontSize
	d.WaitAfterCommand = cfg.WaitAfterCommand
	if len(cfg.EnvVars) > 0 {
		quoted := make([]string, len(cfg.EnvVars))
		for i, v := range cfg.EnvVars {
			quoted[i] = fmt.Sprintf(`"%s"`, escapeAppleScript(v))
		}

		d.EnvVarsList = strings.Join(quoted, ", ")
	}

	d.HasConfig = cfg.Command != "" || cfg.WorkingDir != "" || cfg.InitialInput != "" || cfg.FontSize > 0 || cfg.WaitAfterCommand || len(cfg.EnvVars) > 0
	return d
}
