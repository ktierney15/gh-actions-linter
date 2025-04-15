// rule struct and shared logic
package lint

type Rule struct {
    Name        string
    Description string
	Weight      int    // 1-10
	// Consider defining Severity and Category types with only those options
	Severity 	string // "info", "warning", "critical"
	Category 	string // "syntax", "performance", "security"
    Check       func(map[string]interface{}) bool
}

func InitializeRules() []Rule {
    return []Rule{
        {
            Name:        "NameFieldPresent",
            Description: "Check if 'name' field exists in the GitHub Actions YAML",
			Weight: 	 10,
			Severity:	 "critical",
			Category:	 "syntax",
            Check:       NameFieldPresent,
        },
		{
			Name:        "OnFieldPresent",
            Description: "Check if 'on' field exists in the GitHub Actions YAML",
			Weight: 	 10,
			Severity:	 "critical",
			Category:	 "syntax",
            Check:       OnFieldPresent,
		},
		{
			Name:        "JobsFieldPresent",
            Description: "Check if 'jobs' field exists in the GitHub Actions YAML",
			Weight: 	 10,
			Severity:	 "critical",
			Category:	 "syntax",
            Check:       JobsFieldPresent,
		},
		{
			Name:        "ValidWorkflowTrigger",
            Description: "Check if 'on' field has valid triggers (workflow_dispatch, push, pull_request, schedule, etc...)",
			Weight: 	 10,
			Severity:	 "critical",
			Category:	 "syntax",
            Check:       ValidWorkflowTrigger,
		},
		{
			Name:        "ValidJobNames",
            Description: "Check if job names follow the following rules: kebab-case, lowercase letters, numbers, dashes, underscores",
			Weight: 	 2,
			Severity:	 "low",
			Category:	 "style",
            Check:       ValidJobNames,
		},
    }
}
