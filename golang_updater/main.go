package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	args := os.Args

	//prepare the log file
	cwd, _ := os.Getwd()
	logfile, err := os.Create(filepath.Join(cwd, "updater.log"))
	if err != nil {
		panic(err)
	}
	defer logfile.Close()
	logger := log.New(logfile, "", log.LstdFlags)

	//make sure number of parameters are good to go
	if args != nil && len(args) >= 3 {
		updateAsar := args[1]
		appAsar := args[2]

		//sleep for 5 seconds
		fmt.Println("sleep for 5 seconds")
		time.Sleep(5 * time.Second)
		fmt.Printf("Will move %s to %s\n", updateAsar, appAsar)

		//check updateAsar
		_, updateAsarErr := os.Stat(updateAsar)
		if os.IsNotExist(updateAsarErr) {
			fmt.Fprintf(os.Stderr, "Update.asar not exist.\n")
			logger.Fatal("updateAsar not exist")
			return
		}

		//check appAsar
		_, err = pathExists(appAsar)
		if err != nil {
			logger.Fatal("appAsar not exist")
			logger.Fatal(err)
			return
		}

		// backup app.asar
		_, err = copyFile(appAsar, appAsar+".bak")
		if err != nil {
			logger.Fatal("backup app.asar failed")
			logger.Fatal(err)
			return
		}

		// delete app.asar
		err = os.Remove(appAsar)
		if err != nil {
			fmt.Printf("%s", err)
			time.Sleep(5 * time.Second)
			logger.Fatal("delete app.asar failed")
			logger.Fatal(err)
			copyFile(appAsar+".bak", appAsar)
			return
		}

		// copy update.asar to app.asar
		_, err = copyFile(updateAsar, appAsar)
		if err != nil {
			logger.Fatal("copy update.asar to app.asar failed")
			logger.Fatal(err)
			copyFile(appAsar+".bak", appAsar)
			return
		}

		//check executableFilePath and restart  StsStudentClient
		if len(args) == 4 {
			executabalFilePath := args[3]
			result := exec.Command(executabalFilePath)
			result.Run()
			fmt.Printf("New version started: %s \n", result.ProcessState)
		}

	} else {
		fmt.Fprintf(os.Stderr, "Wrong arguments: %s\n", args)
		logger.Fatal("at least 2 parameters are needed")
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func copyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
