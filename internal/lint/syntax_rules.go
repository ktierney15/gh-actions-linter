// syntax linting rule definitions
package lint

import (
	"fmt"
	"strings"
	"github.com/gorhill/cronexpr"
)

// Check for required fields
func NameFieldPresent(data map[string]interface{}) (*bool, string) {
	// checks if name field is present
	result := true
	errorMsg := ""
	_, exists := data["name"]

	if !exists {
		result = false
		errorMsg = "Missing 'name' field"	
	}
    return &result, errorMsg
}

func OnFieldPresent(data map[string]interface{}) (*bool, string) {
	// checks if on field is present
	result := true
	errorMsg := ""
	_, exists := data["on"]
	if !exists {
		result = false
		errorMsg = "Missing 'on' Field"
	}
    return &result, errorMsg
}

func JobsFieldPresent(data map[string]interface{}) (*bool, string) {
	// checks if job feild is present
	result := true
	errorMsg := ""
    _, exists := data["jobs"]
    if !exists {
        result = false
		errorMsg = "Missing 'jobs' field"
    }
    return &result, errorMsg
}

// workflow trigger syntax
func ValidWorkflowTrigger(data map[string]interface{}) (*bool, string) {
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
		workflowTriggersValid = false
		return &workflowTriggersValid, "Missing 'on' field"
	}

	for event := range onField {
		if !validTriggers[event] {
			workflowTriggersValid = false
			failureOutputMessage += fmt.Sprintf("%s is not a valid workflow trigger, ", event)
		}
	}
    return &workflowTriggersValid, strings.TrimSuffix(failureOutputMessage, ", ") 
}

func ValidJobStructure(data map[string]interface{}) (*bool, string) {
	// checks if the job has runs-on and steps
	jobsAreValid := true
	failureOutputMessage := ""

	jobsField, ok := data["jobs"].(map[string]interface{})
	if !ok {
		jobsAreValid = false	
		return &jobsAreValid, "Missing 'jobs' field"
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

	return &jobsAreValid, strings.TrimSuffix(failureOutputMessage, ", ") 
}

func ValidCron(data map[string]interface{}) (*bool, string) {
	// checks if schedule has a cron and cron is valid
	validCron := true
	failureOutputMessage := ""

	onField, ok := data["on"].(map[string]interface{})
	if !ok {
		validCron = false
		return &validCron, "Missing 'on' field"
	}

	scheduleField, hasSchedule := onField["schedule"]
	if !hasSchedule {
		// No schedule present, rule does not apply
		return nil, ""
	}

	scheduleList, ok := scheduleField.([]interface{})
	if !ok {
		validCron := false
		return &validCron, "'schedule' field is not a valid list"
	}

	for _, scheduleItem := range scheduleList {
		schedMap, ok := scheduleItem.(map[string]interface{})
		if !ok {
			result := false
			return &result, "Invalid structure in 'schedule' list"
		}

		cronExpr, hasCron := schedMap["cron"].(string)
		if !hasCron {
			validCron = false
			failureOutputMessage += "Missing 'cron' key in schedule, "
			continue
		}

		_, err := cronexpr.Parse(cronExpr)
		if err != nil {
			validCron = false
			failureOutputMessage += fmt.Sprintf("Invalid cron expression: %s, ", cronExpr)
		}
	}

	return &validCron, strings.TrimSuffix(failureOutputMessage, ", ") 
}

func NeedsJobExists(data map[string]interface{}) (*bool, string) {
	// makes sure that if there is a 'needs' the job thats dependent on exists
	needsJobExists := true
	failureOutputMessage := ""

	jobField, ok := data["jobs"].(map[string]interface{})
	if !ok {
		return nil, "Missing 'jobs' field"
	}

	jobNames := make(map[string]bool)
	for jobName := range jobField {
		jobNames[jobName] = true
	}

	for jobName, jobValue := range jobField {
		jobMap , ok:= jobValue.(map[string]interface{})
		if !ok {
			continue
		}

		needsValue, exists := jobMap["needs"]
		if !exists {
			continue
		}

		switch needs := needsValue.(type) {
		case string:
			if !jobNames[needs] {
				needsJobExists = false
				failureOutputMessage += fmt.Sprintf("Job '%s' depends on missing job '%s', ", jobName, needs)
			}
		case []interface{}:
			for _, need := range needs {
				needStr, ok := need.(string)
				if !ok {
					continue
				}
				if !jobNames[needStr] {
					needsJobExists = false
					failureOutputMessage += fmt.Sprintf("Job '%s' depends on missing job '%s', ", jobName, needStr)
				}
			}
		default:
			needsJobExists = false
			failureOutputMessage += fmt.Sprintf("Job '%s' has an invalid 'needs' format, ", jobName)
		}

	}

	return &needsJobExists, failureOutputMessage
}