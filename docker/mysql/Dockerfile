FROM mysql/mysql-server
RUN chown -R mysql:root /var/lib/mysql/

ARG MYSQL_DATABASE
ARG MYSQL_USER
ARG MYSQL_PASSWORD
ARG MYSQL_ROOT_PASSWORD

# ADD dumps/sugar.sql /etc/mysql/data.sql
# RUN mysql -uroot -p"$MYSQL_ROOT_PASSWORD" < /etc/mysql/data.sql
# RUN sed -i 's/MYSQL_DATABASE/'$MYSQL_DATABASE'/g' /etc/mysql/data.sql
# RUN cp /etc/mysql/data.sql /docker-entrypoint-initdb.d