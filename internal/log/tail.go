/*
#######
##                                     _
##        ____    _____ ____ _        (_)__ ________ ____  ___
##       (_-< |/|/ / _ `/ _ `/ _     / / _ `/ __/ _ `/ _ \/ _ \
##      /___/__,__/\_,_/\_, / (_) __/ /\_,_/_/  \_, /\___/_//_/
##                     /___/     |___/         /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package log

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"
)

type logFile struct {
	name    string
	file    *os.File
	reLevel *regexp.Regexp
	sigEnd  chan os.Signal
}

func (lf *logFile) openFile() (bool, error) {
	fmt.Printf("--> %s\n", lf.name)

	warning := true

	for {
		file, err := os.Open(lf.name)
		if err == nil {
			fmt.Println("...")

			lf.file = file

			return false, nil
		}

		if os.IsNotExist(err) {
			if warning {
				fmt.Println("--> this file doesn't exist...(wait or ^C ?)")

				warning = false
			}

			select {
			case <-time.After(100 * time.Millisecond):
			case <-lf.sigEnd:
				fmt.Println("END")
				return true, nil
			}
		} else {
			return true, err
		}
	}
}

func (lf *logFile) readFile() error {
	reader := bufio.NewReader(lf.file)

	for {
		line, err := reader.ReadString('\n')
		switch err {
		case nil:
			renderLine(lf.reLevel.FindString(line), line)
		case io.EOF:
			select {
			case <-time.After(10 * time.Millisecond):
			case <-lf.sigEnd:
				fmt.Println("END")
				return nil
			}
		default:
			return err
		}
	}
}

// TailFile AFAIRE
func TailFile(file string) error {
	lf := &logFile{
		name:    file,
		reLevel: regexp.MustCompile(`[{]\w{3}[}]`),
		sigEnd:  make(chan os.Signal, 1),
	}

	defer close(lf.sigEnd)
	signal.Notify(lf.sigEnd, syscall.SIGINT, syscall.SIGTERM)

	end, err := lf.openFile()
	if err != nil {
		return err
	}

	if end {
		return nil
	}

	defer lf.file.Close()

	if err := lf.readFile(); err != nil {
		return err
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
