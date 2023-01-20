package ssh

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"bytes"
	"code.google.com/p/go.crypto/ssh"

	"gihub.com/abdeltan/copyfiles/src/fs"
	"golang.org/x/crypto/ssh"
)

func main() {
	pk, _ := ioutil.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa")
	signer, err := ssh.ParsePrivateKey(pk)
  
	if err != nil {
	  panic(err)
	}
  
	config := &ssh.ClientConfig{
	  User: "root",
	  Auth: []ssh.AuthMethod{
		ssh.PublicKeys(signer),
	  },
	}
  
	client, err := ssh.Dial("tcp", "hostname:22", config)
	
	if err != nil {
	  panic("Failed to dial: " + err.Error())
	}
  
	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
	  panic("Failed to create session: " + err.Error())
	}
	defer session.Close()
  
	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("ls"); err != nil {
	  panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
  }

	// open file
	file, err := os.Open(originalFile)
	if err != nil {
		return err
	}
	defer file.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)

	// copy file to remote server
	go func() {
		hostIn, err := session.StdinPipe()
		if err != nil {
			fmt.Println("Failed to create stdin pipe: ", err)
		}
		defer hostIn.Close()

		io.Copy(hostIn, file)
		wg.Done()
	}()

	// run command on remote server to copy file to remote server and close stdin pipe when done
	session.Run("/usr/bin/scp -t " + remoteFile)
	wg.Wait()

	return nil
}
