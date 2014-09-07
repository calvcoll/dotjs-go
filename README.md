#DotJS Go
A reimplementation of the .js server, written in Go.
I have also unpacked the extension and remade it so that it can fall back to HTTP if needed.
However it should work with the default extension (add link later)

I consider this to be feature complete, but, you'll have to generate your own cert and key.

##Building
'''bash
go get bitbucket.org/kardianos/osext
'''

'''bash
go build
'''

##Cert and Key pair for SSL/TLS

**Note: ** The keys will need to be put in the (executable's path)/djsd_certs/ folder.

**For example:** If djsd's executable is run in */home/calv/bin*, the keys will be put in */home/calv/bin/djsd_certs/*

###Creating a cert and key pair
'''bash
openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout key.pem -out cert.pem
'''

This will give you a cert (valid for 10 years) and key pair that you can use (I guess Unix only[?]).
