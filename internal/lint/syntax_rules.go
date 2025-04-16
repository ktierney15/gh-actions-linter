// syntax linting rule definitions
package lint

import (
	"fmt"
	"strings"
// 	"regexp"
)

// Check for required fields
func NameFieldPresent(data map[string]interface{}) (bool, string) {
    _, exists := data["name"]
	if !exists {
		return false, "Missing 'name' field"	
	}
    return true, ""
}

func OnFieldPresent(data map[string]interface{}) (bool, string) {
	_, exists := data["on"]
	if !exists {
		return false, "Missing 'on' Field"
	}
    return true, ""
}

func JobsFieldPresent(data map[string]interface{}) (bool, string) {
    _, exists := data["jobs"]
    if !exists {
        return false, "Missing 'jobs' field"
    }
    return true, ""
}

// workflow trigger syntax
func ValidWorkflowTrigger(data map[string]interface{}) (bool, string) {
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
		return false, "missing 'on' field"
	}

	for event := range onField {
		if !validTriggers[event] {
			workflowTriggersValid = false
			failureOutputMessage += fmt.Sprintf("%s is not a valid workflow trigger, ", event)
		}
	}
    return workflowTriggersValid, strings.TrimSuffix(failureOutputMessage, ", ") 
}


// job syntax
// func ValidJobNames(data map[string]interface{}) bool {
// 	validJobRegex := regexp.MustCompile(`^[a-z0-9\-_]+$`) // kebab-case, lowercase letters, numbers, dashes, underscores
// 	jobNamesValid := true

// 	jobField, ok := data["jobs"].(map[string]interface{})
// 	if !ok {
// 		fmt.Printf("not ok")
// 		return false
// 	}

// 	for job := range jobField {
// 		if !validJobRegex.MatchString(job) {
// 			fmt.Printf("🔔Style warning: job name '%s' should be lowercase kebab-case \n", job)
// 			jobNamesValid = false
// 		}
// 	}

// 	return jobNamesValid
// }