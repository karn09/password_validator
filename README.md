### Password Validator

**Prerequisites**: 
Have Go installed in order to run `go get`

Passwords **MUST**

1. Have an 8 character minimum
2. AT LEAST 64 character maximum
3. Allow all ASCII characters and spaces (unicode optional)
4. Not be a common password

To install into go/bin:

`go get -u github.com/karn09/password_validator`

Alternative, clone this repo, and from the root directory:

`go run password_validator.go weak_password_list.txt`

Example usage:
```
cat input_passwords.txt | ./password_validator weak_password_list.txt
mom -> Error: Too Short
password1 -> Error: Too Common
*** -> Error: Invalid Charaters
```

A useful weak / common password list can be downloaded here:

https://github.com/danielmiessler/SecLists/raw/master/Passwords/Common-Credentials/10-million-password-list-top-1000000.txt