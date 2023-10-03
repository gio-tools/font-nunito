package nunito

import (
	"sync"

	"gio.tools/fonts/nunito/nunitoblack"
	"gio.tools/fonts/nunito/nunitoblackitalic"
	"gio.tools/fonts/nunito/nunitobold"
	"gio.tools/fonts/nunito/nunitobolditalic"
	"gio.tools/fonts/nunito/nunitoextrabold"
	"gio.tools/fonts/nunito/nunitoextrabolditalic"
	"gio.tools/fonts/nunito/nunitoextralight"
	"gio.tools/fonts/nunito/nunitoextralightitalic"
	"gio.tools/fonts/nunito/nunitoitalic"
	"gio.tools/fonts/nunito/nunitolight"
	"gio.tools/fonts/nunito/nunitolightitalic"
	"gio.tools/fonts/nunito/nunitomedium"
	"gio.tools/fonts/nunito/nunitomediumitalic"
	"gio.tools/fonts/nunito/nunitoregular"
	"gio.tools/fonts/nunito/nunitosemibold"
	"gio.tools/fonts/nunito/nunitosemibolditalic"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

var (
	once       sync.Once
	collection []font.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(nunitoblack.TTF)
		register(nunitoblackitalic.TTF)
		register(nunitobold.TTF)
		register(nunitobolditalic.TTF)
		register(nunitoextrabold.TTF)
		register(nunitoextrabolditalic.TTF)
		register(nunitoextralight.TTF)
		register(nunitoextralightitalic.TTF)
		register(nunitoitalic.TTF)
		register(nunitolight.TTF)
		register(nunitolightitalic.TTF)
		register(nunitomedium.TTF)
		register(nunitomediumitalic.TTF)
		register(nunitoregular.TTF)
		register(nunitosemibold.TTF)
		register(nunitosemibolditalic.TTF)
		// Ensure that any outside appends will not reuse the backing store.
		n := len(collection)
		collection = collection[:n:n]
	})
	return collection
}

func register(src []byte) {
	faces, err := opentype.ParseCollection(src)
	if err != nil {
		panic("failed to parse font: " + err.Error())
	}
	collection = append(collection, faces[0])
}
