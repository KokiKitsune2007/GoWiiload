package GoWiiload

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

const HBC_VERSION_MAJOR = 0
const HBC_VERSION_MINOR = 5

func wiiload_grab_ip() string {

	ip, exist := os.LookupEnv("WII")

	if !exist {
		fmt.Println("WII environment variable not found!")
	}

	return ip
}
func wiiload_grab_file(path string) ([]byte, error) {
	valid := map[string]bool{
		".dol":  true,
		".elf":  true,
		".wuhb": true,
		".rpx":  true,
	}

	ext := strings.ToLower(filepath.Ext(path))
	if !valid[ext] {
		fmt.Println("Unknown Extension!  Wiiload only takes dol, elf, rpx, or wuhb files!")
	}

	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading: %v\n", err)
		return nil, err
	}
	return file, nil

}
func wiiload_connect(ip string) {
	ip = wiiload_grab_ip()
	port := 4299
	addr := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", addr)

	// create HAXX header
	var header bytes.Buffer
	header.Write([]byte("HAXX"))

}
