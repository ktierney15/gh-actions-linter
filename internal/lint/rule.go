// rule struct and shared logic
package lint

type Rule struct {
    Name        string
    Description string
	Severity 	string // "info", "warning", "critical"
	Category 	string // "syntax", "performance", "security"
    Check       func(map[string]interface{}) (*bool, string)
}

func InitializeRules() []Rule {
    return []Rule{
        {
            Name:        "NameFieldPresent",
            Description: "Check if 'name' field exists in the GitHub Actions YAML",
			Severity:	 "critical",
			Category:	 "syntax",
            Check:       NameFieldPresent,
        },
		{
			Name:        "OnFieldPresent",
            Description: "Check if 'on' field exists in the GitHub Actions YAML",
			Severity:	 "critical",
			Category:	 "syntax",
            Check:       OnFieldPresent,
		},
		{
			Name:        "JobsFieldPresent",
            Description: "Check if 'jobs' field exists in the GitHub Actions YAML",
			Severity:	 "critical",
			Category:	 "syntax",
            Check:       JobsFieldPresent,
		},
		{
			Name:        "ValidWorkflowTrigger",
            Description: "Check if 'on' field has valid triggers (workflow_dispatch, push, pull_request, schedule, etc...)",
			Severity:	 "critical",
			Category:	 "syntax",
            Check:       ValidWorkflowTrigger,
		},
		{
			Name:        "ValidJobStructure",
            Description: "Checks if the job has runs-on and steps",
			Severity:	 "critical",
			Category:	 "syntax",
            Check:       ValidJobStructure,
		},
		{
			Name:        "ValidCron",
            Description: "Checks if schedule has a cron and cron is valid",
			Severity:	 "critical",
			Category:	 "syntax",
            Check:       ValidCron,
		},
		{
			Name:        "NeedsJobExists",
            Description: "Makes sure that if there is a 'needs' the job thats dependent on exists",
			Severity:	 "critical",
			Category:	 "syntax",
            Check:       NeedsJobExists,
		},
		{
			Name:        "HasRunsOrUses",
            Description: "Checks if all steps have either 'runs' or uses'",
			Severity:	 "critical",
			Category:	 "syntax",
            Check:       HasRunsOrUses,
		},
		{
			Name:        "UsingActionVersion",
            Description: "Checks if any actions are pointing to a branch or latest",
			Severity:	 "medium",
			Category:	 "style",
            Check:       UsingActionVersion,
		},
		{
			Name:        "RedundantSteps",
            Description: "Checks if there are any duplicate steps in a job",
			Severity:	 "medium",
			Category:	 "style",
            Check:       RedundantSteps,
		},
		{
			Name:        "ValidJobNames",
            Description: "Check if job names follow the following rules: kebab-case, lowercase letters, numbers, dashes, underscores",
			Severity:	 "low",
			Category:	 "style",
            Check:       ValidJobNames,
		},
		{
			Name:        "EachStepHasName",
            Description: "Check if all steps have a name",
			Severity:	 "low",
			Category:	 "style",
            Check:       EachStepHasName,
		},
    }
}

