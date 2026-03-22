package main

import (
	"Iam/Interface"
	"fmt"
)

func main() {
	cirkle := Interface.Cirkle{Radius: 45}
	rectangle := Interface.Rectangle{Width: 45, Height: 6}
	box:=Interface.Box{Len: 88}
	
	Interface.Use(&cirkle)
	Interface.Use(&rectangle)
	Interface.Use(&box)

	fmt.Println("Вношу изменения в репозиторий...")
}