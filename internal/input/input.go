package input

import (
	"github.com/liqmix/slaptrax/internal/user"
)

var (
	M      = newMouse()
	K      = newKeyboard()
	Update = func() {
		M.update()
		K.update()
	}
	JustActioned = func(a Action) bool {
		keys, ok := actionToKey[a]
		if !ok {
			return false
		}
		return K.AreAny(keys, JustPressed)
	}
	IsActioned = func(a Action) bool {
		keys, ok := actionToKey[a]
		if !ok {
			return false
		}
		return K.AreAny(keys, Held)
	}
	NotActioned = func(a Action) bool {
		keys, ok := actionToKey[a]
		if !ok {
			return false
		}
		return !K.AreAny(keys, Held)
	}
)

func InitInput() {
	SetTrackKeys(TrackKeyConfig(user.S().KeyConfig))
}
