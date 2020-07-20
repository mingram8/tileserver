# tileserver

Basic x/y/z server

Runs with directory and listen port command line arguments. Listen Port defaults to 8080.

For example:

./cmd --directory=/Users/mingram

There's 2 endpoints. 

`/heath` will return a 200 and "Ok" if the server is up and running.
`/tiles` takes a layer query parameter, a x param, y param, and z param. For example `http://localhost:8080/tiles?layer=test_dir&x=0&y=1&z=3.png`

It will return whatever png is there. If you need to use a pbf file or anything else, just change the extension.

This will open up your file system, so best to sandbox it with the directory command and not run it in your root file system. 
