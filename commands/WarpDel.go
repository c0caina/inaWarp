package commands

import (
	"github.com/c0caina/inaWarp/global"
	"github.com/df-mc/dragonfly/server/cmd"
)

type WarpDel struct {
	Del  del
	Name string
}

func (wd WarpDel) Run(source cmd.Source, output *cmd.Output) {
	err := global.WarpSqlite.Delete(wd.Name)
	if err != nil {
		output.Errorf("[inaWarp] %v", err)
		return
	}
	output.Printf("[inaWarp] Warp %s deleted.", wd.Name)
}

type del string

func (s del) SubName() string {
	return "del"
}
