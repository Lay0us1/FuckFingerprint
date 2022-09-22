package runner

import (
	"fmt"
	"github.com/logrusorgru/aurora"
)

/**
  @author: yhy
  @since: 2022/9/17
  @desc: //TODO
**/

const banner = `
  █████▒█    ██  ▄████▄   ██ ▄█▀  █████▒██▓███  
▓██   ▒ ██  ▓██▒▒██▀ ▀█   ██▄█▒ ▓██   ▒▓██░  ██▒
▒████ ░▓██  ▒██░▒▓█    ▄ ▓███▄░ ▒████ ░▓██░ ██▓▒
░▓█▒  ░▓▓█  ░██░▒▓▓▄ ▄██▒▓██ █▄ ░▓█▒  ░▒██▄█▓▒ ▒
░▒█░   ▒▒█████▓ ▒ ▓███▀ ░▒██▒ █▄░▒█░   ▒██▒ ░  ░
 ▒ ░   ░▒▓▒ ▒ ▒ ░ ░▒ ▒  ░▒ ▒▒ ▓▒ ▒ ░   ▒▓▒░ ░  ░
 ░     ░░▒░ ░ ░   ░  ▒   ░ ░▒ ▒░ ░     ░▒ ░     
 ░ ░    ░░░ ░ ░ ░        ░ ░░ ░  ░ ░   ░░       
          ░     ░ ░      ░  ░                   
                ░
`

// Version is the current version of FuckFingerprint
const Version = "1.1.0"

// showBanner is used to show the banner to the user
func showBanner() {
	fmt.Println("\t" + aurora.Green(banner).String())
	fmt.Println("\t\t\t" + aurora.Red("v"+Version).String())
	fmt.Println("\t\t\t\t" + aurora.Blue("https://github.com/yhy0/FuckFingerprint").String() + "\n\n")

	fmt.Println(aurora.Red("Use with caution. You are responsible for your actions").String())
	fmt.Println(aurora.Red("Developers assume no liability and are not responsible for any misuse or damage.").String())
}
