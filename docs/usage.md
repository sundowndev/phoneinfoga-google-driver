```shell
$ phoneinfoga

PhoneInfoga is one of the most advanced tools to scan phone numbers using only free resources.

Usage:
  phoneinfoga [command]

Available Commands:
  help        Help about any command
  recon       Launch custom format reconnaissance
  scan        Scan a phone number
  serve       Serve web client
  version     Print current version of the tool

Flags:
  -h, --help   help for phoneinfoga

Use "phoneinfoga [command] --help" for more information about a command.
```

#### Basic scan

```
phoneinfoga -n "(+42) 837544833"
```

Country code and special chars such as `( ) - +` will be escaped so typing US-based numbers stay easy : 

```
phoneinfoga -n "+1 555-444-3333"
```

!!! note "Note that the country code is essential. You don't know which country code to use ? [Find it here](https://www.countrycode.org/)"

#### Output file

Check several numbers at once and send results to a file.

```
phoneinfoga -i numbers.txt -o results.txt
```

Input file must contain one phone number per line. Invalid numbers will be skipped.

#### Footprinting

```
phoneinfoga -n +42837544833 -s footprints
```

#### Custom format reconnaissance

You don't know where to search and what custom format to use ? Let the tool try several custom formats based on the country code for you.

```
phoneinfoga -n +42837544833 -s any --recon
```

## Available scanners

Use `any` to disable this feature. Default value: `all`

- numverify
- ovh
- footprints

**Numverify** provide standard but useful informations such as number's country code, location, line type and carrier.

**OVH** is, besides being a web and cloud hosting company, a telecom provider with several VoIP numbers in the Europe. Thanks to their API-key free [REST API](https://api.ovh.com/), we are able to tell if a number is owned by OVH Telecom or not.

**Footprints** scanner uses Google search engine and [Google Dorks](https://en.wikipedia.org/wiki/Google_hacking) to search phone number's footprints everywhere on the web. It allows you to search for scam reports, social media profiles, documents and more.

## Examples

Check for a number range on OVH :

```
phoneinfoga -n "+33 01 88 33 40 32" -s ovh
```

Output : 

```
[!] ---- Fetching informations for 330188334032 ---- [!]
[*] Running local scan...
[+] International format: +33 1 88 33 40 32
[+] Local format: 188334032
[+] Country found: France (+33)
[+] City/Area: France
[+] Carrier: 
[+] Timezone: Europe/Paris
[i] The number is valid and possible.
[*] Running OVH scan...
[+] 1 result found in OVH database
[+] Number range: 018833xxxx
[+] City: Paris
[+] Zip code: 
Continue scanning ? (y/N) 
[i] Good bye!
```