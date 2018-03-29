# :no_entry_sign: This plugin cannot run because it doesn't support [Fastly Billing API v2](https://docs.fastly.com/api/account#billing).

# mackerel-plugin-fastly-billing

[![License](https://img.shields.io/github/license/masutaka/mackerel-plugin-fastly-billing.svg?maxAge=2592000)][license]
[![GoDoc](https://godoc.org/github.com/masutaka/mackerel-plugin-fastly-billing?status.svg)][godoc]

[license]: https://github.com/masutaka/mackerel-plugin-fastly-billing/blob/master/LICENSE.txt
[godoc]: https://godoc.org/github.com/masutaka/mackerel-plugin-fastly-billing

## Description

[Fastly](https://www.fastly.com/) billing custom metrics plugin for mackerel.io agent.

## Synopsis

    mackerel-plugin-fastly-billing -api-key=<Fastly API Key>

## Example of mackerel-agent.conf

    [plugin.metrics.fastly_billing]
    command = "/path/to/mackerel-plugin-fastly-billing -api-key=<Fastly API Key>"

See also [Finding and managing your account info - Account management and security | Fastly Help Guides](https://docs.fastly.com/guides/account-management-and-security/finding-and-managing-your-account-info#finding-and-regenerating-your-api-key)
