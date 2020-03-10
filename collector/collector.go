package collector

import (
	"conscientia/internalMetrics"
	"io"
	"net"
)

func handleRequest(nCh chan  net.Conn, mCh chan []byte, intM *internalMetrics.InternalMetrics) {
	const maxIndex = 16384
	var index int
	buf := make([]byte, maxIndex * 2)
	local_buf := make([]byte, maxIndex)

	for {

		fd := <- nCh

		for {
			bytesRead, err := fd.Read(local_buf)

			if err == io.EOF {
				break
			}

			if err != nil {
				break
			}


			for i := 0; i < bytesRead; i++ {
				if local_buf[i] == '\n' {
					inData := make([]byte, index)
					copy(inData, buf)
					mCh <- inData

					index = 0
					for j := 0; j < maxIndex; j++ {
						buf[j] = 0
					}
					intM.IncrementGlobalScalar(internalMetrics.MetricReceived)
					continue
				}

				buf[index] = local_buf[i]
				index++
			}
		}

		fd.Close()
	}
}

func Collect(mCh chan []byte, intM *internalMetrics.InternalMetrics) {

	nCh := make(chan net.Conn)

	for i := 0; i < 20; i++ {
		go handleRequest(nCh, mCh, intM)
	}

	ln, err := net.Listen("tcp", ":9769")

	if err != nil {
		return
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()

		if err != nil {
			continue
		}

		nCh <- conn
	}
}