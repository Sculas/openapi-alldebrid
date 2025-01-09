openapi-generator-cli generate -i .\alldebrid.yaml -g go -o . --additional-properties=packageName=client --git-user-id sculas --git-repo-id openapi-alldebrid
rm .\client\ -r -force -ErrorAction Ignore
New-Item -ItemType Directory -Path .\client -Force | Out-Null
Move-Item -Path .\*.go -Destination .\client -Force