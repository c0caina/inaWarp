package main

import (
	"embed"
	"fmt"
	"github.com/c0caina/inaWarp/commands"
	"github.com/c0caina/inaWarp/repository"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	"github.com/soypat/rebed"
	"io/ioutil"
	"os"
)

//go:embed assets
var assets embed.FS

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	config, err := readConfig()
	if err != nil {
		log.Fatalln(err)
	}

	srv := server.New(&config, log)
	srv.CloseOnProgramEnd()
	if err := srv.Start(); err != nil {
		log.Fatalln(err)
	}

	err = rebed.Patch(assets, "")
	if err != nil {
		log.Fatalf("[inaWarp] %v", err)
	}

	db, err := repository.NewSqliteDB("sqlite3", "assets/inaWarp/warp.sqlite3")
	if err != nil {
		log.Fatalf("[inaWarp] %v", err)
	}
	repository.WarpRepo = repository.NewWarpSqlite(db)

	cmd.Register(cmd.New("warp", "Warp system control.", []string{}, commands.WarpSet{}, commands.WarpDel{}, commands.WarpTp{}, commands.WarpList{}))

	for {
		if _, err := srv.Accept(); err != nil {
			return
		}
	}
}

// readConfig reads the configuration from the config.toml file, or creates the file if it does not yet exist.
func readConfig() (server.Config, error) {
	c := server.DefaultConfig()
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return c, fmt.Errorf("failed encoding default config: %v", err)
		}
		if err := ioutil.WriteFile("config.toml", data, 0644); err != nil {
			return c, fmt.Errorf("failed creating config: %v", err)
		}
		return c, nil
	}
	data, err := ioutil.ReadFile("config.toml")
	if err != nil {
		return c, fmt.Errorf("error reading config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return c, fmt.Errorf("error decoding config: %v", err)
	}
	return c, nil
}
