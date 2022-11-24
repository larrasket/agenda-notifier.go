package notifiy

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"time"

	"github.com/0xAX/notificator"
	"github.com/salehmu/notifier.go/internal/messages"
	"github.com/salehmu/notifier.go/pkg/reader"
	. "github.com/salehmu/notifier.go/pkg/reader"
)

func ListenAndServe(config *Config) {
	logger := NewLogger()
	notify := notificator.New(notificator.Options{
		DefaultIcon: IconLoc,
		AppName:     "Emacs",
	})
	cmd := exec.Command(config.DoomScriptLoc, ExportScriptLoc)
	e := Entity{}
	if !config.Doom {
		cmd = exec.Command("emacs", "-script", fmt.Sprintf(ExportScriptLoc, config.EmacsLoc))
	}
	scanInt := time.NewTicker(time.Duration(config.ScanInt) * time.Second)
	defer scanInt.Stop()
	q := make(chan bool)
	defer close(q)
	go func() {
		err := Notify(e, q, notify, config.BeforeNotification)
		if err != nil {
			logger.Error(err)
		}
	}()
	for ; true; <-scanInt.C {
		data, err := ExtractData(*cmd)
		if err != nil {
			if config.Doom {
				logger.Fatal(fmt.Sprintf(messages.DoomscriptErr, err.Error()))
			} else {
				logger.Fatal(fmt.Sprintf("Couldn't extract agenda from emacs: %s", err.Error()))
			}
		}
		ne, err := reader.ComingEntity(data)
		if err != nil && !errors.Is(err, NoEntityErr) {
			logger.Fatal(fmt.Sprintf("Something wrong happend in reading the upcmming entity : %s", err.Error()))
		} else if errors.Is(err, NoEntityErr) {
			logger.Info("No upcoming entity")
			continue
		}
		fmt.Println(ne)
		if *ne != e {
			q <- true
			e = *ne
			go func() {
				err := Notify(e, q, notify, config.BeforeNotification)
				if err != nil {
					logger.Error(err)
				}
			}()
		}

	}

}

func Notify(e Entity, q <-chan bool, notify *notificator.Notificator, timeBefore int) error {
	now, err := time.Parse(TimeFormat, time.Now().Format(reader.TimeFormat))
	t := e.Time.Sub(now)
	coming := time.After(t)

	//calc time before
	re := time.Duration(timeBefore) * time.Minute * -1
	tb := e.Time.Add(re)
	tbt := tb.Sub(now)
	comingb := time.After(tbt)

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
		case <-comingb:
			msg := fmt.Sprintf("After %d mintues: %s", timeBefore, e.Name)
			err = notify.Push(e.Type, msg, IconLoc, notificator.UR_NORMAL)
		}
	}
}
func ExtractData(cmd exec.Cmd) ([]byte, error) {
	data, err := cmd.Output()
	start := bytes.Index(data, []byte(AgendaStart))
	end := bytes.Index(data, []byte(AgendaEnd))
	if err != nil && (start == -1 || end == -1) {
		return nil, err
	}
	if end == -1 {
		return nil, errors.New("couldn't reach end of csv file")
	}
	if start == -1 {
		return nil, errors.New("couldn't reach end of csv file")
	}
	data = (data)[start+len(AgendaStart) : end]
	return data, nil
}
