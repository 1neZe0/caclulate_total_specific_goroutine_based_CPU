// main.go
package CpuProfiler

/*
#include <stdio.h>
#include <sys/resource.h>
#include <time.h>

typedef struct CpuCalculator {
    struct rusage start;
    struct rusage stop;
} CpuCalculator;

void startProfiling(CpuCalculator *calculator) {
    getrusage(RUSAGE_SELF, &calculator->start);
}

void stopProfiling(CpuCalculator *calculator) {
    getrusage(RUSAGE_SELF, &calculator->stop);
}

double calculateUsage(CpuCalculator *calculator) {
    // Calculate the elapsed time in seconds
    double elapsed = calculator->stop.ru_utime.tv_sec + calculator->stop.ru_stime.tv_sec
        - calculator->start.ru_utime.tv_sec - calculator->start.ru_stime.tv_sec
        + 1e-6 * (calculator->stop.ru_utime.tv_usec + calculator->stop.ru_stime.tv_usec
            - calculator->start.ru_utime.tv_usec - calculator->start.ru_stime.tv_usec);

    // Calculate the CPU usage in percent
    return 100.0 * elapsed / CLOCKS_PER_SEC;
}
*/
import "C"
import (
	"unsafe"
)

type CpuCalculator struct {
	Start C.struct_rusage
	Stop  C.struct_rusage
}

func CpuProfiler() CpuCalculator {
	// initialize fields and return the initialized instance
	return CpuCalculator{
		// field values go here
	}
}

func (c *CpuCalculator) StartProfiling() {
	C.startProfiling((*C.CpuCalculator)(unsafe.Pointer(c)))
}

func (c *CpuCalculator) StopProfiling() {
	C.stopProfiling((*C.CpuCalculator)(unsafe.Pointer(c)))
}

func (c *CpuCalculator) CalculateUsage() float64 {
	return float64(C.calculateUsage((*C.CpuCalculator)(unsafe.Pointer(c))))
}
