curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r '(.[] | select(.id==170) | .name), (.[] | select(.id==170) | .powerstats.power), (.[] | select(.id==170) | .appearance.gender)'
