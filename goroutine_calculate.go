package main

import (
	"awrpoj/cpu_profiler"
	"fmt"
	"golang.org/x/crypto/ssh"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	profiler := CpuProfiler.CpuProfiler()
	profiler.StartProfiling()
	// Create a WaitGroup
	var wg sync.WaitGroup
	// Increment the WaitGroup counter to indicate that we have one goroutine
	wg.Add(1)
	// Get the current time
	//f, err := os.Create("cpu.prof")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()

	//startTime := time.Now()
	go func() {
		defer wg.Done()
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 1000000000; i++ {
			_ = rand.Int()
		}
		// Start the goroutine
		f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		auth := []string{"hello:hello", "bye:bye", "salam:salam"}
		proxy := ""
		// Try connecting to the proxy without a username and password
		_, err = ssh.Dial("tcp", proxy, &ssh.ClientConfig{
			User: "",
			Auth: []ssh.AuthMethod{ssh.Password("")},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
			Timeout: time.Second * 30,
		})
		if err == nil {
			fmt.Printf("good ssh = " + proxy + "\n")
			// If the connection is successful, write the proxy's address to the good_proxies file
			_, err = f.WriteString(proxy + "\n")
			if err != nil {
				fmt.Printf("Error writing to good_proxies file: %s\n", err)
			}
			return
		}
		fmt.Println("1")
		for _, a := range auth {
			pair := strings.Split(a, ":")
			username := pair[0]
			password := pair[1]
			_, err := ssh.Dial("tcp", proxy, &ssh.ClientConfig{
				User: username,
				Auth: []ssh.AuthMethod{ssh.Password(password)},
				HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
					return nil
				},
				Timeout: time.Second * 30,
			})
			if err == nil {
				fmt.Printf("good ssh = " + proxy + ":" + username + ":" + password + "\n")
				// If the connection is successful, write the proxy's address to the good_proxies file
				_, _ = f.WriteString(proxy + ":" + username + ":" + password + "\n")
				return
			}
			time.Sleep(1)
		}
		fmt.Println("2")
		// Get the CPU profile of the current process
	}()

	// Wait for the goroutine to finish

	wg.Wait()

	profiler.StopProfiling()
	usage := profiler.CalculateUsage()
	fmt.Printf("CPU usage of goroutine: %.10f%%\n", usage)
	formattedUsage := strconv.FormatFloat(usage, 'f', 10, 64)
	parsedUsage, _ := strconv.ParseFloat(formattedUsage, 64)
	maximum_number_of_goroutines := 100 / parsedUsage
	fmt.Printf("Maximum number of goroutines: %.f\n", maximum_number_of_goroutines)
	//cpuprofiler.StopProfiling(&calculator)
	//cpuUsage := cpuprofiler.CalculateUsage(&calculator)
	//fmt.Printf("CPU usage: %.2f%%\n", cpuUsage)
	//pprof.StopCPUProfile()
	//
	//// Calculate the CPU usage of the goroutine
	//cpuUsage := calculateCPUUsage()
	//fmt.Printf("CPU usage of goroutine: %.2f%%\n", cpuUsage)
	// Second calc
	//numCPUs := runtime.NumCPU()
	//maxProcs := runtime.GOMAXPROCS(-1)
	//fmt.Printf("numCPUs = %g%%\n", float64(numCPUs))
	//fmt.Printf("maxProcs = %g%%\n", float64(maxProcs))
	//
	//cpuUsage := float64(maxProcs) / float64(numCPUs)
	//fmt.Printf("cpuUsage = %g%%\n", cpuUsage)

	// First calc
	//var memStats runtime.MemStats
	//runtime.ReadMemStats(&memStats)
	//
	//// Calculate the total pause time
	//totalPauseTime := uint64(0)
	//for _, pauseTime := range memStats.PauseNs {
	//	totalPauseTime += pauseTime
	//}
	//
	//// Calculate the elapsed time
	//elapsedTime := time.Since(startTime)
	//elapsedTimeInNanoseconds := elapsedTime.Nanoseconds()
	////fmt.Println("elapsedTimeInNanoseconds = %d%%\n", elapsedTimeInNanoseconds)
	//
	//// Calculate the percentage of CPU time that was spent paused for garbage collection
	//pausePercentage := float64(totalPauseTime) / float64(elapsedTimeInNanoseconds)
	//fmt.Printf("pausePercentage = %g%%\n", pausePercentage)
	//
	////fmt.Println("Pause percentage = %d%%\n", pausePercentage)
	//
	//// Calculate the percentage of CPU time that was available for user code
	//availableCPUPercentage := 1.0 - pausePercentage
	//// Calculate the total CPU time of the goroutine
	//
	//// Calculate the CPU usage of the goroutine as a percentage
	////fmt.Println(availableCPUPercentage)
	//fmt.Printf("CPU usage of goroutine: %g%%\n", availableCPUPercentage)
}

//func calculateCPUUsage() float64 {
//	// Load the CPU profile data from the file
//	f, err := os.Open("cpu.prof")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer f.Close()
//	p, err := pprof.Parse(f)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Calculate the total CPU time of the goroutine
//	totalTime := 0.0
//	for _, sample := range p.Profile.Sample {
//		totalTime += float64(sample.Value[0])
//	}
//
//	// Calculate the CPU usage as a percentage
//	numCPUs := runtime.NumCPU()
//	cpuUsage := totalTime / float64(numCPUs) / float64(p.Profile.Period) * 100.0
//	return cpuUsage
//}
