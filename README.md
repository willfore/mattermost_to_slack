# Mattermost to Slack

This is a command line utility to take a bulk.json export from Mattermost and give you multiple import options into slack via the API. Some things to take note of:

* I was in a huge rush when I wrote this so it may break for you
* No there are not tests, yes I was in that big of a hurry
* You will need to have knowledge of mattermost command line as well as creating a slack app to run this
* Users will need to be imported via CSV

# Usage

`mm2slack -h`

This will give you help instructions of all commands available

`mm2slack set_auth --auth-token <token>`

Run this first and provide your bearer token.

`mm2slack get_users --export-file <path_to_file>`

Run this to grab all of your users and create a CSV file to be imported into slack
