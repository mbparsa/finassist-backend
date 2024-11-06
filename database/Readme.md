
Start: brew services start postgresql
Stop: brew services stop postgresql

### configure postgres to only use localhost
listen_addresses = 'localhost'


this is how you can get the status of Postgres on my mac
pg_ctl -D /opt/homebrew/var/postgresql@14 status


## create a db 
createdb finassist 


###  Create a user
finassist=# \du
finassist=# CREATE USER postgres WITH PASSWORD '<My Usuals>';

### grant access
GRANT CONNECT ON DATABASE your_database TO new_username;
GRANT USAGE ON SCHEMA public TO new_username;
GRANT SELECT, INSERT ON ALL TABLES IN SCHEMA public TO new_username;

