package gwr

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Ras struct {
	Bin     string
	Address string
}

type Cluster struct {
	id  string
	ras Ras
}

type Admin struct {
	user string ""
	pwd  string ""
}

type Infobase struct {
	id      string
	Name    string
	cluster Cluster
}

type Session struct {
	Id      string
	cluster Cluster
}

func (r Ras) GetCluster() Cluster {
	srcData := execute(r.Bin, r.Address, "cluster", "list")

	return Cluster{srcData[0]["cluster"], r}
}

func (c Cluster) GetInfobases() []Infobase {
	result := []Infobase{}
	srcData := execute(c.ras.Bin, c.ras.Address, "infobase", "list", fmt.Sprintf("--cluster=%s", c.id))
	for _, i := range srcData {
		result = append(result, Infobase{i["infobase"], i["name"], c})
	}
	return result
}

func (c Cluster) GetSessions() []Session {
	result := []Session{}
	srcData := execute(c.ras.Bin, c.ras.Address, "session", "list", fmt.Sprintf("--cluster=%s", c.id))
	for _, i := range srcData {
		result = append(result, Session{i["session"], c})
	}
	return result

}

func (s Session) Terminate() {
	execute(s.cluster.ras.Bin, s.cluster.ras.Address, "session", "terminate", fmt.Sprintf("--session=%s", s.Id), fmt.Sprintf("--cluster=%s", s.cluster.id))
}

func execute(args ...string) []map[string]string {
	out, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		log.Print(err)
	}

	splittedOut := strings.Split(string(out), "\n\n")
	result := make([]map[string]string, 0)

	for i, v := range splittedOut[:len(splittedOut)-1] {
		v = strings.TrimSpace(v)
		result = append(result, make(map[string]string))

		for _, a := range strings.Split(v, "\n") {
			k := strings.Split(string(a), ":")
			key := strings.TrimSpace(string(k[0]))
			val := strings.TrimSpace(string(k[1]))

			result[i][key] = val
		}
	}

	return result
}
