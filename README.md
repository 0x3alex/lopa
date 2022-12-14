# lopa
More of a collection of apis that can be interacted with via Discord than a Discord bot.

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

<img src="https://i.imgur.com/dChbMCC.png" width="350" height="350">

## Disclaimer 
You are welcomed to use the bot or parts of the bot (api implementations,...) for your personal projects! :)

## General

A basic discord bot, which is built in [discordgo](https://github.com/bwmarrin/discordgo), to interact with some APIs

### Why a discord bot?
The "discord bot" part is only an interface to interact with my implimentations of the APIs, since I am not especially good at front-end development :D

### Implemented APIs

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/2/26/Spotify_logo_with_text.svg/1200px-Spotify_logo_with_text.svg.png" width="340" height="100">
- The Spotify Web API

- The usage of the [spotify web api](https://developer.spotify.com/documentation/web-api/) in this bot is very specific, since only a small part of the api is needed and processed in an very specific manner
- Authentication is done with a `clientID` and an `clientSecrent`, which can be obtained [here](https://developer.spotify.com/dashboard/applications)

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/f/f0/Urban_Dictionary_logo.svg/1200px-Urban_Dictionary_logo.svg.png" width="340" height="100">
- The Urban Dictionary API

- The API is not docutmented by the Urban Dictionary, since its a "secret", that those parts of the API face the public
