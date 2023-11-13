package proffNO

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"
	"unicode"
)

type Result struct {
	Name            string  `json:"name"`            // Name of subsidiary
	OwnedPercentage float64 `json:"ownedPercentage"` // Percentage of ownership
}

type Results struct {
	Results []Result `json:"results"`
}

// findOrgInfo finds the organization number or name of the company
func findOrgInfo(query string, queryIsOrgName bool) (orgInfo string, err error) {
	query = strings.ReplaceAll(query, " ", "+")

	// Get the build ID
	buildID, err := getBuildID()
	if err != nil {
		return "", err
	}

	response, err := http.Get("https://beta.proff.no/_next/data/" + buildID + "/search.json?q=" + query)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var structure SearchStructure
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&structure); err != nil {
		return "", err
	}

	if len(structure.PageProps.HydrationData.SearchStore.Companies.Companies) > 0 {
		if queryIsOrgName {
			orgInfo = structure.PageProps.HydrationData.SearchStore.Companies.Companies[0].Orgnr
		} else {
			orgInfo = structure.PageProps.HydrationData.SearchStore.Companies.Companies[0].Name
		}
	}

	return orgInfo, nil
}

// GetSubsidiaries gets the subsidiaries of a company
func GetSubsidiaries(orgName string, level int, greaterThanPercentage float64) (results []Result, err error) {
	if level <= 0 {
		return results, nil
	}

	subs, _ := collectSubsidiaries(orgName)

	for _, sub := range subs {
		if sub.OwnedPercentage > greaterThanPercentage {
			result := Result{
				Name:            sub.Name,
				OwnedPercentage: sub.OwnedPercentage,
			}
			results = append(results, result)

			// Recursively call the function for sub-subsidiaries
			subSubs, err := GetSubsidiaries(sub.Name, level-1, greaterThanPercentage)
			if err != nil {
				return nil, err
			}
			results = append(results, subSubs...)
		}
	}

	return results, nil
}

func getBuildID() (string, error) {
	// Make a GET request
	resp, err := http.Get("https://beta.proff.no/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Regular expression to find the build ID
	re := regexp.MustCompile(`"buildId":"(.*?)"`)
	matches := re.FindStringSubmatch(string(body))

	if len(matches) < 2 {
		return "", nil // No match found
	}

	// Return the found build ID
	return matches[1], nil
}

// collectSubsidiaries collects the subsidiaries of a company
func collectSubsidiaries(query string) (results []Result, err error) {
	var structure CorporateStructure
	var orgNR, orgName string

	if isNumeric(query) {
		orgNR = query
		orgName, err = findOrgInfo(query, true)
	} else {
		orgName = query
		orgNR, err = findOrgInfo(query, true)
	}

	response, err := http.Get("https://beta.proff.no/api/company/legal/" + orgNR + "/corporateStructure")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&structure); err != nil {
		return nil, err
	}

	if structure.Tree.Name == orgName {
		for _, level1 := range structure.Tree.Sub {
			results = append(results, Result{Name: level1.Name, OwnedPercentage: level1.OwnedPercentage})
		}
	} else {
		for _, level1 := range structure.Tree.Sub {
			if level1.Name == orgName {
				for _, level2 := range level1.Sub {
					results = append(results, Result{Name: level2.Name, OwnedPercentage: level2.OwnedPercentage})
				}
			}
		}

		for _, level1 := range structure.Tree.Sub {
			for _, level2 := range level1.Sub {
				if level2.Name == orgName {
					for _, level3 := range level2.Sub {
						results = append(results, Result{Name: level3.Name, OwnedPercentage: level3.OwnedPercentage})
					}
				}
			}
		}

		for _, level1 := range structure.Tree.Sub {
			for _, level2 := range level1.Sub {
				for _, level3 := range level2.Sub {
					if level3.Name == orgName {
						for _, level4 := range level3.Sub {
							if sub4Map, ok := level4.(map[string]interface{}); ok {
								var name string
								var ownedPercentage float64
								for key, value := range sub4Map {
									switch key {
									case "name":
										name, _ = value.(string)
									case "ownedPercentage":
										ownedPercentage, _ = value.(float64)

									}
								}
								results = append(results, Result{Name: name, OwnedPercentage: ownedPercentage})
							}
						}
					}
				}
			}
		}
	}
	return
}

// isNumeric checks if a string is numeric
func isNumeric(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
