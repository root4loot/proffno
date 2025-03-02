package proffno

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"
	"unicode"
)

type Result struct {
	RawInput        string  `json:"rawInput"`
	Name            string  `json:"name"`
	OwnedPercentage float64 `json:"ownedPercentage"`
}

type Results struct {
	Results []Result `json:"results"`
}

func findOrgInfo(query string, queryIsOrgName bool) (orgInfo string, err error) {
	query = strings.ReplaceAll(query, " ", "+")

	buildID, err := getBuildID()
	if err != nil {
		return "", err
	}

	response, err := http.Get("https://proff.no/_next/data/" + buildID + "/search.json?q=" + query)
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

func GetSubsidiaries(orgName string, greaterThanPercentage float64) ([]Result, error) {
	return getSubsidiaries(orgName, 1, greaterThanPercentage)
}

func getSubsidiaries(orgName string, level int, greaterThanPercentage float64) ([]Result, error) {
	if level <= 0 {
		return nil, nil
	}

	orgName = normalizeOrganizationName(orgName)
	subs, err := collectSubsidiaries(orgName)
	if err != nil {
		return nil, err
	}

	var results []Result
	for _, sub := range subs {
		if sub.OwnedPercentage > greaterThanPercentage {
			results = append(results, Result{
				Name:            sub.Name,
				OwnedPercentage: sub.OwnedPercentage,
			})
			// Recursive call with decreased level
			subSubs, err := getSubsidiaries(sub.Name, level-1, greaterThanPercentage)
			if err != nil {
				return nil, err
			}
			results = append(results, subSubs...)
		}
	}

	return results, nil
}
func getBuildID() (string, error) {
	resp, err := http.Get("https://proff.no/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`"buildId":"(.*?)"`)
	matches := re.FindStringSubmatch(string(body))

	if len(matches) < 2 {
		return "", nil
	}

	return matches[1], nil
}

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

	response, err := http.Get("https://proff.no/api/company/legal/" + orgNR + "/corporateStructure")
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

func normalizeOrganizationName(name string) string {
	words := strings.Fields(name)

	for i, word := range words {
		for _, v := range word {
			u := string(unicode.ToUpper(v))
			word = u + word[len(u):]
			break
		}

		if (strings.ToLower(word) == "as" || strings.ToLower(word) == "asa") && i == len(words)-1 {
			word = strings.ToUpper(word)
		}

		words[i] = word
	}

	return strings.Join(words, " ")
}

func isNumeric(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
