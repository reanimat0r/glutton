package glutton

import (
	"encoding/hex"
	"net"
	"time"

	log "github.com/Sirupsen/logrus"
)

// HandleTCP takes a net.Conn and peeks at the data send
func HandleTCP(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Errorf("[log.tcp ] %v", err)
		}
	}()

	conn.SetReadDeadline(time.Now().Add(10))
	host, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Errorf("[log.tcp ] %v", err)
	}
	if n > 0 {
		log.Infof("[log.tcp ] %s\n%s", host, hex.Dump(buffer[0:n]))
	}
}
