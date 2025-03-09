# `eclosix`

This repository contains the code for both the front and backend of `eclosix`. Documentation on how the launch the website is contained in the `back` and `front` folders.

## Initialize the data

To initialize the data, run the commands:

```
cd scripts && go run . parse_list && cd ..
cd back && go run main.go migrate && cd ..
cd scripts && go run . serve && cd ..
```