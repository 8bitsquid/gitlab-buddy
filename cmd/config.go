package cmd

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const MASK_RUNE rune = '\u002E'

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration settings for Gitlab Buddy",
	Long: `
Manage configuration settings for Gitlab Buddy. Info to connect to Gitlab hosts (such as API Tokens) are stored in 
the '$HOME/.heb-digital/gitlab-buddy.yml'. Gitlab Buddy can create the config file for you, or you can pass in a different
config with the '--config' global flag. Config files can be simple with only API Tokens and a user-friendly host name
to use with this CLI, but can specify different URLs and API Paths per API Token.

Each API Token is considered it's own "client" with a unique human-friendly name associated with it. Each "client" can 
point to the same host, but will be treated as unique entities. This helps prevent possible cross chatter when 
entering commands into Gitlab Buddy.

So be sure you give each API Token an easily distinguishable name!
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func promptAddWhatClient() (int, string, error) {
	prompt := promptui.Select{
		Label: "What would you like to configure",
		Items: []string{
			"Gitlab Token", 
			"Local Git Repo(s) [Not yet supported]", 
			"Nothin' get me outta here."},
	
	}

	return prompt.Run()
}

func promptAddGitlabClient() (map[string]string, error) {
	c := make(map[string]string)
	token, err := promptToken("Gitlab")
	if err != nil {
		return nil, err
	}
	c["token"] = token

	name, err := promptHumanReadableName()
	if err != nil {
		return nil, err
	}
	c["name"] = name
	
	i, _, _ := promptConfirm()
	if i > 0 {
		return promptAddGitlabClient()
	}

	return c, nil
}

func promptToken(clientName string) (string, error) {
	tokenPrompt := promptui.Prompt{
		Label: clientName + " API Token",
		Mask: MASK_RUNE,
		Validate: func(s string) error {
			if err := validateNonEmpty(s, "API Token"); err != nil {
				return err
			}
			return nil
		},
	}
	return tokenPrompt.Run()
}
// TODO: Implement URL and API Path prompts
// func promptURL() (string, error) {
// 	urlPrompt := promptui.Prompt{
// 		Label: "Host URL",
// 		Default: "https://gitlab.com",
// 		Validate: func(s string) error {
// 			if err := validateNonEmpty(s, "URL"); err != nil {
// 				return err
// 			}
// 			// tool.NewURL() validates urls, and 
// 			// checks if they have Scheme and Host defined
// 			if _, err := tools.NewURL(s); err != nil {
// 				return errors.Unwrap(fmt.Errorf("%w. Must contain scheme and host (e.g., https://mydomain.com", err))
// 			}
			
// 			return nil		
// 		},
// 	}
// 	return urlPrompt.Run()
// }

// func promptAPIPath() (string, error) {
// 	apiPrompt := promptui.Prompt{
// 		Label: "API Path",
// 		Default: "/api/v4",Validate: func(s string) error {
// 			if err := validateNonEmpty(s, "API Path"); err != nil {
// 				return err
// 			}
// 			_, err := url.Parse(s)
// 			return err			
// 		},
// 	}
// 	return apiPrompt.Run()
// }

func promptHumanReadableName() (string, error){
	forbidden := `!@#$%^&*()[]{}=+,.<>/?|\;:'"~`
	namePrompt := promptui.Prompt{
		Label: "Human Friendly Name",
		Validate: func(s string) error {
			if err := validateNonEmpty(s, "Name"); err != nil {
				return err
			}

			pattern := regexp.QuoteMeta(forbidden)
			re := regexp.MustCompile("["+pattern+"]")
			if matches := re.MatchString(s); matches {
				return errors.New("Name cannot contain special characters: " + forbidden)
			}
			return nil
		},
	}
	return namePrompt.Run()
}

func promptConfirm() (int, string, error) {
	configAddPrompt := promptui.Select{
		Label: "Everything Look Good?",
		Items: []string{
			"Yup",
			"Nope, start over",
		},
	}
	return configAddPrompt.Run()
}

func validateNonEmpty(s string, label string) error {
	if s == "" {
		return errors.New(label + " is required")
	}
	return nil
}