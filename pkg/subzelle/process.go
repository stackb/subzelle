package subzelle

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
)

func startPlugin(dir, executable string, args []string, envVars []string) (*exec.Cmd, error) {
	ListFiles(dir)

	env := os.Environ()
	env = append(env, envVars...)

	cmd := exec.Command(executable, args...)
	cmd.Dir = dir
	cmd.Env = env
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("could not start subprocess: %v", err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		s := <-c
		if runtime.GOOS != "windows" {
			cmd.Process.Signal(s)
		} else {
			cmd.Process.Kill()
		}
	}()

	// err = cmd.Wait()
	// if err != nil {
	// 	if exitError, ok := err.(*exec.ExitError); ok {
	// 		waitStatus := exitError.Sys().(syscall.WaitStatus)
	// 		// log.Printf("bazel exited with abnormal status: %v", waitStatus.ExitStatus())
	// 		return err, waitStatus.ExitStatus()
	// 	}
	// 	// return fmt.Errorf("could not launch Bazel: %v", err), 1
	// }

	// log.Print("bazelisk exited seemingly normally")
	return cmd, nil
}

// ListFiles - convenience debugging function to log the files under a given dir
func ListFiles(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("%v\n", err)
			return err
		}
		if info.Mode()&os.ModeSymlink > 0 {
			link, err := os.Readlink(path)
			if err != nil {
				return err
			}
			log.Printf("%s -> %s", path, link)
			return nil
		}

		log.Println(path)
		return nil
	})
}
