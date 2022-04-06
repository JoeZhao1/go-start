package command

import (
	"log"
	"os"
	"os/exec"

	"github.com/JoeZhao1/go-start/framework/cobra"
)

// npm just run local go bin
var npmCommand = &cobra.Command{
	Use:   "npm",
	Short: "运行 PATH/npm 的命令",
	RunE: func(c *cobra.Command, args []string) error {
		path, err := exec.LookPath("npm")
		if err != nil {
			log.Fatalln("go-start npm: should install npm in your PATH")
		}

		cmd := exec.Command(path, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		return nil
	},
}
