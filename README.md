![image](https://user-images.githubusercontent.com/16916722/183267270-89b5b40b-52ad-447a-bd39-5057ce9020c9.png)
Dinder
=======================

[![Dependency Status](https://david-dm.org/sahat/hackathon-starter/status.svg?style=flat)](https://david-dm.org/sahat/hackathon-starter) [![devDependencies Status](https://david-dm.org/sahat/hackathon-starter/dev-status.svg)](https://david-dm.org/sahat/hackathon-starter?type=dev) [![Build Status](https://travis-ci.org/sahat/hackathon-starter.svg?branch=master)](https://travis-ci.org/sahat/hackathon-starter) [![Join the chat at https://gitter.im/sahat/hackathon-starter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/sahat/hackathon-starter?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

**Live Demo**: https://hackathon-starter.walcony.com

Jump to [What's new?](https://github.com/sahat/hackathon-starter/blob/master/CHANGELOG.md)

A boilerplate for **Node.js** web applications.

If you have attended any hackathons in the past, then you know how much time it takes to
get a project started: decide on what to build, pick a programming language, pick a web framework,
pick a CSS framework. A while later, you might have an initial project up on GitHub, and only then
can other team members start contributing. Or how about doing something as simple as *Sign in with Facebook*
authentication? You can spend hours on it if you are not familiar with how OAuth 2.0 works.

When I started this project, my primary focus was on **simplicity** and **ease of use**.
I also tried to make it as **generic** and **reusable** as possible to cover most use cases of hackathon web apps,
without being too specific. In the worst case, you can use this as a learning guide for your projects,
if for example you are only interested in **Sign in with Google** authentication and nothing else.


## IBM Bluemix Cloud Platform

**NOTE** *At this point it appears that Bluemix's free tier to host NodeJS apps is limited to 30 days.  If you are looking for a free tier service to host your app, Heroku might be a better choice at this point*

1. Create a Bluemix Account

    [Sign up](https://console.ng.bluemix.net/registration/?target=%2Fdashboard%2Fapps) for Bluemix, or use an existing account.

1. Download and install the [Cloud Foundry CLI](https://github.com/cloudfoundry/cli) to push your applications to Bluemix.

1. Create a `manifest.yml` file in the root of your application.
  ```
  applications:
  - name:      <your-app-name>
    host:      <your-app-host>
    memory:    128M
    services:
    - myMongo-db-name
  ```

  The host you use will determinate your application URL initially, e.g. `<host>.mybluemix.net`.
  The service name 'myMongo-db-name' is a declaration of your MongoDB service.  If you are using other services like Watson for example, then you would declare them the same way.

1. Connect and login to Bluemix via the Cloud-foundry CLI
  ```
  $ cf login -a https://api.ng.bluemix.net
  ```

1. Create a [MongoDB service](https://www.ng.bluemix.net/docs/#services/MongoDB/index.html#MongoDB)
  ```
  $ cf create-service mongodb 100 [your-service-name]
  ```
  **Note:** this is a free and experiment verion of MongoDB instance.
  Use the MongoDB by Compose instance for production applications:
  ```
  $ cf create-service compose-for-mongodb Standard [your-service-name]'
  ```


1. Push the application

    ```
    $ cf push
    ```
    ```
    $ cf env <your-app-name >
    (To view the *environment variables* created for your application)

    ```

**Done**, now go to the staging domain (`<host>.mybluemix.net`) and see your app running.

[Cloud Foundry Commands](https://console.ng.bluemix.net/docs/cli/reference/bluemix_cli/index.html)
[More Bluemix samples](https://ibm-bluemix.github.io/)
[Simple ToDo app in a programming language of your choice](https://github.com/IBM-Bluemix/todo-apps)


### IBM Watson
Be sure to check out the full list of Watson services to forwarder enhance your application functionality with a little effort. Watson services are easy to get going; it is simply a RESTful API call. Here is an example of a [Watson Toner Analyzer](https://tone-analyzer-demo.mybluemix.net/) to understand the emotional context of a piece of text that you send to Watson.


#### Watson catalog of services

**<img src="https://wbi.mybluemix.net/icons/conversation.svg?version=2" width="25"> [Conversation](https://www.ibm.com/watson/services/conversation/)** -     Quickly build and deploy chatbots and virtual agents across a variety of channels, including mobile devices, messaging platforms, and even robots.

**<img src="https://wbi.mybluemix.net/icons/discovery.svg" width="25"> [Discovery](https://www.ibm.com/watson/services/discovery/)** - Unlock hidden value in data to find answers, monitor trends and surface patterns with the world’s most advanced cloud-native insight engine.

**<img src="https://wbi.mybluemix.net/icons/language-translator.svg?version=4" width="20" width="25"> [Language Translator](https://www.ibm.com/watson/services/language-translator/)** - Translate text from one language to another.

**<img src="https://wbi.mybluemix.net/icons/natural-language-classifier.svg?version=2" width="25"> [Natural Language Classifier](https://www.ibm.com/watson/services/natural-language-classifier/)** - Interpret and classify natural language with confidence.

**<img src="https://wbi.mybluemix.net/icons/natural-language-understanding.svg?version=2" width="25"> [Natural Language Understanding](https://www.ibm.com/watson/services/natural-language-understanding/)** - Analyze text to extract meta-data from content such as concepts, entities, keywords and more.

**<img src="https://wbi.mybluemix.net/icons/personality-insights.svg?version=2" width="25"> [Personality Insights](https://www.ibm.com/watson/services/personality-insights/)** - Predict personality characteristics, needs and values through written text.

**<img src="https://wbi.mybluemix.net/icons/speech-to-text.svg?version=2" width="25"> [Speech to Text](https://www.ibm.com/watson/services/speech-to-text/)** - Convert audio and voice into written text for quick understanding of content.

**<img src="https://wbi.mybluemix.net/icons/text-to-speech.svg?version=2" width="25"> [Text to Speech](https://www.ibm.com/watson/services/text-to-speech/)** - Convert written text into natural sounding audio in a variety of languages and voices.

**<img src="https://wbi.mybluemix.net/icons/tone-analyzer.svg?version=2" width="25"> [Tone Analyzer](https://www.ibm.com/watson/services/tone-analyzer/)** - Understand emotions, social tendencies and perceived writing style.

**<img src="https://1.cms.s81c.com/sites/default/files/2019-02-21/1_VisualRecognitionHero.png" width="25"> [Visual Recognition](https://www.ibm.com/watson/services/visual-recognition/)** - Tag, classify and search visual content using machine learning.



[Click here](https://www.ibm.com/watson/developercloud/services-catalog.html) for live demos of each Watson service.


---

###  Google Cloud Platform
<img src="https://avatars2.githubusercontent.com/u/2810941?v=3&s=64" width="64">

- [Download and install Node.js](https://nodejs.org/)
- [Select or create](https://console.cloud.google.com/project) a Google Cloud Platform Console project
- [Enable billing](https://support.google.com/cloud/answer/6293499#enable-billing) for your project (there's a $300 free trial)
- Install and initialize the [Google Cloud SDK](https://cloud.google.com/sdk/docs/quickstarts)
- Create an `app.yaml` file at the root of your `hackathon-starter` folder with the following contents:

    ```yaml
    runtime: nodejs
    env: flex
    manual_scaling:
      instances: 1
    ```
- Make sure you've set `MONGODB_URI` in `.env.example`
- Run the following command to deploy the `hackathon-starter` app:

    ```bash
    gcloud app deploy
    ```
- [Monitor your deployed app](https://console.cloud.google.com/appengine) in the Cloud Console
- [View the logs](https://console.cloud.google.com/logs/viewer) for your app in the Cloud Console

Production
---------
 If you are starting with this boilerplate to build an application for prod deployment, or if after your hackathon you would like to get your project hardened for production use, see [prod-checklist.md](https://github.com/sahat/hackathon-starter/blob/master/prod-checklist.md).

Changelog
---------

You can find the changelog for the project in: [CHANGELOG.md](https://github.com/sahat/hackathon-starter/blob/master/CHANGELOG.md)


Contributing
------------

If something is unclear, confusing, or needs to be refactored, please let me know.
Pull requests are always welcome, but due to the opinionated nature of this
project, I cannot accept every pull request. Please open an issue before
submitting a pull request. This project uses
[Airbnb JavaScript Style Guide](https://github.com/airbnb/javascript) with a
few minor exceptions. If you are submitting a pull request that involves
Pug templates, please make sure you are using *spaces*, not tabs.

License
-------

The MIT License (MIT)

Copyright (c) 2014-2022 Sahat Yalkabov

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.


