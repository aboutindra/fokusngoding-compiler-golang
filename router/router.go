package router

import (
	"encoding/json"
	"fc-golang/data"
	"fmt"
	"github.com/gofiber/fiber"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

type Router struct {
}

var out, _ = exec.Command("uuidgen").Output()

func (r Router) Exec(c *fiber.Ctx) {

	var Src data.Codes
	json.Unmarshal([]byte(c.Body()), &Src)

	fmt.Print("\nIni isi body : ", c.Body(), "\n Ini isi struct nya : ", Src.Code)

	createFile()
	writeFile(Src.Code)
	path := string(out)
	path += ".go"
	cmd1 := exec.Command("go", "run", path, string(out))
	cmd1.Stdin = os.Stdin
	stdout, err := cmd1.CombinedOutput()

	if err != nil {
		c.Status(200).Send(err)
	}

	response, _ := json.Marshal(string(stdout))
	fmt.Print("\nIni out : ", string(response))
	fmt.Println("\nMaxRSS:", cmd1.ProcessState.SysUsage().(*syscall.Rusage).Maxrss)
	c.Send(response)
	deleteFile()
	PrintMemUsage()

}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func createFile() {
	// check if file exists
	path := string(out)
	path += ".go"
	fmt.Print("Ini file nya : ", path)
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(string(path))
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("File Created Successfully", path)
}

func writeFile(str string) string {
	path := string(out)
	path += ".go"
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return "Err"
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString(str)
	if isError(err) {
		return "Err"
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return "Err"
	}

	fmt.Println("File Updated Successfully.")
	return "File Updated Successfully."
}

func deleteFile() {
	path := string(out)
	path += ".go"
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("File Deleted")
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
