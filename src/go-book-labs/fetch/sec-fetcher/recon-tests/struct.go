package main

type ApiResponse struct {
    Matches []Match   `json:"matches"`
    Facets  Facets    `json:"facets"`
    Total   int       `json:"total"`
}

type Match struct {
    Product   string    `json:"product"`
    Hash      int       `json:"hash"`
    IP        int       `json:"ip"`
    Org       string    `json:"org"`
    ISP       string    `json:"isp"`
    Transport string    `json:"transport"`
    Cpe       []string  `json:"cpe"`
    Data      string    `json:"data"`
    ASN       string    `json:"asn"`
    Port      int       `json:"port"`
    Hostnames []string  `json:"hostnames"`
    Location  Location  `json:"location"`
    Timestamp string    `json:"timestamp"`
    Domains   []string  `json:"domains"`
    Http      Http      `json:"http"`
    OS        *string   `json:"os"`  // Nullable
    Shodan    Shodan    `json:"_shodan"`
    IPStr     string    `json:"ip_str"`
}

type Location struct {
    City         string  `json:"city"`
    RegionCode   string  `json:"region_code"`
    Longitude    float64 `json:"longitude"`
    CountryCode  string  `json:"country_code"`
    Latitude     float64 `json:"latitude"`
    PostalCode   *string `json:"postal_code"`
    CountryName  string  `json:"country_name"`
}

type Http struct {
    Title      string  `json:"title"`
    Server     string  `json:"server"`
    Host       string  `json:"host"`
    Html       string  `json:"html"`
    Location   string  `json:"location"`
    Components map[string]interface{} `json:"components"`
}

type Shodan struct {
    Crawler string `json:"crawler"`
    PTR     bool   `json:"ptr"`
    ID      string `json:"id"`
    Module  string `json:"module"`
    Options map[string]interface{} `json:"options"`
}

type Facets struct {
    Country []Country `json:"country"`
}

type Country struct {
    Count int    `json:"count"`
    Value string `json:"value"`
}

