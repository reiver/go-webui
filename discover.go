package webui

import (
	"bytes"
	"path/filepath"
	"os"
)

// discoverChromium tries to find the Chromium web browser, or the Google Chrome Web Browser on the user's computer.
//
// It first tried to do this using the ‘which’ command that is commonly found on Linux, and Unix-like systems.
//
// If it cannot find it that way, then it tries looking at specific locations on the user's computer.
func discoverChromium() (string, error) {

	var path string
	var err  error

	// These are different potential file names that the Chromium web browser, and the Google Chrome web browser,
	// might be named as on the user's system.
	fileNames := []string{
		"chromium",
		"chromium-browser",
		"google-chrome",
		"google-chrome-stable",
	}

	// Try to use the ‘which’ command to discover where the Chromium web browser, or the Google Chrome web browser,
	// is on the user's system.
	//
	// If find it, then return the path.
	for _, fileName := range fileNames {

		// which $fileName
		path, err = which(fileName)
		if nil == err {
			return path, nil
		}
	}

	// Try to look in specific directories to see if the Chromium web browser, or the Google Chrome web browser, is there.
	//
	// If find it, then return the path.
	{
		dirs := []string{
			"/usr/bin/",
		}

		for _, dir := range dirs {
			for _, fileName := range fileNames {
				var buffer bytes.Buffer

				buffer.WriteString(dir)
				buffer.WriteString(fileName)

				path := buffer.String()

				if _, err := os.Stat(path); !os.IsNotExist(err) {
					return path, nil
				}
			}
		}
	}

	// Try to look at specific full paths to see if the Chromium web browser, or the Google Chrome web browser, is there.
	//
	// If find it, then return the path.
	{
		paths := []string{
			"/Applications/Chromium.app/Contents/MacOS/Chromium",
			"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
			"/Applications/Google Chrome Canary.app/Contents/MacOS/Google Chrome Canary",

			"C:/Program Files/Google/Chrome/Application/chrome.exe",
			"C:/Program Files (x86)/Google/Chrome/Application/chrome.exe",
		}

		for _, path := range paths {
			if _, err := os.Stat(path); !os.IsNotExist(err) {
				return path, nil
			}
		}
	}

	// Try to look at specific full paths for the user to see if the Chromium web browser, or the Google Chrome web browser, is there.
	//
	// If find it, then return the path.
	if userName := username(); "" != userName {
		paths := []string{
			filepath.Join("C:/Users/", userName, "/AppData/Local/Chromium/Application/chromium.exe"),
			filepath.Join("C:/Users/", userName, "/AppData/Local/Chromium/Application/chrome.exe"),
			filepath.Join("C:/Users/", userName, "/AppData/Local/Google/Chrome/Application/chromium.exe"),
			filepath.Join("C:/Users/", userName, "/AppData/Local/Google/Chrome/Application/chrome.exe"),
		}

		for _, path := range paths {
			if _, err := os.Stat(path); !os.IsNotExist(err) {
				return path, nil
			}
		}
	}

	return "", errNotFound
}
