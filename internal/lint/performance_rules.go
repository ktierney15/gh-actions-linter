// rules that deter slowing the performance of a pipeline
package lint

import (
	"fmt"
	"strings"
	"encoding/json"
)

func RedundantSteps(data map[string]interface{}) (*bool, string) {
	// checks if there are any duplicate steps in a job
	noRedundantSteps := true
	failureOutputMessage := ""

	jobField, ok := data["jobs"].(map[string]interface{})
	if !ok {
	}

	for jobName, jobValue := range jobField {
		jobMap , ok:= jobValue.(map[string]interface{})
		if !ok {
			continue
		}

		steps, ok := jobMap["steps"].([]interface{})
		if !ok {
			continue
		}

		seenSteps := make(map[string]bool)

		for _, step := range steps {
			stepMap, ok := step.(map[string]interface{})
			if !ok {
				continue
			}
			stepBytes, err := json.Marshal(stepMap)
			if err != nil {
				continue
			}

			stepKey := string(stepBytes)
			if seenSteps[stepKey] {
				noRedundantSteps = false
				failureOutputMessage += fmt.Sprintf("Job '%s' has a redundant step: %s, ", jobName, stepKey)
			} else {
				seenSteps[stepKey] = true
			}
		}
	}

	return &noRedundantSteps, strings.TrimSuffix(failureOutputMessage, ", ")
}


