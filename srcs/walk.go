package norminette

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func WalkInsideProject(path string) {
	cSourceFile, _ := regexp.Compile(".*\\.c$")
	hSourceFile, _ := regexp.Compile(".*\\.h$")
	Makefile, _ := regexp.Compile("^Makefile$")

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		//fmt.Printf(" > Checking %q\n", info.Name())
		if info == nil{
			return fmt.Errorf("Can't open file: %q", path)
		}

		if info.IsDir() {
			//fmt.Printf(" Skipping %q... because it's a DIR\n", path)
			return nil
		}

		if cSourceFile.MatchString(info.Name()) {
			fmt.Printf("cSourceFile\t|  FOUND %q ! NORMINETTING\n", info.Name())
			// TODO: Add go routine for C file checking
			return nil
		}

		if hSourceFile.MatchString(info.Name()) {
			fmt.Printf("hSourceFile\t|  FOUND %q ! NORMINETTING\n", info.Name())
			// TODO: Add go routine for H file checking
			return nil
		}

		if Makefile.MatchString(info.Name()) && Cfg.Options.BypassMakefile == false {
			fmt.Printf("Makefile\t|  FOUND %q ! NORMINETTING\n", info.Name())
			// TODO: Add go routine for Makefile checking
			return nil
		}

		return nil
	})

	if err != nil {
		fmt.Println("Something went wrong")
		fmt.Println(err)
	}

}
