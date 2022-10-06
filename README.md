## Slack avatar cycler

Update your slack avatar according to a schedule. 

## Purpose 
I have no actual use for this but I thought it might be interesting. 


## Usage
It is designed to run to a schedule, such as via a cron job that runs every hour.

It will evaluate the time of day and update your avatar to the configured image file.

First, define your schedule in this format

```
schedule:
  - morning:
     imageFile: "./Homestar/cranky.png"
     startTime: "08:00AM"
     endTime: "10:00AM"
  - lunch:
     imageFile: "./Homestar/closed_mouth.png"
     startTime: "10:00AM"
     endTime: "2:30PM"
  - arvo:
     imageFile: "./Homestar/open_mouth.png"
     startTime: "2:30PM"
     endTime: "5:00PM"
  - default:
      imageFile: "./Homestar/sleepy.png"
```
Each schedule period ends at the defined start and end time. This feels a bit awkward. 
## Auth

The Slack user token needs to be provided by setting the `SLACK_TOKEN` env var.

## Images

Images need to be stored on the filesystem in a location accessible to the app

## Todo

add better logging

