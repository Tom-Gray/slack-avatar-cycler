## Slack avatar cycler

Update your slack avatar according to a schedule. 

## Purpose 
Annoy your work colleagues by constantly cycling your avatar according to time of day.

I thought this might be interesting to write. 


## Usage
It is designed to run to a schedule, such as via a cron job that runs every hour.

It will evaluate the time of day and update your avatar to the configured image file.

First, define your schedule in this format

```
schedule:
  - morning:
      imageFile: "./Homestar/cranky.png"
      startTime: "8:00"
      endTime: "10:00"
  - lunch:
      imageFile: "./Homestar/closed_mouth.png"
      startTime: "10:00"
      endTime: "14:30"
  - default:
      imageFile: "./Homestar/sleepy.png"
```
Each schedule period ends at the defined start and end time. This feels a bit awkward.

Times must be represented in the 24hour format in the above example.

## Auth

The Slack user token needs to be provided by setting the `SLACK_TOKEN` env var.

Slack user tokens are prefixed with `xoxb`

## Images

Images need to be stored on the filesystem in a location accessible to the app

## Todo

add better logging



