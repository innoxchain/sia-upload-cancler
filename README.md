# Sia-upload-cancler Tool
This may be used to cancel ongoing uploads on a Sia node. Therefore some requirements need to be met:
- Sia is installed to /var/sia (may be changed in code)
- The tool must be run with the same user running the Sia daemon (siad), so the API password is available from ~/.sia/apipassword

## Run the tool
Get and run the tool by:
```
go get github.com/innoxchain/sia-upload-cancler
$GOPATH/bin/sia-upload-cancler
```