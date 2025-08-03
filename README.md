# Fursona TUI

✨Show info about your fursona or OC right in your terminal✨
![](assets/Screenshot.png)

## Installation

Install using go

```sh
go get github.com/tekofx/fursona-tui
```

## Configuration

- In `$HOME/.config/fursona-tui` place a picture (.JPG or .PNG).
- Run the app `fursona-tui`. This will generate a new settings.json.

### Customize

The data shown can be customized in the JSON file `$HOME/.config/fursona-tui/settings.json`.

```json
{
  "name": "Name",
  "palette": ["#FFFFFF", "#000000"],
  "data": {
    "Gender": "Gender",
    "Pronouns": "Pronouns",
    "Species": "Species"
  }
}
```

The fields in data can be any pair of strings.

# Example

![](assets/Screenshot2.png)

## Config

```json
{
  "name": "Mr Meow",
  "palette": ["#FC9B05", "#FCF2B7", "#FCD36A", "#FCE28E", "#FC9C54"],
  "data": {
    "📍Location": "Europe",
    "Pronouns": "🐱🐈",
    "Favorite Food": "Tuna 🐟"
  }
}
```
