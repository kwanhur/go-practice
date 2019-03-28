package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	flagSet   = flag.NewFlagSet("affinity", flag.ExitOnError)
	processor = flagSet.Int("processor", -1, "how many processors")
	cpu       = flagSet.String("cpu", "", "which cpu for affinity,format:1-3|1,3,4|1")
)

func main() {
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	processors, cpus := *processor, *cpu
	if processors <= 0 {
		log.Fatalln("no processor could be affinity")
	}
	if cpus == "" {
		log.Fatalln("please set which cpu be affinity")
	}
	//format 1-3
	if strings.Index(cpus, "-") > 0 {
		cpuIds := strings.Split(cpus, "-")
		if len(cpuIds) != 2 {
			log.Fatalf("format should be 1-3,but get %s", cpus)
		}

		p, err := strconv.Atoi(cpuIds[0])
		if err != nil {
			log.Fatalf("not valid cpu id with %s", cpuIds[0])
		}
		q, err := strconv.Atoi(cpuIds[1])
		if err != nil {
			log.Fatalf("not valid cpu id with %s", cpuIds[1])
		}
		if p > q {
			log.Fatalf("not valid cpu id range with:%s", cpus)
		}
		if q > processors {
			log.Fatalf("cpu id:%d should not greater than processors:%d", q, processors)
		}
		if p <= 0 {
			log.Fatalf("not valid cpu id range with:%s", cpus)
		}
		if p == q {
			// 40 3-3 0[37]100
			affinity := strings.Repeat("0", processors-q) + "1" + strings.Repeat("0", q-1)
			log.Printf("cpu affinity:%s ", affinity)
		} else {
			// 40 11-15
			affinity := ""
			for p <= q {
				affinity += strings.Repeat("0", processors-p) + "1" + strings.Repeat("0", p-1) + " "
				p += 1
			}
			log.Printf("cpu affinity:%s ", affinity)
		}
	} else if strings.Index(cpus, ",") > 0 { //format 1,2,3
		cpuIds := strings.Split(cpus, ",")
		affinity := ""
		for _, p := range cpuIds {
			q, err := strconv.Atoi(p)
			if err != nil || q > processors || q <= 0 {
				log.Fatalf("not valid cpu id with %s", p)
			}
			affinity += strings.Repeat("0", processors-q) + "1" + strings.Repeat("0", q-1) + " "
		}
		log.Printf("cpu affinity:%s ", affinity)
	} else { //format 1
		q, err := strconv.Atoi(cpus)
		if err != nil || q > processors || q <= 0 {
			log.Fatalf("not valid cpu id with %s", cpus)
		}
		affinity := strings.Repeat("0", processors-q) + "1" + strings.Repeat("0", q-1)
		log.Printf("cpu affinity:%s ", affinity)
	}
}
