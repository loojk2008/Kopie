#!/bin/sh

psql postgres -c "CREATE USER kopie WITH PASSWORD 'kopietestpw';"

chmod 0600 PGPASSFILE
export PGPASSFILE=./PGPASSFILE

createdb kopie_test --owner kopie

psql kopie_test -c "create table test(name varchar, age numeric);"
psql kopie_test -c "alter table test owner to kopie;"

psql kopie_test -c "create table test2(price numeric);"
psql kopie_test -c "alter table test2 owner to kopie;"


go test

dropdb kopie_test
dropdb kopie_test2

dropuser kopie