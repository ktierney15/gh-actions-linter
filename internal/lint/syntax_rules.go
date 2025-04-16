// syntax linting rule definitions
package lint

import (
	"fmt"
	"regexp"
)

// Check for required fields
func NameFieldPresent(data map[string]interface{}) bool {
    _, exists := data["name"]
    return exists
}

func OnFieldPresent(data map[string]interface{}) bool {
	_, exists := data["on"]
    return exists
}

func JobsFieldPresent(data map[string]interface{}) bool {
	_, exists := data["jobs"]
    return exists
}

// workflow trigger syntax
func ValidWorkflowTrigger(data map[string]interface{}) bool {
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

	onField, ok := data["on"].(map[string]interface{})
	if !ok {
		return false
	}

	for event := range onField {
		if !validTriggers[event] {
			fmt.Printf("🛑Invalid trigger found: '%s'\n", event)
			workflowTriggersValid = false
		}
	}
    return workflowTriggersValid
}


// job syntax
func ValidJobNames(data map[string]interface{}) bool {
	validJobRegex := regexp.MustCompile(`^[a-z0-9\-_]+$`) // kebab-case, lowercase letters, numbers, dashes, underscores
	jobNamesValid := true

	jobField, ok := data["jobs"].(map[string]interface{})
	if !ok {
		fmt.Printf("not ok")
		return false
	}

	for job := range jobField {
		if !validJobRegex.MatchString(job) {
			fmt.Printf("🔔Style warning: job name '%s' should be lowercase kebab-case \n", job)
			jobNamesValid = false
		}
	}

	return jobNamesValid
}