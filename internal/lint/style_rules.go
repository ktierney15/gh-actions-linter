// rule functions of the style type. These rules are generally lower severity and more for best practice
package lint

import (
	"fmt"
	"strings"
	"regexp"
)

func ValidJobNames(data map[string]interface{}) (bool, string) {
	// Checks if all job names are kebab-case, lowercase letters, numbers, dashes, underscores
	validJobRegex := regexp.MustCompile(`^[a-z0-9\-_]+$`) 
	jobNamesValid := true
	failureOutputMessage := ""

	jobField, ok := data["jobs"].(map[string]interface{})
	if !ok {
		return false, "Missing 'jobs' field"
	}

	for job := range jobField {
		if !validJobRegex.MatchString(job) {
			jobNamesValid = false
			failureOutputMessage += fmt.Sprintf("job name '%s' should be lowercase kebab-case, ", job)
		}
	}

	return jobNamesValid, strings.TrimSuffix(failureOutputMessage, ", ")
}

// func UniqueJobName(data map[string]interface{}) (bool, string) {
// 	// checks if any jobs have the same name THIS FUNCTION IS NOT NEEDED BECAUSE THE YAML CANT BE PARSED IF ITS TRUE
// 	jobNamesValid := true
// 	failureOutputMessage := ""
// 	jobSet := make(map[string]bool)

// 	jobField, ok := data["jobs"].(map[string]interface{})
// 	if !ok {
// 		return false, "Missing 'jobs' field"
// 	}

// 	for job := range jobField {
// 		if jobSet[job] {
// 			jobNamesValid = false
// 			failureOutputMessage += fmt.Sprintf("multiple jobs named '%s', ", job)
// 		}
// 		jobSet[job] = true
// 	}

// 	return jobNamesValid, strings.TrimSuffix(failureOutputMessage, ", ")
// }

func EachStepHasName(data map[string]interface{}) (bool, string) {
	// checks if each step has a name
	eachStepHasName := true
	failureOutputMessage := ""

	jobField, ok := data["jobs"].(map[string]interface{})
	if !ok {
		return false, "Missing 'jobs' field"
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

	return eachStepHasName, strings.TrimSuffix(failureOutputMessage, ", ")
}