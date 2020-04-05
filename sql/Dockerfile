FROM mysql:5.7

COPY *.sql /docker-entrypoint-initdb.d/
COPY utf8mb4.cnf binlog.cnf /etc/mysql/conf.d/

CMD ["--sql-mode="]
