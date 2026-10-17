package kuruma

// Placeholders in Art that animate per frame:
//
//	@  wheel hub, spinning through - \ | /
//	!  warning light, blinking through Lights
//	%  crawler track tread, shifting each frame
//
// Vehicles face left, since they drive from the right edge to the left.
var All = []Vehicle{
	{
		Name:   "firetruck",
		Label:  "消防車",
		Color:  red,
		Lights: []string{red, white},
		Art: []string{
			`      _!____!_________________________________`,
			`  ___/  |  |=|=|=|=|=|=|=|=|=|=|=|=|=|=|=|=|=|`,
			` / [] | |  |_________________________________|`,
			`/_____|_|  |  ()  FIRE DEPT  119  ()  [====] |`,
			`| ==      ||_________________________________|`,
			`'---(@)-------------------------(@)---(@)----'`,
		},
	},
	{
		Name:   "police",
		Label:  "パトカー",
		Color:  white,
		Lights: []string{red, blue},
		Art: []string{
			`            __!==!__`,
			`       ____/  |  |  \_________`,
			`  ____/ [__]  |  | [__]  [__] \__`,
			` | ==   POLICE    110    POLICE  |`,
			` '---(@)------------------(@)----'`,
		},
	},
	{
		Name:   "ambulance",
		Label:  "救急車",
		Color:  white,
		Lights: []string{red, white},
		Art: []string{
			`        _!____________________!__`,
			`   ____/ | |                     |`,
			`  / [] | | |   +    AMBULANCE    |`,
			` /_____|_| |  +++      119       |`,
			`| ==       |   +                 |`,
			`'---(@)----'---------------(@)---'`,
		},
	},
	{
		Name:   "garbage",
		Label:  "ゴミ収集車",
		Color:  green,
		Lights: []string{yellow, white},
		Art: []string{
			`       _!______________________`,
			`   ___/  |  |                 |\`,
			`  / []|  |  |    GARBAGE      | \`,
			` /____|__|  |    RECYCLE (*)  |  |`,
			`| ==     |__|_________________|__|`,
			`'---(@)----------------(@)-(@)---'`,
		},
	},
	{
		Name:   "dump",
		Label:  "ダンプカー",
		Color:  yellow,
		Lights: []string{yellow, white},
		Art: []string{
			`        _!__   ____________________`,
			`   ____/ |  | |\ .:.:.:.:.:.:.:.: /`,
			`  / [] | |  | | \________________/`,
			` /_____|_|  |_|___________________\`,
			`| ==       |_______________________|`,
			`'---(@)----'-------------(@)-(@)---'`,
		},
	},
	{
		Name:   "mixer",
		Label:  "ミキサー車",
		Color:  cyan,
		Lights: []string{yellow, white},
		Art: []string{
			`                    _____________`,
			`       _!__     _.-' \  \  \  \  '-._`,
			`   ___/ |  |   (  \  \  \  \  \  \   )=`,
			`  / []| |  |    '-._\__\__\__\__\_.-'`,
			` /____|_|  |______/____________\_______`,
			`| ==       |___________________________|`,
			`'---(@)----'-----------------(@)-(@)---'`,
		},
	},
	{
		Name:   "excavator",
		Label:  "ショベルカー",
		Color:  yellow,
		Lights: []string{yellow, white},
		Art: []string{
			`      ______`,
			`     /      \___`,
			`    /   __      \___         __!___`,
			`   | __/  \__       \______ |  []  |___`,
			`   |/        \__           \|______|___|`,
			`  /_\           \___________|__________|_`,
			`  \_/                  .-----------------.`,
			`                      (%%%%%%%%%%%%%%%%%%%)`,
			`                       '-----------------'`,
		},
	},
	{
		Name:   "bulldozer",
		Label:  "ブルドーザー",
		Color:  yellow,
		Lights: []string{yellow, white},
		Art: []string{
			`                  ___!__`,
			`   ___           | [] | |______`,
			`  |   |          |____|_|______|_`,
			`  |   |=========|  ==           |`,
			`  |   |    .----------------------.`,
			`  |___|   (%%%%%%%%%%%%%%%%%%%%%%%%)`,
			`           '----------------------'`,
		},
	},
}

// Lookup returns the vehicle with the given Name.
func Lookup(name string) (Vehicle, bool) {
	for _, v := range All {
		if v.Name == name {
			return v, true
		}
	}
	return Vehicle{}, false
}
