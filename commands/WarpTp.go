package commands

import (
	"github.com/c0caina/inaWarp/global"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type WarpTp struct {
	Tp   tp
	Name string
}

func (wt WarpTp) Run(source cmd.Source, output *cmd.Output) {
	XYZ, err := global.WarpSqlite.SelectName(wt.Name)
	if err != nil {
		output.Errorf("[inaWarp] %v", err)
		return
	}

	p, _ := source.(*player.Player)
	p.Teleport(XYZ)
	output.Printf("[inaWarp] You are transported warp %s.", wt.Name)
}

type tp string

func (s tp) SubName() string {
	return "tp"
}
