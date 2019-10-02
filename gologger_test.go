// @author nikoeleison
package gologger

import "testing"

// @public

func BenchmarkGologger(b *testing.B) {
	logger := New(
		"./log",
		"benchmark_golang",
	)
	defer logger.Kill()

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sample(logger, i)
	}
}

// @private

func sample(logger *Driver, i int) {
	logger.Info("print ", i)
	logger.Infof("print:%d", i)

	logger.Error("print ", i)
	logger.Errorf("print:%d", i)

	logger.Debug("print ", i)
	logger.Debugf("print:%d", i)

	logger.Panic("print ", i)
	logger.Panicf("print:%d", i)
}
