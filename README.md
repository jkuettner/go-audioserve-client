# Audioserve client for go

This is a go client library for [audioserve api](https://github.com/izderadicka/audioserve).
The specs are from the [openapi documentation](https://github.com/izderadicka/audioserve/blob/master/docs/api.md) and were generated with [oapi-codegen](https://github.com/deepmap/oapi-codegen).

Originally I had planned to use this client for an alternative Android client, but I don't know if this will ever happen due to time and motivation.
It also depends on how much time I have to maintain this project and if I can keep it up to date for future API changes.
So please bear with me.

# Usage

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "github.com/jkuettner/go-audioserve-client/pkg/api/v1"
    apiv1specs "github.com/jkuettner/go-audioserve-client/pkg/api/v1/specs"
    "time"
)

func main() {
    ctx := context.Background()

    // create the client instance
    client, err := apiv1.NewClient(&apiv1.ClientOpts{

        // public credentials of the official demo server from the Audioserve project
        ServerURL:      "https://audioserve.zderadicka.eu",
        SharedSecret:   "mypass",
        
        // timeout of the http handler
        RequestTimeout: 60 * time.Second,
    })
    if err != nil {
        panic(err)
    }

    [...]
    // see examples below
}
```

# Examples

## List the collections. Normally there is only one anyway

```go
collections, err := client.GetCollections(ctx)
if err != nil {
    panic(err)
}

out, err := json.MarshalIndent(collections, "", "    ")
if err != nil {
    panic(err)
}

fmt.Println(string(out))
```
Output:
```yaml
{
    "count": 1,
    "folder_download": true,
    "names": [
        "audiobooks"
    ],
    "shared_positions": true,
    "version": "0.17.0" 
} 
```

## Request folder informations

```go
folder, err := client.GetFolderPath(ctx, 0, apiv1specs.Path("Čapek Karel/R.U.R"), &apiv1specs.GetColIdFolderPathParams{})
if err != nil {
    panic(err)
}

out, err := json.MarshalIndent(folder, "", "    ")
if err != nil {
    panic(err)
}

fmt.Println(string(out))
```
Output:
```yaml
{
  "cover": {
    "mime": "image/jpeg",
    "path": "Čapek Karel/R.U.R/rur_rossums_universal_robots_1409.jpg"
  },
  "description": {
    "mime": "text/markdown",
    "path": "Čapek Karel/R.U.R/info.md"
  },
  "files": [
    {
      "meta": {
        "bitrate": 64,
        "duration": 2276,
        "tags": null
      },
      "mime": "audio/mpeg",
      "name": "rur_00_capek_64kb.mp3",
      "path": "Čapek Karel/R.U.R/rur_00_capek_64kb.mp3",
      "section": null
    },
    {
      "meta": {
        "bitrate": 64,
        "duration": 1928,
        "tags": null
      },
      "mime": "audio/mpeg",
      "name": "rur_01_capek_64kb.mp3",
      "path": "Čapek Karel/R.U.R/rur_01_capek_64kb.mp3",
      "section": null
    },
    {
      "meta": {
        "bitrate": 64,
        "duration": 1354,
        "tags": null
      },
      "mime": "audio/mpeg",
      "name": "rur_02_capek_64kb.mp3",
      "path": "Čapek Karel/R.U.R/rur_02_capek_64kb.mp3",
      "section": null
    },
    {
      "meta": {
        "bitrate": 64,
        "duration": 887,
        "tags": null
      },
      "mime": "audio/mpeg",
      "name": "rur_03_capek_64kb.mp3",
      "path": "Čapek Karel/R.U.R/rur_03_capek_64kb.mp3",
      "section": null
    }
  ],
  "is_file": false,
  "modified": 1638100000000,
  "subfolders": [],
  "tags": null,
  "total_time": 6445
}

```

## Download a track

```go
audioData, err := client.GetColIdAudioPath(ctx, 0, "Austene Jane/Pride And Prejudice/prideandprejudice_09_austen_64kb.mp3", &apiv1specs.GetColIdAudioPathParams{})
if err != nil {
    panic(err)
}

if err := os.WriteFile("prideandprejudice_09_austen_64kb.mp3", audioData.Data, 0644); err != nil {
    panic(err)
}
```

`audioData` contains the (possibly transcoded) data of the requested file.
In this example we save the track to a file.

# Generate api specs

To build all supported api specs, the following command in the root-folder of this project is sufficient:

```go
go generate ./...
```