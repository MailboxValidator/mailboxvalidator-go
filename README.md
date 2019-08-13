[![Go Report Card](https://goreportcard.com/badge/github.com/MailboxValidator/mailboxvalidator-go)](https://goreportcard.com/report/github.com/MailboxValidator/mailboxvalidator-go)

MailboxValidator Go Package
===========================

This Go Package provides an easy way to call the MailboxValidator API which validates if an email address is a valid one.

This module can be used in many types of projects such as:

 - validating a user's email during sign up
 - cleaning your mailing list prior to an email marketing campaign
 - a form of fraud check


Installation
============

```
go get github.com/MailboxValidator/mailboxvalidator-go
```


Dependencies
============

An API key is required for this module to function.

Go to https://www.mailboxvalidator.com/plans#api to sign up for FREE API plan and you'll be given an API key.


Usage for validating emails
===========================

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

Functions
=========

### SetAPIKey(api_key)

Sets the MailboxValidator API key.

### ValidateEmail(email_address)

Performs email validation on the supplied email address.

Result Fields
=============

### EmailAddress

The input email address.

### Domain

The domain of the email address.

### IsFree

Whether the email address is from a free email provider like Gmail or Hotmail.

Return values: True, False

### IsSyntax

Whether the email address is syntactically correct.

Return values: True, False

### IsDomain

Whether the email address has a valid MX record in its DNS entries.

Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable)

### IsSMTP

Whether the mail servers specified in the MX records are responding to connections.

Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable)

### IsVerified

Whether the mail server confirms that the email address actually exist.

Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable)

### IsServerDown

Whether the mail server is currently down or unresponsive.

Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable)

### IsGreylisted

Whether the mail server employs greylisting where an email has to be sent a second time at a later time.

Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable)

### IsDisposable

Whether the email address is a temporary one from a disposable email provider.

Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable)

### IsSuppressed

Whether the email address is in our blacklist.

Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable)

### IsRole

Whether the email address is a role-based email address like admin@example.net or webmaster@example.net.

Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable)

### IsHighRisk

Whether the email address contains high risk keywords.

Return values: True, False, -&nbsp;&nbsp;&nbsp;(- means not applicable)

### IsCatchAll

Whether the email address is a catch-all address.

Return values: True, False, Unknown, -&nbsp;&nbsp;&nbsp;(- means not applicable)

### MailboxValidatorScore

Email address reputation score.

Score > 0.70 means good; score > 0.40 means fair; score <= 0.40 means poor.

### TimeTaken

The time taken to get the results in seconds.

### Status

Whether our system think the email address is valid based on all the previous fields.

Return values: True, False

### CreditsAvailable

The number of credits left to perform validations.

### ErrorCode

The error code if there is any error. See error table below.

### ErrorMessage

The error message if there is any error. See error table below.


Usage for checking if an email is from a disposable email provider
==================================================================

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

Functions
=========

### SetAPIKey(api_key)

Sets the MailboxValidator API key.

### DisposableEmail(email_address)

Check if the supplied email address is from a disposable email provider.

Result Fields
=============

### EmailAddress

The input email address.

### IsDisposable

Whether the email address is a temporary one from a disposable email provider.

Return values: True, False

### CreditsAvailable

The number of credits left to perform validations.

### ErrorCode

The error code if there is any error. See error table below.

### ErrorMessage

The error message if there is any error. See error table below.


Usage for checking if an email is from a free email provider
============================================================

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

Functions
=========

### SetAPIKey(api_key)

Sets the MailboxValidator API key.

### FreeEmail(email_address)

Check if the supplied email address is from a free email provider.

Result Fields
=============

### EmailAddress

The input email address.

### IsFree

Whether the email address is from a free email provider like Gmail or Hotmail.

Return values: True, False

### CreditsAvailable

The number of credits left to perform validations.

### ErrorCode

The error code if there is any error. See error table below.

### ErrorMessage

The error message if there is any error. See error table below.


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

Copyright (C) 2019 by MailboxValidator.com, support@mailboxvalidator.com
