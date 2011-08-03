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

//#include "SDL.h"
import "C"
import "unsafe"

type Texture struct {
	Magic      *[0]byte
	Format     uint32
	Access     int32
	W          int32
	H          int32
	ModMode    int32
	BlendMode  *C.SDL_BlendMode
	R, G, B, A uint8

	Native *Texture
	Yuv    *[0]byte

	Pixels *[0]byte

	Pitch       int32
	Locked_rect Rect

	Driverdata *[0]byte

	Prev *Texture
	Next *Texture
}

func (t *Texture) Get() *C.SDL_Texture {
	return (*C.SDL_Texture)(cast(t))
}

func (t *Texture) Release() {
	C.SDL_DestroyTexture(t.Get())
}

func (t *Texture) SetAlpha(_alpha uint8) {
	C.SDL_SetTextureAlphaMod(t.Get(), C.Uint8(_alpha))
}

func (t *Texture) RenderCopy(_renderer *Renderer, _srcrect *Rect, _dstrect *Rect) {
	src := (*C.SDL_Rect)(cast(_srcrect))
	dst := (*C.SDL_Rect)(cast(_dstrect))
	C.SDL_RenderCopy(_renderer.Get(), t.Get(), src, dst)
}

func (t *Texture) SetColorMod(_red uint8, _green uint8, _blue uint8) {
	C.SDL_SetTextureColorMod(t.Get(), C.Uint8(_red), C.Uint8(_green), C.Uint8(_blue))
}

func (t *Texture) SetBlendMode(_blendmode int) {
	C.SDL_SetTextureBlendMode(t.Get(), C.SDL_BlendMode(_blendmode))
}

func GetNumRenderDrivers() int {
	return int(C.SDL_GetNumRenderDrivers())
}

type RendererInfo struct {
	name                *byte
	flags               uint32
	mod_modes           uint32
	blend_modes         uint32
	scale_modes         uint32
	num_texture_formats uint32
	texture_formats     [50]uint32
	max_texture_width   int32
	max_texture_height  int32
}

func GetRenderDriverInfo(_index int) *RendererInfo {
	var rendererInfo *RendererInfo = &RendererInfo{}
	C.SDL_GetRenderDriverInfo(C.int(_index), (*C.SDL_RendererInfo)(cast(rendererInfo)))
	return rendererInfo
}

func GetRenderDriverName(_index int) string {
	info := GetRenderDriverInfo(_index)
	strname := ""
	for c := 0; ; c++ {
		var name = uintptr(unsafe.Pointer(info.name)) + uintptr(c)
		ch := (*uint8)(cast(name))
		if *ch == uint8(0) {
			break
		}
		strname += string(*ch)
	}
	return strname
}

type Renderer C.SDL_Renderer

func (r *Renderer) Get() *C.SDL_Renderer {
	return (*C.SDL_Renderer)(cast(r))
}

func (r *Renderer) Release() {
	C.SDL_DestroyRenderer(r.Get())
}

func CreateRenderer(_window *Window, _index int) (renderer *Renderer, error string) {
	raw := C.SDL_CreateRenderer(_window.window, C.int(_index), C.SDL_RENDERER_PRESENTVSYNC|C.SDL_RENDERER_ACCELERATED)
	if raw == nil {
		error = GetError()
		return
	}
	renderer = (*Renderer)(cast(raw))
	return
}

func RenderClear(_renderer *Renderer) {
	C.SDL_RenderClear(_renderer.Get())
}

func RenderPresent(_renderer *Renderer) {
	C.SDL_RenderPresent(_renderer.Get())
}

func RenderFillRect(_renderer *Renderer, _rect Rect) {
	C.SDL_RenderFillRect(_renderer.Get(), (*C.SDL_Rect)(cast(&_rect)))
}

func SetRenderDrawColor(_renderer *Renderer, _r uint8, _g uint8, _b uint8, _a uint8) {
	C.SDL_SetRenderDrawColor(_renderer.Get(), C.Uint8(_r), C.Uint8(_g), C.Uint8(_b), C.Uint8(_a))
}

func SetRenderDrawBlendMode(_renderer *Renderer, _mode int) {
	C.SDL_SetRenderDrawBlendMode(_renderer.Get(), C.SDL_BlendMode(_mode))
}
