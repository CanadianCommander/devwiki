package template

import (
	"fmt"
	"log/slog"
	"regexp"
)

// ===============================================
// Types
// ===============================================

const (
	ThemeDark  = "/assets/css/tailwinds-dark-pack.css"
	ThemeLight = "/assets/css/tailwinds-light-pack.css"
)

// ===============================================
// Public
// ===============================================

// GetCurrentThemeCss gets the path to the current theme css
func GetCurrentThemeCss() string {
	// TODO switch between light and dark variants of the theme based on the current theme
	return ThemeDark
}

// GetImageInTheme gets the path to an image in the current theme. This automatically switches between light and dark
// variants of the image based on the current theme.
func GetImageInTheme(imagePath string) string {
	// TODO switch between light and dark variants of the image based on the current theme

	regxp := regexp.MustCompile("^(.*)\\.(\\w+)$")
	matches := regxp.FindStringSubmatch(imagePath)

	if len(matches) != 3 {
		// because this method is called during template rendering, we can't return an error here
		// suppress the error and log it instead
		slog.Error("could not get image in theme: %s", imagePath)
		return imagePath
	} else {
		return fmt.Sprintf("%s-dark.%s", matches[1], matches[2])
	}
}
