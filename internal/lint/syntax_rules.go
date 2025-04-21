// syntax linting rule definitions
package lint

import (
	"fmt"
	"strings"
)

// Check for required fields
func NameFieldPresent(data map[string]interface{}) (bool, string) {
	// checks if name field is present
    _, exists := data["name"]
	if !exists {
		return false, "Missing 'name' field"	
	}
    return true, ""
}

func OnFieldPresent(data map[string]interface{}) (bool, string) {
	// checks if on field is present
	_, exists := data["on"]
	if !exists {
		return false, "Missing 'on' Field"
	}
    return true, ""
}

func JobsFieldPresent(data map[string]interface{}) (bool, string) {
	// checks if job feild is present
    _, exists := data["jobs"]
    if !exists {
        return false, "Missing 'jobs' field"
    }
    return true, ""
}

// workflow trigger syntax
func ValidWorkflowTrigger(data map[string]interface{}) (bool, string) {
	// checks if all workflow triggers are valid
	validTriggers := map[string]bool{
		"push":               true,
		"pull_request":       true,
		"workflow_dispatch":  true,
		"schedule":           true,
		"release":            true,
		"workflow_call":      true,
		"repository_dispatch":true,
		"issue_comment":      true,
		"check_run":          true,
	}
	workflowTriggersValid := true
	failureOutputMessage := ""

	onField, ok := data["on"].(map[string]interface{})
	if !ok {
		return false, "Missing 'on' field"
	}

	for event := range onField {
		if !validTriggers[event] {
			workflowTriggersValid = false
			failureOutputMessage += fmt.Sprintf("%s is not a valid workflow trigger, ", event)
		}
	}
    return workflowTriggersValid, strings.TrimSuffix(failureOutputMessage, ", ") 
}

func ValidJobStructure(data map[string]interface{}) (bool, string) {
	// checks if the job has runs-on and steps
	jobsAreValid := true
	failureOutputMessage := ""

	jobsField, ok := data["jobs"].(map[string]interface{})
	if !ok {
		return false, "Missing 'jobs' field"
	}

	for jobName, jobValue := range jobsField {
		jobMap, ok := jobValue.(map[string]interface{})
		if !ok {
			jobsAreValid = false
			failureOutputMessage += fmt.Sprintf("Job '%s' is not a valid yaml object, ", jobName)
		}

		if _, exists := jobMap["runs-on"]; ! exists {
			jobsAreValid = false
			failureOutputMessage += fmt.Sprintf("Job '%s' is missing a runs-on value, ", jobName)
		}

		steps, stepsExists := jobMap["steps"]
		if !stepsExists {
			jobsAreValid = false
			failureOutputMessage += fmt.Sprintf("Job '%s' is missing steps value, ", jobName)
		} else {
			stepList, ok := steps.([]interface{})
			if !ok || len(stepList) == 0 {
				jobsAreValid = false
				failureOutputMessage += fmt.Sprintf("Job '%s' is missing steps, ", jobName)
			}
		}
	}

	return jobsAreValid, strings.TrimSuffix(failureOutputMessage, ", ") 
}