#!/bin/bash
libreoffice6.4 --convert-to csv --outdir report report/2019-ncov-cases.ods || exit 1
# cat report/2019-ncov-cases.csv | go run main.go  || exit 1
go run main.go report/2019-ncov-cases.csv report/tokyo/130001_tokyo_covid19_patients.csv
rm report/2019-ncov-cases.csv
