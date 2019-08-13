package mailboxvalidator

import (
	// "fmt"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type ValidationResult struct {
	EmailAddress          string `json:"email_address"`
	Domain                string `json:"domain"`
	IsFree                string `json:"is_free"`
	IsSyntax              string `json:"is_syntax"`
	IsDomain              string `json:"is_domain"`
	IsSMTP                string `json:"is_smtp"`
	IsVerified            string `json:"is_verified"`
	IsServerDown          string `json:"is_server_down"`
	IsGreylisted          string `json:"is_greylisted"`
	IsDisposable          string `json:"is_disposable"`
	IsSuppressed          string `json:"is_suppressed"`
	IsRole                string `json:"is_role"`
	IsHighRisk            string `json:"is_high_risk"`
	IsCatchAll            string `json:"is_catchall"`
	MailboxValidatorScore string `json:"mailboxvalidator_score"`
	TimeTaken             string `json:"time_taken"`
	Status                string `json:"status"`
	CreditsAvailable      uint32 `json:"credits_available"`
	ErrorCode             string `json:"error_code"`
	ErrorMessage          string `json:"error_message"`
}

type ValidationResultDisposable struct {
	EmailAddress     string `json:"email_address"`
	IsDisposable     string `json:"is_disposable"`
	CreditsAvailable uint32 `json:"credits_available"`
	ErrorCode        string `json:"error_code"`
	ErrorMessage     string `json:"error_message"`
}

type ValidationResultFree struct {
	EmailAddress     string `json:"email_address"`
	IsFree           string `json:"is_free"`
	CreditsAvailable uint32 `json:"credits_available"`
	ErrorCode        string `json:"error_code"`
	ErrorMessage     string `json:"error_message"`
}

var apiKey string

func APIKey() string {
	return apiKey
}

func SetAPIKey(key string) {
	apiKey = key
}

func ValidateEmail(email string) (*ValidationResult, error) {
	url := "https://api.mailboxvalidator.com/v1/validation/single?email=" + url.QueryEscape(email) + "&key=" + url.QueryEscape(apiKey)

	result := &ValidationResult{}
	err := getJson(url, result)

	if err != nil {
		return result, err
	}

	// PrintResult(result)

	return result, nil
}

func DisposableEmail(email string) (*ValidationResultDisposable, error) {
	url := "https://api.mailboxvalidator.com/v1/email/disposable?email=" + url.QueryEscape(email) + "&key=" + url.QueryEscape(apiKey)

	result := &ValidationResultDisposable{}
	err := getJson(url, result)

	if err != nil {
		return result, err
	}

	// PrintResult(result)

	return result, nil
}

func FreeEmail(email string) (*ValidationResultFree, error) {
	url := "https://api.mailboxvalidator.com/v1/email/free?email=" + url.QueryEscape(email) + "&key=" + url.QueryEscape(apiKey)

	result := &ValidationResultFree{}
	err := getJson(url, result)

	if err != nil {
		return result, err
	}

	// PrintResult(result)

	return result, nil
}

func getJson(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// for debugging purposes
// func PrintResult(x *ValidationResult) {
// fmt.Printf("EmailAddress: %s\n", x.EmailAddress)
// fmt.Printf("Domain: %s\n", x.Domain)
// fmt.Printf("IsFree: %s\n", x.IsFree)
// fmt.Printf("IsSyntax: %s\n", x.IsSyntax)
// fmt.Printf("IsDomain: %s\n", x.IsDomain)
// fmt.Printf("IsSMTP: %s\n", x.IsSMTP)
// fmt.Printf("IsVerified: %s\n", x.IsVerified)
// fmt.Printf("IsServerDown: %s\n", x.IsServerDown)
// fmt.Printf("IsGreylisted: %s\n", x.IsGreylisted)
// fmt.Printf("IsDisposable: %s\n", x.IsDisposable)
// fmt.Printf("IsSuppressed: %s\n", x.IsSuppressed)
// fmt.Printf("IsRole: %s\n", x.IsRole)
// fmt.Printf("IsHighRisk: %s\n", x.IsHighRisk)
// fmt.Printf("IsCatchAll: %s\n", x.IsCatchAll)
// fmt.Printf("MailboxValidatorScore: %s\n", x.MailboxValidatorScore)
// fmt.Printf("TimeTaken: %s\n", x.TimeTaken)
// fmt.Printf("Status: %s\n", x.Status)
// fmt.Printf("CreditsAvailable: %d\n", x.CreditsAvailable)
// fmt.Printf("ErrorCode: %s\n", x.ErrorCode)
// fmt.Printf("ErrorMessage: %s\n", x.ErrorMessage)
// }
