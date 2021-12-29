package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"syscall"
)

const UID = 123

func main() {
	const bin = "/tmp/airupnp-static"
	const dir = "/perm/airupnp"

	if err := os.WriteFile(bin, airupnpPrebuilt, 0555); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(dir, 0755); err != nil && !errors.Is(err, os.ErrExist) {
		log.Fatal(err)
	}

	if err := syscall.Chown(dir, UID, UID); err != nil {
		log.Fatal(err)
	}

	if err := syscall.Chmod(dir, 0755); err != nil {
		log.Fatal(err)
	}

	if err := syscall.Chdir(dir); err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(bin)
	cmd.Env = append(os.Environ())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Credential: &syscall.Credential{
			Uid: UID,
			Gid: UID,
		},
	}
	log.Fatal(cmd.Run())
}
