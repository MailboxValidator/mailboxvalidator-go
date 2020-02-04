[![Go Report Card](https://goreportcard.com/badge/github.com/MailboxValidator/mailboxvalidator-go)](https://goreportcard.com/report/github.com/MailboxValidator/mailboxvalidator-go)

MailboxValidator Go Package
===========================

This Go Package enables user to easily validate if an email address is valid, a type of disposable email or free email.

This module can be useful in many types of projects, for example

 - to validate an user's email during sign up
 - to clean your mailing list prior to email sending
 - to perform fraud check
 - and so on


Installation
============

```
go get github.com/MailboxValidator/mailboxvalidator-go
```


Dependencies
============

An API key is required for this module to function.

Go to https://www.mailboxvalidator.com/plans#api to sign up for FREE API plan and you'll be given an API key.

Functions
=========

## SetAPIKey(api_key)

Sets the MailboxValidator API key.

## ValidateEmail(email_address)

Performs email validation on the supplied email address.

### Return Fields

| Field Name | Description |
|-----------|------------|
| email_address | The input email address. |
| domain | The domain of the email address. |
| is_free | Whether the email address is from a free email provider like Gmail or Hotmail. Return values: True, False |
| is_syntax | Whether the email address is syntactically correct. Return values: True, False |
| is_domain | Whether the email address has a valid MX record in its DNS entries. Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable) |
| is_smtp | Whether the mail servers specified in the MX records are responding to connections. Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable) |
| is_verified | Whether the mail server confirms that the email address actually exist. Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable) |
| is_server_down | Whether the mail server is currently down or unresponsive. Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable) |
| is_greylisted | Whether the mail server employs greylisting where an email has to be sent a second time at a later time. Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable) |
| is_disposable | Whether the email address is a temporary one from a disposable email provider. Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable) |
| is_suppressed | Whether the email address is in our blacklist. Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable) |
| is_role | Whether the email address is a role-based email address like admin@example.net or webmaster@example.net. Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable) |
| is_high_risk | Whether the email address contains high risk keywords. Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable) |
| is_catchall | Whether the email address is a catch-all address. Return values: True, False, Unknown, -&nbsp;&nbsp;&nbsp;(- means not applicable) |
| mailboxvalidator_score | Email address reputation score. Score > 0.70 means good; score > 0.40 means fair; score <= 0.40 means poor. |
| time_taken | The time taken to get the results in seconds. |
| status | Whether our system think the email address is valid based on all the previous fields. Return values: True, False |
| credits_available | The number of credits left to perform validations. |
| error_code | The error code if there is any error. See error table in the below section. |
| error_message | The error message if there is any error. See error table in the below section. |

## DisposableEmail(email_address)

Check if the supplied email address is from a disposable email provider.

### Return Fields

| Field Name | Description |
|-----------|------------|
| email_address | The input email address. |
| is_disposable | Whether the email address is a temporary one from a disposable email provider. Return values: True, False |
| credits_available | The number of credits left to perform validations. |
| error_code | The error code if there is any error. See error table in the below section. |
| error_message | The error message if there is any error. See error table in the below section. |

## FreeEmail(email_address)

Check if the supplied email address is from a free email provider.

### Return Fields

| Field Name | Description |
|-----------|------------|
| email_address | The input email address. |
| is_free | Whether the email address is from a free email provider like Gmail or Hotmail. Return values: True, False |
| credits_available | The number of credits left to perform validations. |
| error_code | The error code if there is any error. See error table in the below section. |
| error_message | The error message if there is any error. See error table below. |


Sample Codes
============

## Validate email

```go
package main

import (
	"fmt"
	"github.com/mailboxvalidator/mailboxvalidator-go"
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
	
	fmt.Printf("EmailAddress: %s\n", result.EmailAddress)
	fmt.Printf("Domain: %s\n", result.Domain)
	fmt.Printf("IsFree: %s\n", result.IsFree)
	fmt.Printf("IsSyntax: %s\n", result.IsSyntax)
	fmt.Printf("IsDomain: %s\n", result.IsDomain)
	fmt.Printf("IsSMTP: %s\n", result.IsSMTP)
	fmt.Printf("IsVerified: %s\n", result.IsVerified)
	fmt.Printf("IsServerDown: %s\n", result.IsServerDown)
	fmt.Printf("IsGreylisted: %s\n", result.IsGreylisted)
	fmt.Printf("IsDisposable: %s\n", result.IsDisposable)
	fmt.Printf("IsSuppressed: %s\n", result.IsSuppressed)
	fmt.Printf("IsRole: %s\n", result.IsRole)
	fmt.Printf("IsHighRisk: %s\n", result.IsHighRisk)
	fmt.Printf("IsCatchAll: %s\n", result.IsCatchAll)
	fmt.Printf("MailboxValidatorScore: %s\n", result.MailboxValidatorScore)
	fmt.Printf("TimeTaken: %s\n", result.TimeTaken)
	fmt.Printf("Status: %s\n", result.Status)
	fmt.Printf("CreditsAvailable: %d\n", result.CreditsAvailable)
	fmt.Printf("ErrorCode: %s\n", result.ErrorCode)
	fmt.Printf("ErrorMessage: %s\n", result.ErrorMessage)
}
```

## Check if an email is from a disposable email provider

```go
package main

import (
	"fmt"
	"github.com/mailboxvalidator/mailboxvalidator-go"
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
	
	fmt.Printf("EmailAddress: %s\n", result.EmailAddress)
	fmt.Printf("IsDisposable: %s\n", result.IsDisposable)
	fmt.Printf("CreditsAvailable: %d\n", result.CreditsAvailable)
	fmt.Printf("ErrorCode: %s\n", result.ErrorCode)
	fmt.Printf("ErrorMessage: %s\n", result.ErrorMessage)
}
```

## Check if an email is from a free email provider

```go
package main

import (
	"fmt"
	"github.com/mailboxvalidator/mailboxvalidator-go"
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
	
	fmt.Printf("EmailAddress: %s\n", result.EmailAddress)
	fmt.Printf("IsFree: %s\n", result.IsFree)
	fmt.Printf("CreditsAvailable: %d\n", result.CreditsAvailable)
	fmt.Printf("ErrorCode: %s\n", result.ErrorCode)
	fmt.Printf("ErrorMessage: %s\n", result.ErrorMessage)
}
```


Errors
======

| error_code | error_message |
| ---------- | ------------- |
| 100 | Missing parameter. |
| 101 | API key not found. |
| 102 | API key disabled. |
| 103 | API key expired. |
| 104 | Insufficient credits. |
| 105 | Unknown error. |

Copyright
=========

Copyright (C) 2019-2020 by MailboxValidator.com, support@mailboxvalidator.com
