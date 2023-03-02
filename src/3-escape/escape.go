package escape

import "fmt"

func CanFastAttack(isKnightAwake bool) bool {
	return !isKnightAwake
}

func CanSpy(isKnightAwake, isArcherAwake, isPrisonerAwake bool) bool {
	return isKnightAwake || isArcherAwake || isPrisonerAwake
}

func CanSignalPrisoner(isArcherAwake, isPrisonerAwake bool) bool {
	return !isArcherAwake && isPrisonerAwake
}

func CanFreePrisoner(isKnightAwake, isArcherAwake, isPrisonerAwake, isDogPresent bool) bool {
	return (isDogPresent && !isArcherAwake) || (!isDogPresent && !isKnightAwake && !isArcherAwake && isPrisonerAwake)
}

func Run() {
	isKnightAwake := true
	isArcherAwake := false
	isPrisonerAwake := true
	isDogPresent := true

	fmt.Println("isKnightAwake: ", isKnightAwake, "\nisArcherAwake: ", isArcherAwake, "\nisPrisonerAwake: ", isPrisonerAwake, "\nisDogPresent: ", isDogPresent)
	fmt.Println("Can fast attack: ", CanFastAttack(isKnightAwake))
	fmt.Println("Can spy: ", CanSpy(isKnightAwake, isArcherAwake, isPrisonerAwake))
	fmt.Println("Can signal prisoner: ", CanSignalPrisoner(isArcherAwake, isPrisonerAwake))
	fmt.Println("Can free prisoner: ", CanFreePrisoner(isKnightAwake, isArcherAwake, isPrisonerAwake, isDogPresent))
}
