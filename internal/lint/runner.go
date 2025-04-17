// orchestrates applying rules to your files
package lint

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "io/ioutil"
)

func Run(fileName string) error {
	// read the yaml file
	data, err := ioutil.ReadFile(fileName)
    if err != nil {
        return fmt.Errorf("error reading file %s: %v", fileName, err)
    }

	// parse yaml file into a map
	parsedYaml := make(map[string]interface{})
    err = yaml.Unmarshal(data, &parsedYaml)
	
    if err != nil {
		// fmt.Println(err)
        return fmt.Errorf("error parsing YAML file %s: %v", fileName, err)
    }

	// initialize rules and run them against the yaml
	rules := InitializeRules()
	totalWeight := 0
	failedWeightSum := 0

	fmt.Println("----------------------------------------------------------------------")

	for _, rule := range rules {
		passed, message := rule.Check(parsedYaml)
		severityInfo := SeverityMap[rule.Severity]
		totalWeight += severityInfo.Weight

		if !passed {
			failedWeightSum += severityInfo.Weight

			fmt.Printf("%s [%s] %s\n   ↳ %s\n", severityInfo.Emoji , rule.Severity, rule.Name, message)
			fmt.Println("----------------------------------------------------------------------")
		}
	}

	// calculate and print score
	score := float64(totalWeight - failedWeightSum) / float64(totalWeight) * 100
	fmt.Printf("\n🏁 Lint Score: %.2f%% 🏁 \n\n", score)
	
    return nil // if no errors
}