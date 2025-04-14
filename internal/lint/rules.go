// linting rule definitions
package lint

// import (
//     "fmt"
//     "gopkg.in/yaml.v2"
// )

type Rule struct {
    Name        string
    Description string
	Weight      int    // 1-10
	// Consider definging Severity and Category types with only those options
	Severity 	string // "info", "warning", "critical"
	Category 	string // "syntax", "performance", "security"
    Check       func(map[string]interface{}) bool
}

func NameFieldPresent(data map[string]interface{}) bool {
    _, exists := data["name"]
    return exists
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
    }
}
