package bot_test

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
)

func TestAccBotChannelsRegistration(t *testing.T) {
	// NOTE: this is a combined test rather than separate split out tests due to
	// Azure only being able provision against one app id at a time
	acceptance.RunTestsInSequence(t, map[string]map[string]func(t *testing.T){
		"basic": {
			"basic":    testAccBotChannelsRegistration_basic,
			"update":   testAccBotChannelsRegistration_update,
			"complete": testAccBotChannelsRegistration_complete,
		},
		"connection": {
			"basic":    testAccBotConnection_basic,
			"complete": testAccBotConnection_complete,
		},
		"channel": {
			"alexaBasic":                     testAccBotChannelAlexa_basic,
			"alexaUpdate":                    testAccBotChannelAlexa_update,
			"alexaRequiresImport":            testAccBotChannelAlexa_requiresImport,
			"slackBasic":                     testAccBotChannelSlack_basic,
			"slackUpdate":                    testAccBotChannelSlack_update,
			"smsBasic":                       testAccBotChannelSMS_basic,
			"smsRequiresImport":              testAccBotChannelSMS_requiresImport,
			"msteamsBasic":                   testAccBotChannelMsTeams_basic,
			"msteamsUpdate":                  testAccBotChannelMsTeams_update,
			"directlineBasic":                testAccBotChannelDirectline_basic,
			"directlineComplete":             testAccBotChannelDirectline_complete,
			"directlineUpdate":               testAccBotChannelDirectline_update,
			"directLineSpeechBasic":          testAccBotChannelDirectLineSpeech_basic,
			"directLineSpeechComplete":       testAccBotChannelDirectLineSpeech_complete,
			"directLineSpeechUpdate":         testAccBotChannelDirectLineSpeech_update,
			"directLineSpeechRequiresImport": testAccBotChannelDirectLineSpeech_requiresImport,
			"facebookBasic":                  testAccBotChannelFacebook_basic,
			"facebookUpdate":                 testAccBotChannelFacebook_update,
			"facebookRequiresImport":         testAccBotChannelFacebook_requiresImport,
			"lineBasic":                      testAccBotChannelLine_basic,
			"lineComplete":                   testAccBotChannelLine_complete,
			"lineUpdate":                     testAccBotChannelLine_update,
			"lineRequiresImport":             testAccBotChannelLine_requiresImport,
			"webchatBasic":                   testAccBotChannelWebChat_basic,
			"webchatComplete":                testAccBotChannelWebChat_complete,
			"webchatUpdate":                  testAccBotChannelWebChat_update,
			"webchatRequiresImport":          testAccBotChannelWebChat_requiresImport,
		},
		"web_app": {
			"basic":    testAccBotWebApp_basic,
			"update":   testAccBotWebApp_update,
			"complete": testAccBotWebApp_complete,
		},
	})
}
