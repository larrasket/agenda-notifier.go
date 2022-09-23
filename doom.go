package main

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
)

func InitConfig() (doom bool, err error) {
	doom = false
	var loc *string
	doom, loc = IsDoom()
	if !doom {
		fmt.Println(doomNotFound)
		_, loc = IsEmacs()
	}
	if loc != nil {
		fmt.Println(emacsFound + *loc)
	} else {
		fmt.Println(nonFound)
		err = ErrNotFound
		return
	}
	conf, _ := yaml.Marshal(MakeConfig(doom, *loc))
	err = ioutil.WriteFile(ConfigDirDef, conf, 0644)
	return
}
