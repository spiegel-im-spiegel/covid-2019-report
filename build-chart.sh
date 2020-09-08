#!/bin/bash
pushd report/
./get-csv.sh
popd
# go run main.go report/WHO-COVID-19-japan-data.csv report/tokyo/130001_tokyo_covid19_patients.csv
go run main.go
