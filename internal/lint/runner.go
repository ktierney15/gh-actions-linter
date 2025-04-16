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
	fmt.Println("----------------------------------------------------------------------")

	for _, rule := range rules {
		passed, message := rule.Check(parsedYaml)
		if !passed {
			emoji := "❌" // fallback
			if val, ok := SeverityMap[rule.Severity]; ok {
				emoji = val.Emoji
			}
			fmt.Printf("%s [%s] %s\n   ↳ %s\n", emoji, rule.Severity, rule.Name, message)
			fmt.Println("----------------------------------------------------------------------")
		}
	}
	
    return nil // if no errors
}