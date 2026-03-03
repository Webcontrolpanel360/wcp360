package agent

import (
    "bytes"
    "golang.org/x/crypto/ssh"
    "fmt"
)

// RemoteServer représente un serveur distant esclave
type RemoteServer struct {
    IP       string
    User     string
    Password string
}

// ExecuteRemote exécute une commande sur un serveur distant via SSH
func (s *RemoteServer) ExecuteRemote(command string) (string, error) {
    config := &ssh.ClientConfig{
        User: s.User,
        Auth: []ssh.AuthMethod{
            ssh.Password(s.Password),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    client, err := ssh.Dial("tcp", s.IP+":22", config)
    if err != nil {
        return "", err
    }
    defer client.Close()

    session, err := client.NewSession()
    if err != nil {
        return "", err
    }
    defer session.Close()

    var b bytes.Buffer
    session.Stdout = &b
    if err := session.Run(command); err != nil {
        return "", err
    }

    return b.String(), nil
}
