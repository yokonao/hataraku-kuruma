# hataraku-kuruma

A random working vehicle (はたらくくるま) drives across your terminal, like [sl](https://github.com/mtoyoda/sl).

```
      _*____*_________________________________
  ___/  |  |=|=|=|=|=|=|=|=|=|=|=|=|=|=|=|=|=|
 / [] | |  |_________________________________|
/_____|_|  |  ()  FIRE DEPT  119  ()  [====] |
| ==      ||_________________________________|
'---(-)-------------------------(-)---(-)----'
```

## Install

Download a prebuilt binary from [GitHub Releases](https://github.com/yokonao/hataraku-kuruma/releases).
Every release ships with a build provenance attestation.

See [docs/install.md](docs/install.md) for other install methods and how to verify a release.

## Usage

```
hataraku-kuruma [--kind NAME] [--loop] [--delay DURATION] [--no-color]
hataraku-kuruma --list
```

| flag           | default | description                              |
| -------------- | ------- | ---------------------------------------- |
| `-k`, `--kind` | random  | vehicle to drive (see `--list`)          |
| `-l`, `--list` |         | list the vehicles                        |
| `--loop`       |         | keep driving vehicles until Ctrl-C       |
| `--delay`      | `40ms`  | time per frame                           |
| `--no-color`   |         | disable colors; `NO_COLOR` also disables |

Vehicles: `firetruck` 消防車, `police` パトカー, `ambulance` 救急車, `garbage` ゴミ収集車, `dump` ダンプカー, `mixer` ミキサー車, `excavator` ショベルカー, `bulldozer` ブルドーザー.

Unlike sl, Ctrl-C stops the vehicle.
