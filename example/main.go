package main

import (
	"fmt"

	"github.com/mllsn/gwr"
)

func main() {
	r := gwr.Ras{"/opt/1cv8/x86_64/8.3.20.1710/rac", "192.168.234.133"}

	cluster := r.GetCluster()

	ibs := cluster.GetInfobases()

	for _, i := range ibs {
		fmt.Println(i.Name)
	}

	sns := r.GetSessions()

	for _, i := range sns {
		fmt.Println(i.Id)
	}
}
