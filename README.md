METRON
---

## Notice
**Not intended for production!**
This project is kind of a playground. I'm testing and learning a variety of architecture and design pattern, 
languages and technologies. Because of this, im not necessarily trying to reach the defined goal, and rather 
im trying to learn as much as I can

**EDIT 17.04.2022**: TikTok seems to be extending and improving their bot protection at the moment. Because of that I will put this project on-hold, let's see if I will pick it up again someday.

## Goal
Creating a third-party analytics platform for tiktok.com.

### Challenges
- TikTok does not provide full api access, so the data in demand is parsed from the public web-app.
- TikTok has its entity related data spread between different pages. E.g Video:
  - Views only on account page visible
  - Other meta information only on video detail page
  - Comments only viewable as logged in user
- TikTok is actively developing and deploying measures aimed to make this as hard as possible. This Includes:
  - Detecting non-human users
  - Locking interaction behind captchas as soon as a bot is suspected
  - Failing requests/page loads if a bot is suspected

## Documentation
Detailed documentation can be found in the `/docs` folder.
