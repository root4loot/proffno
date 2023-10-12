package proffNO

type SearchStructure struct {
	PageProps struct {
		RolePersons struct {
			BusinessPersons []interface{} `json:"businessPersons"`
			Hits            int           `json:"hits"`
			CurrentPage     int           `json:"currentPage"`
			Pages           int           `json:"pages"`
		} `json:"rolePersons"`
		CompaniesByName struct {
			Companies   []interface{} `json:"companies"`
			Hits        int           `json:"hits"`
			Pages       int           `json:"pages"`
			CurrentPage int           `json:"currentPage"`
			Filters     []struct {
				DisplayName string        `json:"displayName"`
				Name        interface{}   `json:"name"`
				FilterType  string        `json:"filterType"`
				Filters     []interface{} `json:"filters"`
			} `json:"filters"`
			RelatedIndustryNames []interface{} `json:"relatedIndustryNames"`
		} `json:"companiesByName"`
		HydrationData struct {
			SearchStore struct {
				InactiveCompanies struct {
					Hits        int           `json:"hits"`
					Pages       int           `json:"pages"`
					CurrentPage int           `json:"currentPage"`
					Companies   []interface{} `json:"companies"`
				} `json:"inactiveCompanies"`
				Shareholders struct {
					Hits        int `json:"hits"`
					Pages       int `json:"pages"`
					CurrentPage int `json:"currentPage"`
					Entities    []struct {
						ID                 string `json:"id"`
						Owners             int    `json:"owners"`
						IndirectOwners     int    `json:"indirectOwners"`
						OwnerShips         int    `json:"ownerShips"`
						IndirectOwnerShips int    `json:"indirectOwnerShips"`
						TotalShares        int    `json:"totalShares"`
						Name               string `json:"name"`
						OrganisationNumber string `json:"organisationNumber"`
						ZipCode            string `json:"zipCode"`
						Location           string `json:"location"`
						EntityType         string `json:"entityType"`
						Country            string `json:"country"`
					} `json:"entities"`
				} `json:"shareholders"`
				Companies struct {
					Companies []struct {
						MainOffice          interface{} `json:"mainOffice"`
						Name                string      `json:"name"`
						Orgnr               string      `json:"orgnr"`
						BusinessUnitID      interface{} `json:"businessUnitId"`
						CompanyID           string      `json:"companyId"`
						ListingID           string      `json:"listingId"`
						CustomerID          interface{} `json:"customerId"`
						Phone               string      `json:"phone"`
						Phone2              interface{} `json:"phone2"`
						Mobile              string      `json:"mobile"`
						Mobile2             interface{} `json:"mobile2"`
						FaxNumber           interface{} `json:"faxNumber"`
						HomePage            interface{} `json:"homePage"`
						Email               interface{} `json:"email"`
						MarketingProtection bool        `json:"marketingProtection"`
						Industries          []struct {
							Code        string      `json:"code"`
							Name        string      `json:"name"`
							Description interface{} `json:"description"`
							Style       interface{} `json:"style"`
							CompanyID   string      `json:"companyId"`
							SalesRank   interface{} `json:"salesRank"`
						} `json:"industries"`
						Location struct {
							CountryPart  string `json:"countryPart"`
							County       string `json:"county"`
							Municipality string `json:"municipality"`
							Coordinates  []struct {
								CoordinateSystem string  `json:"coordinateSystem"`
								Xcoordinate      float64 `json:"xcoordinate"`
								Ycoordinate      float64 `json:"ycoordinate"`
							} `json:"coordinates"`
						} `json:"location"`
						CurrentIndustry struct {
							Code        string      `json:"code"`
							Name        string      `json:"name"`
							Description interface{} `json:"description"`
							CompanyID   interface{} `json:"companyId"`
						} `json:"currentIndustry"`
						VisitorAddress struct {
							AddressLine    string      `json:"addressLine"`
							BoxAddressLine interface{} `json:"boxAddressLine"`
							ZipCode        string      `json:"zipCode"`
							PostPlace      string      `json:"postPlace"`
						} `json:"visitorAddress"`
						PostalAddress struct {
							AddressLine    string      `json:"addressLine"`
							BoxAddressLine interface{} `json:"boxAddressLine"`
							ZipCode        string      `json:"zipCode"`
							PostPlace      string      `json:"postPlace"`
						} `json:"postalAddress"`
						Description                    interface{} `json:"description"`
						AdvertType                     interface{} `json:"advertType"`
						Logo                           interface{} `json:"logo"`
						SearchResultImage              interface{} `json:"searchResultImage"`
						LegalName                      string      `json:"legalName"`
						Revenue                        string      `json:"revenue"`
						Currency                       string      `json:"currency"`
						Profit                         string      `json:"profit"`
						CompanyAccountsLastUpdatedDate string      `json:"companyAccountsLastUpdatedDate"`
						Employees                      string      `json:"employees"`
						ContactPerson                  struct {
							Type           string `json:"type"`
							Name           string `json:"name"`
							Role           string `json:"role"`
							ID             string `json:"id"`
							BirthDate      string `json:"birthDate"`
							BusinessPerson bool   `json:"businessPerson"`
						} `json:"contactPerson"`
						StatusRemarks []interface{} `json:"statusRemarks"`
						Certificates  []interface{} `json:"certificates"`
					} `json:"companies"`
					Hits        int `json:"hits"`
					Pages       int `json:"pages"`
					CurrentPage int `json:"currentPage"`
					Filters     []struct {
						DisplayName string `json:"displayName"`
						Name        string `json:"name"`
						FilterType  string `json:"filterType"`
						Filters     []struct {
							DisplayName string `json:"displayName"`
							Hits        int    `json:"hits"`
						} `json:"filters"`
					} `json:"filters"`
					RelatedIndustryNames []string `json:"relatedIndustryNames"`
				} `json:"companies"`
				Query              string      `json:"query"`
				ProffIndustry      interface{} `json:"proffIndustry"`
				CountryPartFilter  interface{} `json:"countryPartFilter"`
				CountyFilter       interface{} `json:"countyFilter"`
				MunicipalityFilter interface{} `json:"municipalityFilter"`
				PostplaceFilter    interface{} `json:"postplaceFilter"`
			} `json:"searchStore"`
		} `json:"hydrationData"`
	} `json:"pageProps"`
	HydrationData struct {
		AppStore struct {
			Site              string   `json:"site"`
			SiteName          string   `json:"siteName"`
			Languages         []string `json:"languages"`
			DefaultLanguage   string   `json:"defaultLanguage"`
			Theme             string   `json:"theme"`
			Country           string   `json:"country"`
			DefaultDataSource string   `json:"defaultDataSource"`
			BaseURL           string   `json:"baseUrl"`
			DefaultSEO        struct {
				Noindex  bool `json:"noindex"`
				Nofollow bool `json:"nofollow"`
			} `json:"defaultSEO"`
			RecaptchaSiteKey string `json:"recaptchaSiteKey"`
			Lang             string `json:"lang"`
			SiteOwner        struct {
				Name    string `json:"name"`
				Address string `json:"address"`
				Zipcode string `json:"zipcode"`
				City    string `json:"city"`
				Phone   string `json:"phone"`
			} `json:"siteOwner"`
			ActiveCustomer interface{} `json:"activeCustomer"`
		} `json:"appStore"`
	} `json:"hydrationData"`
	I18N struct {
		ServerInstance  interface{} `json:"serverInstance"`
		InitialLanguage string      `json:"initialLanguage"`
		InitialStore    interface{} `json:"initialStore"`
	} `json:"i18n"`
	NSsp bool `json:"__N_SSP"`
}

type CorporateStructure struct {
	OrgNumbers        []string      `json:"orgNumbers"`
	CorporateAccounts []interface{} `json:"corporateAccounts"`
	Tree              struct {
		Name               string        `json:"name"`
		OrganisationNumber string        `json:"organisationNumber"`
		OwnedPercentage    interface{}   `json:"ownedPercentage"`
		SourceDate         interface{}   `json:"sourceDate"`
		InactiveDate       interface{}   `json:"inactiveDate"`
		Source             interface{}   `json:"source"`
		Employees          interface{}   `json:"employees"`
		CountryCode        string        `json:"countryCode"`
		CompanyAccounts    []interface{} `json:"companyAccounts"`
		Sub                []struct {
			Name               string      `json:"name"`
			OrganisationNumber string      `json:"organisationNumber"`
			OwnedPercentage    float64     `json:"ownedPercentage"`
			SourceDate         interface{} `json:"sourceDate"`
			InactiveDate       interface{} `json:"inactiveDate"`
			Source             interface{} `json:"source"`
			Employees          string      `json:"employees"`
			CountryCode        string      `json:"countryCode"`
			CompanyAccounts    []struct {
				Year         string      `json:"year"`
				Period       string      `json:"period"`
				PeriodStart  string      `json:"periodStart"`
				PeriodEnd    string      `json:"periodEnd"`
				Length       string      `json:"length"`
				Combined     bool        `json:"combined"`
				Currency     string      `json:"currency"`
				ReferenceURL interface{} `json:"referenceUrl"`
				Accounts     []struct {
					Code   string `json:"code"`
					Amount string `json:"amount"`
				} `json:"accounts"`
				AccIncompleteCode interface{} `json:"accIncompleteCode"`
				AccIncompleteDesc interface{} `json:"accIncompleteDesc"`
			} `json:"companyAccounts"`
			Sub []struct {
				Name               string      `json:"name"`
				OrganisationNumber string      `json:"organisationNumber"`
				OwnedPercentage    float64     `json:"ownedPercentage"`
				SourceDate         interface{} `json:"sourceDate"`
				InactiveDate       interface{} `json:"inactiveDate"`
				Source             interface{} `json:"source"`
				Employees          interface{} `json:"employees"`
				CountryCode        string      `json:"countryCode"`
				CompanyAccounts    []struct {
					Year         string      `json:"year"`
					Period       string      `json:"period"`
					PeriodStart  string      `json:"periodStart"`
					PeriodEnd    string      `json:"periodEnd"`
					Length       string      `json:"length"`
					Combined     bool        `json:"combined"`
					Currency     string      `json:"currency"`
					ReferenceURL interface{} `json:"referenceUrl"`
					Accounts     []struct {
						Code   string `json:"code"`
						Amount string `json:"amount"`
					} `json:"accounts"`
					AccIncompleteCode interface{} `json:"accIncompleteCode"`
					AccIncompleteDesc interface{} `json:"accIncompleteDesc"`
				} `json:"companyAccounts"`
				Sub []struct {
					Name               string      `json:"name"`
					OrganisationNumber string      `json:"organisationNumber"`
					OwnedPercentage    float64     `json:"ownedPercentage"`
					SourceDate         string      `json:"sourceDate"`
					InactiveDate       interface{} `json:"inactiveDate"`
					Source             string      `json:"source"`
					Employees          interface{} `json:"employees"`
					CountryCode        string      `json:"countryCode"`
					CompanyAccounts    []struct {
						Year         string      `json:"year"`
						Period       string      `json:"period"`
						PeriodStart  string      `json:"periodStart"`
						PeriodEnd    string      `json:"periodEnd"`
						Length       string      `json:"length"`
						Combined     bool        `json:"combined"`
						Currency     string      `json:"currency"`
						ReferenceURL interface{} `json:"referenceUrl"`
						Accounts     []struct {
							Code   string `json:"code"`
							Amount string `json:"amount"`
						} `json:"accounts"`
						AccIncompleteCode interface{} `json:"accIncompleteCode"`
						AccIncompleteDesc interface{} `json:"accIncompleteDesc"`
					} `json:"companyAccounts"`
					Sub []interface{} `json:"sub"`
					Seq int           `json:"seq"`
				} `json:"sub"`
				Seq int `json:"seq"`
			} `json:"sub"`
			Seq int `json:"seq"`
		} `json:"sub"`
		Seq int `json:"seq"`
	} `json:"tree"`
}
