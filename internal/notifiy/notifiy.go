package notifiy

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/0xAX/notificator"
	"github.com/salehmu/notifier.go/internal/messages"
	"github.com/salehmu/notifier.go/pkg/reader"
	. "github.com/salehmu/notifier.go/pkg/reader"
	"os/exec"
	"time"
)

func ListenAndServe(config *Config) {
	logger := NewLogger()
	notify := notificator.New(notificator.Options{
		DefaultIcon: IconLoc,
		AppName:     "Emacs",
	})
	cmd := exec.Command(config.DoomScript, ExportScriptLoc)
	e := Entity{}
	if !config.Doom {
		cmd = exec.Command("emacs", "-batch", "-l", config.EmacsLoc, "-eval", reader.EmacsExporter)
	}
	scanInt := time.NewTicker(time.Duration(config.ScanInt) * time.Second)
	defer scanInt.Stop()
	q := make(chan bool)
	defer close(q)
	go func() {
		err := Notify(e, q, notify)
		if err != nil {
			logger.Error(err)
		}
	}()
	for ; true; <-scanInt.C {
		fmt.Println(e)
		data, err := ExtractData(*cmd)
		if err != nil {
			if config.Doom {
				logger.Fatal(fmt.Sprintf(messages.DoomscriptErr, err))
			} else {
				logger.Fatal(fmt.Sprintf("Couldn't extract agenda from emacs: %s", err))
			}
		}
		ne, err := reader.ComingEntity(data)
		if err != nil && !errors.Is(err, NoEntityErr) {
			logger.Fatal(fmt.Sprintf("Something wrong happend in reading the upcmming entity : %s", err))
		} else if errors.Is(err, NoEntityErr) {

			logger.Info("No upcoming entity")
			continue
		}
		if *ne != e {
			q <- true
			e = *ne
			go func() {
				err := Notify(e, q, notify)
				if err != nil {
					logger.Error(err)
				}
			}()
		}

	}

}

func Notify(e Entity, q <-chan bool, notify *notificator.Notificator) error {
	now, err := time.Parse(TimeFormat, time.Now().Format(reader.TimeFormat))
	t := e.Time.Sub(now)
	coming := time.After(t)
	sent := false
	for {
		select {
		case <-q:
			now, _ = time.Parse(TimeFormat, time.Now().Format(reader.TimeFormat))
			if !sent && e.Time.Sub(now).Minutes() <= 1 {
				err = notify.Push(e.Type, e.Name, IconLoc, notificator.UR_NORMAL)
				if err != nil {
					return err
				}
			}
			return nil
		case <-coming:
			err = notify.Push(e.Type, e.Name, IconLoc, notificator.UR_NORMAL)
			sent = true
		}
	}
}
func ExtractData(cmd exec.Cmd) ([]byte, error) {
	data, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	start := bytes.Index(data, []byte(AgendaStart))
	end := bytes.Index(data, []byte(AgendaEnd))
	data = (data)[start+len(AgendaStart) : end]
	return data, nil
}
