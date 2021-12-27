# dakoku
dakoku is a CLI application that integrates easily with [`Toggl`](https://toggl.com/), [`Asana`](https://asana.com/ja).

You can fetch tasks from Asana or fixed tasks from JSON, and just press the timer to record the time of the task by toggl.

Tie tasks and projects permanently and this makes time recording easier.

## Usage
### ◇ Deploy locally
```bash
$ dakoku init 'URL'
```
### ◇ How to use
```bash
$ dakoku tasks
Current
    ouk: [alfa]
    dje: [bravo]
    fug: [charlie]
Fixed
    mtg: []
    sup: [サポート]
```
```bash
$ dakoku beg ouk
Let’s begin [alfa].
```
```bash
$ dakoku fin ouk
Finished [alfa].
```
```bash
$ dakoku [today | yest | 2021/12/26]
alfa    -> 3h
bravo   -> 3h
charlie -> 2h
```
```bash
$ dakoku week
alfa    ->  8h
bravo   -> 12h
charlie -> 13h
delta   ->  7h
```
```bash
$ dakoku mos
alfa    ->  8h
bravo   -> 12h
charlie -> 13h
delta   ->  7h
```
```bash
$ dakoku --help
Time Tracking:
    $ dakoku show
    $ dakoku beg [task id]
    $ dakoku fin [task id]

Task Review:
    $ dakoku [today | yest | 2021/12/26]
    $ dakoku week
    $ dakoku mos
```

## Dependency
* go: v1.17
* [spf13 / cobra](https://github.com/spf13/cobra): v.1.3.0
go install github.com/spf13/cobra@v.1.3.0
