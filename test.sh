#!/bin/sh

psql postgres -c "CREATE USER kopie WITH PASSWORD 'kopietestpw';"
psql postgres -c "ALTER USER kopie WITH SUPERUSER;"


chmod 0600 PGPASSFILE
export PGPASSFILE=./PGPASSFILE

createdb kopie_test --owner kopie
createdb kopie_test2 --owner kopie

psql kopie_test -c "create table test(name varchar, age numeric);"
psql kopie_test -c "create table test2(price numeric);"

psql kopie_test -c "alter table test owner to kopie;"
psql kopie_test -c "alter table test2 owner to kopie;"


go test -v

#dropdb kopie_test
#dropdb kopie_test2
#dropuser kopie