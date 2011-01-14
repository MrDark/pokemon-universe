/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.*/

/*
Simple SDL 1.3 wrapper
The goal is not to make a complete SDL wrapper, but to wrap only the SDL functions that the PU client needs.
*/
package sdl

//#include "SDL_ttf.h"
import "C"
import "unsafe"

func InitTTF() int {
	return int(C.TTF_Init())
}

func QuitTTF() {
	C.TTF_Quit();
}

type Font C.TTF_Font

func LoadFont(_file string, _size int) *Font {
	cfile := C.CString(_file); defer C.free(unsafe.Pointer(cfile))
	font := C.TTF_OpenFont(cfile, C.int(_size))
	if font == nil {
		return nil
	}
	return (*Font)(cast(font))
}

func (f *Font) Get() *C.TTF_Font {
	return (*C.TTF_Font)(cast(f))
}

func (f *Font) Release() {
	C.TTF_CloseFont(f.Get())
}

func (f *Font) RenderText_Blended(_text string, color Color) *Surface {
	ccolor := (*C.SDL_Color)(cast(&color))
	ctext := C.CString(_text); defer C.free(unsafe.Pointer(ctext))
	sf := C.TTF_RenderText_Blended(f.Get(), ctext, *ccolor)
	return (*Surface)(cast(sf))
}

func (f *Font) GetHeight() int {
	return int(C.TTF_FontHeight(f.Get()))
}

func (f *Font) GetMetrics(_ch uint16) (int,int,int,int,int) {
	var minx, maxx, miny, maxy, advance C.int
	C.TTF_GlyphMetrics(f.Get(), C.Uint16(_ch), (*C.int)(cast(&minx)), (*C.int)(cast(&maxx)), (*C.int)(cast(&miny)), (*C.int)(cast(&maxy)), (*C.int)(cast(&advance)));
	return int(minx), int(maxx), int(miny), int(maxy), int(advance)
}

func (f *Font) GetKerning(_previous int, _current int) int {
	return int(C.TTF_GetFontKerningSize(f.Get(),C.int(_previous), C.int(_current)))
}

func (f *Font) SetStyle(_bold bool, _italic bool, _underline bool) {
    var flags int

    if(_bold) {
    	flags |= C.TTF_STYLE_BOLD;
    }
    if(_italic) {
        flags |= C.TTF_STYLE_ITALIC;
	}
    if(_underline) {
    	flags |= C.TTF_STYLE_UNDERLINE;
    }
        
    if(flags == 0) {
    	flags = C.TTF_STYLE_NORMAL;
    }
        
	C.TTF_SetFontStyle(f.Get(), C.int(flags));
}
