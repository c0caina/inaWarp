package commands

import (
	"github.com/c0caina/inaWarp/repository"
	"github.com/c0caina/inaWarp/repository/models"
	"github.com/df-mc/dragonfly/server/cmd"
)

type WarpSet struct {
	Set  set
	Name string
}

func (ws WarpSet) Run(source cmd.Source, output *cmd.Output) {
	err := repository.WarpRepo.Insert(models.Warp{Name: ws.Name, XYZ: source.Position()})
	if err != nil {
		output.Errorf("[inaWarp] %v", err)
		return
	}

	output.Printf("[inaWarp] Warp %s created.", ws.Name)
}

type set string

func (s set) SubName() string {
	return "set"
}
