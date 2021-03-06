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
package main

import (
	pul "pulogic"
)

type Outfit struct {
	data []int
}

func NewOutfit() Outfit {
	return Outfit{data: make([]int, 6)}
}

func NewOutfitExt(_head, _nek, _upper, _lower, _feet int) Outfit {
	outfit := NewOutfit()
	outfit.data[pul.OUTFIT_HEAD] = _head
	outfit.data[pul.OUTFIT_NEK] = _nek
	outfit.data[pul.OUTFIT_UPPER] = _upper
	outfit.data[pul.OUTFIT_LOWER] = _lower
	outfit.data[pul.OUTFIT_FEET] = _feet

	return outfit
}

func (o Outfit) GetData() []int {
	return o.data
}

func (o Outfit) SetOutfitKey(_part pul.OutfitPart, _key int) {
	o.data[_part] = _key
}

func (o Outfit) GetOutfitKey(_part pul.OutfitPart) int {
	return o.data[_part]
}

func (o Outfit) SetOutfitStyle(_part pul.OutfitPart, _style int) {
	o.data[_part] = (_style << 24) | o.GetOutfitColour(_part)
}

func (o Outfit) GetOutfitStyle(_part pul.OutfitPart) int {
	key := o.data[_part]
	return int((int8)(key >> 24))
}

func (o Outfit) SetOutfitColour(_part pul.OutfitPart, _colour int) {
	o.data[_part] = (o.GetOutfitStyle(_part) << 24) | _colour
}

func (o Outfit) GetOutfitColour(_part pul.OutfitPart) int {
	key := o.data[_part]
	return (key & 0xFFFFFF)
}
