// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/VetalRacer/gtdlib/tdlib"
)

// SearchBackground Searches for a background by its name
// @param name The name of the background
func (client *Client) SearchBackground(name string) (*tdlib.Background, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "searchBackground",
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var background tdlib.Background
	err = json.Unmarshal(result.Raw, &background)
	return &background, err

}

// SetBackground Changes the background selected by the user; adds background to the list of installed backgrounds
// @param background The input background to use, null for filled backgrounds
// @param typeParam Background type; null for default background. The method will return error 404 if type is null
// @param forDarkTheme True, if the background is chosen for dark theme
func (client *Client) SetBackground(background tdlib.InputBackground, typeParam tdlib.BackgroundType, forDarkTheme bool) (*tdlib.Background, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":          "setBackground",
		"background":     background,
		"type":           typeParam,
		"for_dark_theme": forDarkTheme,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var backgroundDummy tdlib.Background
	err = json.Unmarshal(result.Raw, &backgroundDummy)
	return &backgroundDummy, err

}
