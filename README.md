# `tarp`

For example, you might want to sum the size of files in a tar file without doing string surgery.

```
tarp < example.tar | jq .Size | paste -sd+ - | bc
77857283
```
