package survey

import "github.com/AlecAivazis/survey/v2"

type PromptAsk interface {
	Select(message string, options []string) (int, string, error)
	MultiSelect(message string, options []string) ([]int, []string, error)
}

type ask struct{}

// GetAsk return ask that implements PromptAsk
func GetAsk() PromptAsk {
	return ask{}
}

// Select create a survey to select one option
func (a ask) Select(message string, options []string) (int, string, error) {
	var (
		selIndex int
		err      error
	)
	err = survey.AskOne(
		&survey.Select{
			Message: message,
			Options: options,
		},
		&selIndex,
	)

	return selIndex, options[selIndex], err
}

// MultiSelect create a survey to select multiple values from an option
func (a ask) MultiSelect(message string, options []string) ([]int, []string, error) {
	var (
		selIndexes []int
		selOptions []string
		err        error
	)

	err = survey.AskOne(
		&survey.MultiSelect{
			Message: message,
			Options: options,
		},
		&selIndexes,
	)

	for _, index := range selIndexes {
		selOptions = append(selOptions, options[index])
	}

	return selIndexes, selOptions, err
}
