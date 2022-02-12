package solution

import (
	emoji "github.com/kyokomi/emoji"
)

func GetMessage() string {
	return emoji.Sprintf("Hello :world_map:!")
}
