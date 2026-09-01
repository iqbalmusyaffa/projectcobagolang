package utils

import (
	"fmt"
	"strings"
)

// ParseUserAgent menganalisis string User-Agent dari browser dan mengembalikan nama perangkat, browser, sistem operasi, serta jenis perangkat (desktop, mobile, tablet).
func ParseUserAgent(ua string) (deviceName, browser, os, deviceType string) {
	if ua == "" {
		return "Perangkat Tidak Dikenal", "Browser Lain", "OS Lain", "desktop"
	}

	uaLower := strings.ToLower(ua)

	// 1. Deteksi Sistem Operasi (OS)
	if strings.Contains(uaLower, "windows nt 10.0") || strings.Contains(uaLower, "windows nt 11.0") {
		os = "Windows 10/11"
	} else if strings.Contains(uaLower, "windows nt 6.3") {
		os = "Windows 8.1"
	} else if strings.Contains(uaLower, "windows nt 6.1") {
		os = "Windows 7"
	} else if strings.Contains(uaLower, "windows") {
		os = "Windows"
	} else if strings.Contains(uaLower, "macintosh") || strings.Contains(uaLower, "mac os x") {
		os = "macOS"
	} else if strings.Contains(uaLower, "iphone") {
		os = "iOS (iPhone)"
	} else if strings.Contains(uaLower, "ipad") {
		os = "iPadOS"
	} else if strings.Contains(uaLower, "android") {
		os = "Android"
	} else if strings.Contains(uaLower, "linux") {
		os = "Linux"
	} else {
		os = "Sistem Lain"
	}

	// 2. Deteksi Jenis Perangkat (Device Type)
	if strings.Contains(uaLower, "ipad") || strings.Contains(uaLower, "tablet") {
		deviceType = "tablet"
	} else if strings.Contains(uaLower, "mobile") || strings.Contains(uaLower, "android") || strings.Contains(uaLower, "iphone") {
		deviceType = "mobile"
	} else {
		deviceType = "desktop"
	}

	// 3. Deteksi Browser
	if strings.Contains(uaLower, "postmanruntime") {
		browser = "Postman"
	} else if strings.Contains(uaLower, "edg/") || strings.Contains(uaLower, "edge/") {
		browser = "Microsoft Edge"
	} else if strings.Contains(uaLower, "opr/") || strings.Contains(uaLower, "opera") {
		browser = "Opera"
	} else if strings.Contains(uaLower, "chrome") && !strings.Contains(uaLower, "edg") && !strings.Contains(uaLower, "opr") {
		browser = "Google Chrome"
	} else if strings.Contains(uaLower, "firefox") {
		browser = "Mozilla Firefox"
	} else if strings.Contains(uaLower, "safari") && !strings.Contains(uaLower, "chrome") {
		browser = "Apple Safari"
	} else {
		browser = "Web Browser"
	}

	// 4. Susun Nama Perangkat yang Elegan
	deviceName = fmt.Sprintf("%s di %s", browser, os)

	return deviceName, browser, os, deviceType
}
