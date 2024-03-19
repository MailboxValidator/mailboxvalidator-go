# Quickstart

## Dependencies

An API key is required for this module to function.

Go to https://www.mailboxvalidator.com/plans#api to sign up for FREE API plan and you'll be given an API key.

## Installation

Run the following command to install the module:

```
go get github.com/mailboxvalidator/mailboxvalidator-go/v2
```

## Sample Codes

### Validate email

You can validate whether an email address is invalid or not as below:

```go
package main

import (
	"fmt"
	"github.com/mailboxvalidator/mailboxvalidator-go/v2"
)

func main() {
	mailboxvalidator.SetAPIKey("PASTE_YOUR_API_KEY_HERE")
	
	result, err := mailboxvalidator.ValidateEmail("example@example.com")
	
	if err != nil {
		panic(err.Error())
	}
	
	if result.ErrorMessage != "" {
		panic(result.ErrorMessage)
	}
	
	fmt.Printf("EmailAddress: %+v\n", result.EmailAddress)
	fmt.Printf("BaseEmailAddress: %+v\n", result.BaseEmailAddress)
	fmt.Printf("Domain: %+v\n", result.Domain)
	fmt.Printf("IsFree: %+v\n", result.IsFree)
	fmt.Printf("IsSyntax: %+v\n", result.IsSyntax)
	fmt.Printf("IsDomain: %+v\n", result.IsDomain)
	fmt.Printf("IsSMTP: %+v\n", result.IsSMTP)
	fmt.Printf("IsVerified: %+v\n", result.IsVerified)
	fmt.Printf("IsServerDown: %+v\n", result.IsServerDown)
	fmt.Printf("IsGreylisted: %+v\n", result.IsGreylisted)
	fmt.Printf("IsDisposable: %+v\n", result.IsDisposable)
	fmt.Printf("IsSuppressed: %+v\n", result.IsSuppressed)
	fmt.Printf("IsRole: %+v\n", result.IsRole)
	fmt.Printf("IsHighRisk: %+v\n", result.IsHighRisk)
	fmt.Printf("IsCatchAll: %+v\n", result.IsCatchAll)
	fmt.Printf("IsDMARCEnforced: %+v\n", result.IsDMARCEnforced)
	fmt.Printf("IsStrictSPF: %+v\n", result.IsStrictSPF)
	fmt.Printf("WebsiteExist: %+v\n", result.WebsiteExist)
	fmt.Printf("MailboxValidatorScore: %+v\n", result.MailboxValidatorScore)
	fmt.Printf("TimeTaken: %+v\n", result.TimeTaken)
	fmt.Printf("Status: %+v\n", result.Status)
	fmt.Printf("CreditsAvailable: %+v\n", result.CreditsAvailable)
	fmt.Printf("ErrorCode: %+v\n", result.ErrorCode)
	fmt.Printf("ErrorMessage: %+v\n", result.ErrorMessage)
}
```

### Check if an email is from a disposable email provider

You can validate whether an email address is disposable email address or not as below:

```go
package main

import (
	"fmt"
	"github.com/mailboxvalidator/mailboxvalidator-go/v2"
)

func main() {
	mailboxvalidator.SetAPIKey("PASTE_YOUR_API_KEY_HERE")
	
	result, err := mailboxvalidator.DisposableEmail("example@example.com")
	
	if err != nil {
		panic(err.Error())
	}
	
	if result.ErrorMessage != "" {
		panic(result.ErrorMessage)
	}
	
	fmt.Printf("EmailAddress: %+v\n", result.EmailAddress)
	fmt.Printf("IsDisposable: %+v\n", result.IsDisposable)
	fmt.Printf("CreditsAvailable: %+v\n", result.CreditsAvailable)
	fmt.Printf("ErrorCode: %+v\n", result.ErrorCode)
	fmt.Printf("ErrorMessage: %+v\n", result.ErrorMessage)
}
```

### Check if an email is from a free email provider

You can validate whether an email address is free email address or not as below:

```go
package main

import (
	"fmt"
	"github.com/mailboxvalidator/mailboxvalidator-go/v2"
)

func main() {
	mailboxvalidator.SetAPIKey("PASTE_YOUR_API_KEY_HERE")
	
	result, err := mailboxvalidator.FreeEmail("example@example.com")
	
	if err != nil {
		panic(err.Error())
	}
	
	if result.ErrorMessage != "" {
		panic(result.ErrorMessage)
	}
	
	fmt.Printf("EmailAddress: %+v\n", result.EmailAddress)
	fmt.Printf("IsFree: %+v\n", result.IsFree)
	fmt.Printf("CreditsAvailable: %+v\n", result.CreditsAvailable)
	fmt.Printf("ErrorCode: %+v\n", result.ErrorCode)
	fmt.Printf("ErrorMessage: %+v\n", result.ErrorMessage)
}
```
