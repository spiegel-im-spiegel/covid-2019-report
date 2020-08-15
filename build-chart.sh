#!/bin/bash
pushd report/matsue/
sh ./build-csv.sh
popd
curl -L -sS "https://stopcovid19.metro.tokyo.lg.jp/data/130001_tokyo_covid19_patients.csv" | gnkf nl -n lf -o ./report/tokyo/130001_tokyo_covid19_patients.csv || exit 1
libreoffice7.0 --convert-to csv --outdir report report/2019-ncov-cases.ods || exit 1
# cat report/2019-ncov-cases.csv | go run main.go  || exit 1
go run main.go report/2019-ncov-cases.csv report/tokyo/130001_tokyo_covid19_patients.csv
# rm report/2019-ncov-cases.csv
