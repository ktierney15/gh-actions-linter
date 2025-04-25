// rule functions of the style type. These rules are generally lower severity and more for best practice
package lint

import (
	"fmt"
	"strings"
	"regexp"
)

func ValidJobNames(data map[string]interface{}) (*bool, string) {
	// Checks if all job names are kebab-case, lowercase letters, numbers, dashes, underscores
	validJobRegex := regexp.MustCompile(`^[a-z0-9\-_]+$`) 
	jobNamesValid := true
	failureOutputMessage := ""

	jobField, ok := data["jobs"].(map[string]interface{})
	if !ok {
		jobNamesValid = false
		return &jobNamesValid, "Missing 'jobs' field"
	}

	for job := range jobField {
		if !validJobRegex.MatchString(job) {
			jobNamesValid = false
			failureOutputMessage += fmt.Sprintf("job name '%s' should be lowercase kebab-case, ", job)
		}
	}

	return &jobNamesValid, strings.TrimSuffix(failureOutputMessage, ", ")
}

func EachStepHasName(data map[string]interface{}) (*bool, string) {
	// checks if each step has a name
	eachStepHasName := true
	failureOutputMessage := ""

	jobField, ok := data["jobs"].(map[string]interface{})
	if !ok {
		eachStepHasName = false
		return &eachStepHasName, "Missing 'jobs' field"
	}

	for jobName, jobValue := range jobField {
		jobMap , ok:= jobValue.(map[string]interface{})
		if !ok {
			eachStepHasName = false
			failureOutputMessage += fmt.Sprintf("Job '%s' is not a valid yaml object, ", jobName)
		}
		
		steps, ok := jobMap["steps"].([]interface{})
		if !ok {
			eachStepHasName = false
			failureOutputMessage += fmt.Sprintf("Job %s does not have steps, ", jobName)
		} else {
			for i, step := range steps {
				stepMap := step.(map[string]interface{})
	
				name, hasName := stepMap["name"]
				if !hasName || name == "" {
					eachStepHasName = false
					failureOutputMessage += fmt.Sprintf("Step %d in job %s does not have a name, ", i+1, jobName)
				}
			}
		}
	}

	return &eachStepHasName, strings.TrimSuffix(failureOutputMessage, ", ")
}

func UsingActionVersion(data map[string]interface{}) (*bool, string) {
	// checks if any actions are pointing to a branch or latest
	isUsingOnlyVersions := true
	failureOutputMessage := ""

	jobField, ok := data["jobs"].(map[string]interface{})
	if !ok {
		// pass because there are no 'uses' if there is nothing in the job
	}

	for jobName, jobValue := range jobField {
		jobMap , ok:= jobValue.(map[string]interface{})
		if !ok {
			continue
		}

		steps, ok := jobMap["steps"].([]interface{})
		if !ok {
			// pass because there are no 'uses' if there is nothing in the steps
			continue
		} else {
			for _, step := range steps {
				stepMap, ok := step.(map[string]interface{})
				if !ok {
					continue
				}
				if usesVal, ok := stepMap["uses"].(string); ok {
					if !strings.Contains(usesVal, "@") || strings.HasSuffix(usesVal, "@main") || strings.HasSuffix(usesVal, "@master") || strings.HasSuffix(usesVal, "@latest") {
						isUsingOnlyVersions = false
						failureOutputMessage += fmt.Sprintf("Step in job '%s' uses unversioned action: '%s', ", jobName, usesVal)
					}
				}
			}
		}

	}

	return &isUsingOnlyVersions, strings.TrimSuffix(failureOutputMessage, ", ")
}

func NoLongRunCommands(data map[string]interface{}) (*bool, string) {
	// determines if a run command is too long (and should be put in a script)
	noLongRunCommands := true
	failureOutputMessage := ""

	jobField, ok := data["jobs"].(map[string]interface{})
	if !ok {
		// pass because there are no 'uses' if there is nothing in the job
	}


	for jobName, jobValue := range jobField {
		jobMap , ok:= jobValue.(map[string]interface{})
		if !ok {
			continue
		}

		steps, ok := jobMap["steps"].([]interface{})
		if !ok {
			continue
		} else {
			for i, step := range steps {
				stepMap, ok := step.(map[string]interface{})
				if !ok {
					continue
				}
				if runsVal, ok := stepMap["run"].(string); ok {
					if len(runsVal) > 400 {
						noLongRunCommands = false
						failureOutputMessage += fmt.Sprintf("Step %d in job '%s' has a run command that is over 400 characters, ", i+1, jobName)
					}
				}
			}
		}
	}

	return &noLongRunCommands, strings.TrimSuffix(failureOutputMessage, ", ")
}

func InputsHaveDescriptions(data map[string]interface{}) (*bool, string) {
	// checks to make sure input values have descriptions
	inputsHaveDescriptions := true
	failureOutputMessage := ""

	onField, ok := data["on"].(map[string]interface{})
	if !ok {
		inputsHaveDescriptions = false
		return &inputsHaveDescriptions, "Missing 'on' field"
	}

	dispatchField, hasDispatchField := onField["workflow_dispatch"]
	if !hasDispatchField {
		return nil, "No workflow_dispatch, skipping rule"
	}

	dispatchMap, ok := dispatchField.(map[string]interface{})
	if !ok {
		return nil, "no inputs, skipping rule"
	}

	inputs, ok := dispatchMap["inputs"].(map[string]interface{})
	if !ok {
		return nil, "no inputs, skipping rule"
	}

	for inputName, inputValue := range inputs {
		inputMap, ok := inputValue.(map[string]interface{})
		if !ok {
			continue
		}
		if _, hasDescription := inputMap["description"]; !hasDescription {
			inputsHaveDescriptions = false
			failureOutputMessage += fmt.Sprintf("Input '%s' is missing a description, ", inputName)
		}
	}

	return &inputsHaveDescriptions, strings.TrimSuffix(failureOutputMessage, ", ")
}