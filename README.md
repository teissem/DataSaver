# Data Saver

The aim of this project is to designed a data saver.

## Configuration file format

Configuration file is in JSON format :

```json
{
    "destination" : "...",
    "log" : "...",
    "compression" : "...",
    "git" : {
        "username" : "...",
        "password" : "...",
        "repositories" : [
            { "source" : "...", "destination" : "..." },
            { "source" : "...", "destination" : "..." }
        ]
    },
    "folder" : {
        "path" : [
            { "source" : "...", "destination" : "..." },
            { "source" : "...", "destination" : "..." }
        ]
    }
}
```

## Architecture of the project

To see the architecture of the project, go to `doc/architecture.md`
