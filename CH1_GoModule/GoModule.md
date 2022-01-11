## go.mod
1. Define a namespace for the project
2. Specify direct dependencies used in the module

Example :
<pre>
go.mod
module namespace
go version
require direct_dependencies latest_version

go.sum
direct_dependencies hash_code
indirect_dependencies hash_code
</pre>

 ## go.sum 
 File to **ensure that future downloads of these modules retrieve the same bits as the first download**, to ensure the modules your project depends on do not change unexpectedly, whether for malicious, accidental, or other reasons

Both go.mod and go.sum should be checked into version control.

## go get module@version_tag

version_tag; latest, v1.0.0, etc (if version tag not specify will using default - latest)

