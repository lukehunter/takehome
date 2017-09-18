TakeHome Nmap Scanner
=====================
This is demo software that scans hosts using nmap and stores the results in a database.

Warnings
--------
This tool uses nmap, and running nmap on public hosts may get you in trouble with your ISP, the host in question, or both. Recommend only using this tool for local network scans. See https://nmap.org/book/legal-issues.html

Setup
-----
It is assumed that you already have nmap and MySQL installed on the server where this app will be running. Create a new database and user in MySQL by running createdb.sh then run the following using the database name you used when running createdb.sh:

cat initdb.sql | mysql -uroot -p -hhost mydbname

Set the TAKEHOME_DB_CONN variable with the database connection string using the user and password you used when running createdb.sh, e.g. if running MySQL on the same machine as the app:

export TAKEHOME_DB_CONN="newuser:newpassword@/mydbname"

URLs
----
* / (Index)
  * Provides html form interface for scanning a hostname or IP
* /scan POST
  * Result page for web form, ip or hostname in POST parameters
* /scan/{target} GET
  * Web service interface, does not perform a scan, only returns previous scan results in json format

Limitations
-----------
This tool is very much a demo project, and the following features would be needed in a production implementation.

* DOS protection (currently no rate limiting on scans)
* Handling of long running scans -- currently they will simply timeout
* Pagination support (currently the entire scan history is listed on one page)
* Support for ipv6 addresses
* Unit tests (probably requires some refactoring)
* Logging
* Rest api
* Client side formatting in js
* Configurable port range
