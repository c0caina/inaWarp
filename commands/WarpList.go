package commands

import (
	"github.com/c0caina/inaWarp/repository"
	"github.com/df-mc/dragonfly/server/cmd"
)

type WarpList struct {
	List list
}

func (wl WarpList) Run(source cmd.Source, output *cmd.Output) {
	warps, err := repository.WarpRepo.SelectAll()
	if err != nil {
		output.Errorf("[inaWarp] %v", err)
		return
	}

	output.Print("-------- Warps --------")
	for _, w := range warps {
		output.Printf("%v coordinates: %v %v %v", w.Name, int(w.XYZ[0]), int(w.XYZ[1]), int(w.XYZ[2]))
	}
}

type list string

func (s list) SubName() string {
	return "list"
}
