# Mattermost to Slack

This is Open Source software and
This is a command line utility to take a bulk.json export from Mattermost and give you multiple import options into slack with API assistance. Some things to take note of:

* I was in a huge rush when I wrote this so it may break for you
* No there are not tests, yes I was in that big of a hurry
* You will need to have knowledge of mattermost command line as well as creating a slack app to run this
* Currently the slack import option is breaking. Slack support was "unable" to help
* Cannot import DM's into slack at this time
# Setup

Set an environment variable called SLACK_API_TOKEN

`export SLACK_API_TOKEN="<my_token>"`

The first thing to do is to get a bulk export from your mattermost instance. You will need command line access to do execute the binary. Instructions can be found here:
https://docs.mattermost.com/manage/bulk-export-tool.html

You will also need to create a slack app and add proper permissions to be able to list onversations, users, and chat. For info on creating a slack app to do this go here:
https://api.slack.com/start/designing

# Usage

## *Make sure to do the following steps in order to create the proper directory/file structure to import into slack*

`mm2slack -h`

This will give you help instructions of all commands available

`mm2slack make_user_csv --export-file <path_to_file>`

Run this to grab all of your users and create a CSV file to be imported into slack. You can use this to create inactive users while you work on importing the rest of your data.
Once you have this file go ahead and import it. You will need to have your users created to properly finish the rest of the process.

`mm2slack get_users --export-file <path_to_file> --slack-team-id <slack_team_id>`

This will create the proper users.json file to be included in the import

`mm2slack get_channels --export-file <path_to_file> --team-name <team_name> --slack-team-id <slack_team_id>`

Since the export uses all team names. Please specify your team name when importing channels from mattermost as well as the team-id in slack. This will create both Private and Public Channels in slack. It will also match up the exsisting users membership to the channels exported

## At this point you will need to export the posts as a CSV file and import them that way. The rest of this process is not currently working and Slack cannot answer why.

`mm2slack csv_posts --export-file <path_to_file> --slack-team-id <slack_team_id> --team-name <team_name>`

This will use the same process as below but will create csv files instead that you can import.

`mm2slack get_posts --export-file <path_to_file> --team-name <team_name> --slack-team-id <slack_team_id>`

This is the meat of the program. This pices goes through all the posts from mattermost and creates a folder with the channel name. Once this is done it creates a posts.json file and inserts all of the posts for that channel into the posts.json file

Once everything above is complete you will need to compress the folder into a zip file. You can also run `mm2slack cleanup` to delete all files created by the program

# Notes

DMs are not currently support while the code is in place to read the DMs from Mattermost. I didn't have time to request a full export from my env to figure out the proper structure/json layout. I will get into that in the future.

This project is far from perfect. As I said I needed to get this migration done in only a couple of days and so I went with the bare minimum to complete it.


# Contributing

Feel free to submit a pull request to improve this. As of writing there are literally no other options out there I could find.

# Notices
This is Open Source software and as such comes with no support, warranties or guarentees. Use at your own risk.
