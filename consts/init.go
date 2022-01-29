package consts

import (
	"bufio"
	"github.com/utopia-eco/walletReader/utils"
	"os"
)

var BlacklistTokens = map[string]bool{}

func InitBlacklist() {
	file, err := os.Open("consts/Blacklist.csv")
	if err != nil {
		utils.Logger.Fatal("cannot read file err: %+v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		BlacklistTokens[scanner.Text()] = true
	}
}
