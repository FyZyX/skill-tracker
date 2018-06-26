# Skill Tracker

## Overview
Skill Tracker is an application that helps users track their mastery of various skills. Skill mastery will dimish over time according to a model of memory degredation. The goal is to recommend skills for users to review to keep thier mastery scores up.

This project uses Go on the backend and HTML/CSS/JavaScript on the frontend.

## Run the Development Server
From the root of the repository, execute in bash:
```bash
$ go run main.go
```

In a brosweser, visit `localhost:8080`.

## Models
### `Skill`
Skills represnt abilities that users possess with varying levels of proficiency. As time without review goes by, skill proficiency diminishes. Proficiency is increased by reviewing the skill. The rate at which proficiency decreases depends on how much and how frequencly a user reviews the skill.

Attribute | Type | Description
--- | --- | ---
`Name` | string | The skill's unique identifier
`mastery` | float32 | the mastery is value between 0 and 1 representing the users proficiency in a skill

## Views and Controllers
Route | Controller | Template
--- | --- | ---
[`/`](#index-route) | `index.go` | `index.gohtml`

### Index Route
This is the application's homepage. The next step is to create user accounts so the user's can display on the homepage if they are logged in.
