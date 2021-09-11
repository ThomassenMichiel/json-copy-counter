# JSON Copy Counter

***

## Description

***

This is a tiny application solely needed to easily parse JSON-files from a 3rd party tool used by the writer's company.
The main take is to expect nothing from this tool if you aren't associated with the writer's company.

## Usage

***

Provide `Kopies.json` as a file in the root directory of project and run `go build .\src\Main.go` and look at the
generated output.

### Requirements

The structure of the JSON-file to input is as follows:

```json
{
  "Kopies_View": [
    {
      "Kopie": "A4ZW",
      "Hoeveelheid": "10",
      "Organisatie": "Organisation"
    }
  ]
}
```

`Kopies_View` is the parent container.
`Kopie` is the type of the copy made. 
`Hoeveelheid` is the amount of copies made. It unfortunately is provided as a string.
`Organisatie` is which organisation has ordered the copies.
All other fields are ignored.