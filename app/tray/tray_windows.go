package tray

import (
	"github.com/CopperCarbonateCollective/ollama-DreamingMetal/app/tray/commontray"
	"github.com/CopperCarbonateCollective/ollama-DreamingMetal/app/tray/wintray"
)

func InitPlatformTray(icon, updateIcon []byte) (commontray.OllamaTray, error) {
	return wintray.InitTray(icon, updateIcon)
}
