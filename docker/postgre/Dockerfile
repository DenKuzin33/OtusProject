FROM postgres
COPY 01.dbschema_init.sql /docker-entrypoint-initdb.d/
COPY 02.dbdata_init.sql /docker-entrypoint-initdb.d/

ENV POSTGRES_DB banners_rotator
ENV POSTGRES_USER user
ENV POSTGRES_PASSWORD user