set -euo pipefail

DONTBELAST_DATABASE_USER=dontbelast
DONTBELAST_DATABASE_NAME=dontbelast

psql --username "${POSTGRES_USER}" <<SQL
CREATE USER ${DONTBELAST_DATABASE_USER} PASSWORD '${DONTBELAST_DATABASE_PASSWORD}';
CREATE DATABASE ${DONTBELAST_DATABASE_NAME} OWNER ${DONTBELAST_DATABASE_USER};
SQL