FROM library/postgres

ENV POSTGRES_USER gcuser
ENV POSTGRES_PASSWORD gcpass

COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432
