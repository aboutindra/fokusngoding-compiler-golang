package router

import (
	"encoding/json"
	"fc-golang/data"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/exec"
)

type Router struct {
}

var out, _ = exec.Command("uuidgen").Output()

func (r Router) Exec(c *fiber.Ctx) error {

	var Src data.Codes
	json.Unmarshal([]byte(c.Body()), &Src)
	createFile()
	fmt.Println("Ini code : ", Src.Code);
	writeFile(Src.Code)
	path := string(out)
	path += ".go"
	cmd1 := exec.Command("go", "run", path, string(out))
	cmd1.Stdin = os.Stdin
	stdout, err := cmd1.CombinedOutput()

	if err != nil {
		c.Status(500).SendString(err.Error())
	}

	response, _ := json.Marshal(string(stdout))

	if response == nil {
		c.Status(204).SendString("No Content");
	}

	deleteFile()
	return c.Send(response)

}

func createFile() {
	// check if file exists
	path := string(out)
	fmt.Println("uuid : ", path);
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

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
