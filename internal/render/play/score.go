package play

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/liqmix/ebiten-holiday-2024/internal/assets"
	"github.com/liqmix/ebiten-holiday-2024/internal/types"
	"github.com/liqmix/ebiten-holiday-2024/internal/ui"
	"github.com/tinne26/etxt"
)

// var (
// 	fadeOutHitMs       = int64(500)
// 	lastDisplayedHitMs int64
// 	prevHit            *types.HitRecord
// )

func (r *Play) drawScore(screen *ebiten.Image) {
	score := r.state.Score

	// Draw the score at the top of the screen
	p := ui.Point{
		X: 0.95,
		Y: 0.05,
	}
	opts := &ui.TextOptions{
		Align: etxt.Right,
		Scale: 1.0,
		Color: types.White,
	}
	perfectText := fmt.Sprintf(assets.String(types.L_HIT_PERFECT)+": %d", score.Perfect)
	goodText := fmt.Sprintf(assets.String(types.L_HIT_GOOD)+": %d", score.Good)
	badText := fmt.Sprintf(assets.String(types.L_HIT_BAD)+": %d", score.Bad)
	missText := fmt.Sprintf(assets.String(types.L_HIT_MISS)+": %d", score.Miss)

	hitDiffText := fmt.Sprintf("Diff: %v", score.GetLastHitRecord())
	ui.DrawTextBlockAt(screen, []string{
		perfectText,
		goodText,
		badText,
		missText,
		hitDiffText,
	}, &p, opts)

	// Draw the combo text in the center of the combo box ? bit cluttered...
	// Draw the combo above play area
	combo := r.state.Score.Combo
	if combo > 0 {
		comboText := fmt.Sprintf("%d", r.state.Score.Combo)
		ui.DrawTextAt(
			screen,
			comboText,
			&headerCenterPoint,
			// &playCenterPoint,
			&ui.TextOptions{
				Align: etxt.Center,
				Scale: 3.0,
				Color: types.White,
			},
		)
	}

	// Draw the last hit text
	lastHit := r.state.Score.GetLastHitRecord()
	if lastHit != nil {
		// currentMs := r.state.CurrentTime()

		// if lastHit != prevHit {
		// 	lastDisplayedHitMs = currentMs
		// }
		// opacity := 1.0 - float64(currentMs-lastDisplayedHitMs)/float64(fadeOutHitMs)
		hitType := lastHit.HitType
		c := hitType.Color()
		// c.A = uint8(opacity * 255)
		ui.DrawTextAt(
			screen,
			hitType.String(),
			&ui.Point{
				X: headerCenterPoint.X,
				Y: headerCenterPoint.Y + 0.05,
			},
			&ui.TextOptions{
				Align: etxt.Center,
				Scale: 1.0,
				Color: c,
			},
		)
	}
}
