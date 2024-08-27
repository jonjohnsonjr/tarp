# `tarp`

For example, you might want to sum the size of files in a tar file without doing string surgery.

```
tarp < example.tar | jq .Size | paste -sd+ - | bc
77857283
```

Or you might want to view some information that `tar -tv` hides from you, like PAX records:

```
< go-1.22-1.22.2-r1.apk gunzip | tarp | jq .PAXRecords | head
null
null
null
null
{
  "APK-TOOLS.checksum.SHA1": "773011eb6962b1bd6cfd190351484b9787656e17"
}
{
  "APK-TOOLS.checksum.SHA1": "99649e3820655d0ba3c612207c4d23c6e53a5874"
}
```
