# Simple CSV to template

![Go](https://github.com/lagren/csv2tmpl/workflows/Go/badge.svg?branch=main)

Use `csv2tmpl` to take a CSV file and convert to different output using templates. This makes it easy to convert CSV files to e.g. SQL scripts.

    csv2tmpl --input cmd/testdata/employees.csv --header-row --template "INSERT INTO employees VALUES ('{{ .employee_num }}', '{{ .name }}', '{{ .department }}');"
    INSERT INTO employees VALUES ('001', 'Bill', 'Engineering');
    INSERT INTO employees VALUES ('002', 'Jane', 'Management');

### Installing

Using csv2tmpl is easy. First, use `go get` to install the latest version:

    go get -u github.com/lagren/csv2tmpl/csv2tmpl

### Functions

Use functions to convert values to another format. Supported functions:

- String manipulation using `lower`, `upper`, `kebab`, `snake`, `camel`, `prefix`, or `suffix`
- Calculate hash sum using `md5`, `sha`, `sha256`, or `sha512` (hex output)

Example template usage:
- `{{ upper .last_name }}` or `{{ .last_name | upper }}` outputs `SMITH`
- `{{ prefix "Mr " .last_name }}` or `{{ .last_name | prefix "Mr " }}` outputs `Mr Smith`
- `{{ suffix "-san" .last_name }}` or `{{ .last_name | suffix "-san" }}` outputs `Smith-san`