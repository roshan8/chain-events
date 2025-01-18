## Change-events

### Motivation
- One place to receive system events from all the places including kubernetes, different cloud provider, and webhook to receive notifications from other users application.
- A minimalistic dashboard to view all the events, support custom filters. 
- A rule engine for users to define alerting rules, and get notified via Slack and other notification channels. 
- Become an important too in SRE arsenal when they are start triaging the incident.

### To get clarity on following
- Are there any open-source tool exist for the above?
- Why can't we use loki, it supports accepting logs from various places. - May be we can't monitor logs based on some condition later on?
- Can sentry be used? - It has integration with most of them & their project seems to be open-source as well.


### How to develop
```sh
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
oapi-codegen -config oapi-codegen-config.yaml openapi.yaml
```