https://f004.backblazeb2.com/file/odaspro-images/odoo-jammy_lxc_amd64.tar.xz
https://f004.backblazeb2.com/file/odaspro-images/odoo-noble_lxc_amd64.tar.xz

121235
quest15data
10.10.121.235/23
10.10.120.1

12Gb
6


ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIHJ0bdsdngPAX0+DLS1b0nhmaH3/vLkqLeukxq/vBDmQ ppreeper@THNK-PP


sudo /usr/local/bin/update
sudo apt purge -y postgresql-18 postgresql-client-18
sudo apt install -y postgresql-17 postgresql-17-pgvector




psql -c "CREATE EXTENSION IF NOT EXISTS pg_stat_statements;"
psql -c "CREATE ROLE odoo WITH CREATEDB NOSUPERUSER ENCRYPTED PASSWORD 'odooodoo' LOGIN;"
psql -c "CREATE DATABASE template_odoo TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'C' LC_CTYPE = 'C' IS_TEMPLATE = TRUE OWNER=odoo;"
sleep 2
psql -c "CREATE EXTENSION pg_stat_statements" template_odoo
psql -c "CREATE EXTENSION vector" template_odoo
psql -c "CREATE EXTENSION pg_trgm" template_odoo
psql -c "update pg_extension set extowner = (select usesysid from pg_user where usename='odoo');" template_odoo
