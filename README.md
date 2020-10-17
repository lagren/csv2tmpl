# Simple CSV to template

Use `csv2tmpl` to take a CSV file and convert to different output using templates. This makes it easy to convert CSV files to e.g. SQL scripts.

    csv2tmpl --input cmd/testdata/employees.csv --header-row --template "INSERT INTO employees VALUES ('{{ .employee_num }}', '{{ .name }}', '{{ .department }}');"
    INSERT INTO employees VALUES ('001', 'Bill', 'Engineering');
    INSERT INTO employees VALUES ('002', 'Jane', 'Management');

### Installing

Using csv2tmpl is easy. First, use `go get` to install the latest version:

    go get -u github.com/lagren/csv2tmpl/csv2tmpl
