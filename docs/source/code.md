# MailboxValidator Go API

```{py:function} SetAPIKey(api_key)
Configure MailboxValidator API key.

:param str api_key: (Required) MailboxValidator API key.
```

```{py:function} ValidateEmail(email)
Validate whether an email address is a valid email or not.

:param str email: (Required) The email address.

:return: Returns the validation result in JSON object.
:rtype: Object

**Successful Response Parameters**
| Field Name | Description |
|-----------|------------|
| EmailAddress | The input email address. |
| BaseEmailAddress | The input email address after sanitizing the username of the dots (only Gmail) and [subaddressing](https://en.wikipedia.org/wiki/Email_address#Sub-addressing). |
| Domain | The domain of the email address. |
| IsFree | Whether the email address is from a free email provider like Gmail or Hotmail. Return values: true, false, null  (null means not applicable) |
| IsSyntax | Whether the email address is syntactically correct. Return values: true, false |
| IsDomain | Whether the email address has a valid MX record in its DNS entries. Return values: true, false, null  (null means not applicable) |
| IsSMTP | Whether the mail servers specified in the MX records are responding to connections. Return values: true, false, null  (null means not applicable) |
| IsVerified | Whether the mail server confirms that the email address actually exist. Return values: true, false, null  (null means not applicable) |
| IsServerDown | Whether the mail server is currently down or unresponsive. Return values: true, false, null  (null means not applicable) |
| IsGreylisted | Whether the mail server employs greylisting where an email has to be sent a second time at a later time. Return values: true, false, null  (null means not applicable) |
| IsDisposable | Whether the email address is a temporary one from a disposable email provider. Return values: true, false, null  (null means not applicable) |
| IsSuppressed | Whether the email address is in our blacklist. Return values: true, false, null  (null means not applicable) |
| IsRole | Whether the email address is a role-based email address like admin@example.net or webmaster@example.net. Return values: true, false, null  (null means not applicable) |
| IsHighRisk | Whether the email address contains high risk keywords. Return values: true, false, null  (null means not applicable) |
| IsCatchAll | Whether the email address is a catch-all address. Return values: true, false, null  (null means not applicable) |
| IsDMARCEnforced | Whether the email domain is enforcing DMARC. Return values: true, false |
| IsStrictSPF | Whether the email domain is using strict SPF. Return values: true, false |
| WebsiteExist | Whether the email domain is a reachable website. Return values: true, false |
| MailboxValidatorScore | Email address reputation score. Score > 0.70 means good; score > 0.40 means fair; score <= 0.40 means poor. |
| TimeTaken | The time taken to get the results in seconds. |
| Status | Whether our system think the email address is valid based on all the previous fields. Return values: True, False |
| CreditsAvailable | The number of credits left to perform validations. |

**Error Response Parameters**
| Field Name | Description |
|-----------|------------|
| ErrorCode | The error code if there is any error. See error table in the [Error Codes](reference.md) section. |
| ErrorMessage | The error message if there is any error. See error table in the [Error Codes](reference.md) section. |

```

```{py:function} DisposableEmail(email)
Validate whether an email address is a disposable email or not.

:param str email: (Required) The email address.

:return: Returns the validation result in JSON object.
:rtype: Object

**Successful Response Parameters**
| Field Name | Description |
|-----------|------------|
| EmailAddress | The input email address. |
| IsDisposable | Whether the email address is a temporary one from a disposable email provider. Return values: True, False |
| CreditsAvailable | The number of credits left to perform validations. |


**Error Response Parameters**
| Field Name | Description |
|-----------|------------|
| ErrorCode | The error code if there is any error. See error table in the [Error Codes](reference.md) section. |
| ErrorMessage | The error message if there is any error. See error table in the [Error Codes](reference.md) section. |

```

```{py:function} FreeEmail(email)
Validate whether an email address is a free email or not.

:param str email: (Required) The email address.

:return: Returns the validation result in JSON object.
:rtype: Object

**Successful Response Parameters**
| Field Name | Description |
|-----------|------------|
| EmailAddress | The input email address. |
| IsFree | Whether the email address is from a free email provider like Gmail or Hotmail. Return values: True, False |
| CreditsAvailable | The number of credits left to perform validations. |


**Error Response Parameters**
| Field Name | Description |
|-----------|------------|
| ErrorCode | The error code if there is any error. See error table in the [Error Codes](reference.md) section. |
| ErrorMessage | The error message if there is any error. See error table in the [Error Codes](reference.md) section. |

```

