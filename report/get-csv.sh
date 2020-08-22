#!/bin/bash
pushd matsue/
bash ./build-csv.sh
popd
curl -L -sS "https://stopcovid19.metro.tokyo.lg.jp/data/130001_tokyo_covid19_patients.csv" | gnkf nl -n lf -o ./tokyo/130001_tokyo_covid19_patients.csv || exit 1
curl -L -sS "https://covid19.who.int/WHO-COVID-19-global-data.csv" -o WHO-COVID-19-global-data.csv || exit 1
cat WHO-COVID-19-global-data.csv | jvgrep Japan > WHO-COVID-19-japan-data.csv
