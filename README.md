## JASS DB - Just A Simple Storage 
This is a strict learners project where the code will be clean, concise and self explanatory to learn how to build a simple database(not decided what kind)
but for learners process, once this project is a success we will move onto giving this database the power of being distributed
Lets ensure as an open source community i get support and motivation to complete this project.

oh one more thing ill be doing it in my fav programming language GOLANG.

## some of the primary resources
1. Database internals book
2. Designing data intensive applications
3. https://build-your-own.org/database/
4. https://medium.com/felixklauke/database-i-developing-your-own-data-storage-engine-aka-create-your-own-database-ed4560c8d80a
5. https://www.codeproject.com/Articles/1107279/Writing-a-MySQL-Storage-Engine-from-Scratch
6. https://nakabonne.dev/posts/write-tsdb-from-scratch/
7. https://github.com/cockroachdb/cockroach/blob/master/pkg/storage/fs/fs.go

Some awesome DB projects to take inspiration from:
1. https://github.com/pingcap/tidb
2. https://github.com/cockroachdb/cockroach
3. https://www.yugabyte.com/
4. https://github.com/vitessio/vitess
5. https://github.com/boltdb/bolt
6. https://github.com/syndtr/goleveldb

##Notes Rough work
1. We can always build DB by writing into a file and accessing the file intelligently but the problem with the approach is that 
Reading files the conventional way is not so good as concurrent reads sometimes wont read the complete file, concurrent writes will file, open and write will actually delete the older file and create a new file. so the better way to do this is (see file.go in rough work)