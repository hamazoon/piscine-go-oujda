#! /bin/bash 

curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r '.[] | select(.id == 70) | .powerstats .name | .appearance .gender '
