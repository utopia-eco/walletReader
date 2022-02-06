package consts

import (
	"bufio"
	"github.com/utopia-eco/walletReader/utils"
	"os"
	"strings"
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
		address := strings.ToLower(scanner.Text())
		BlacklistTokens[address] = true
	}
}
