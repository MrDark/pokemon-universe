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

type PU_BattleEvent_Sleep struct {
	ticks uint32
}

func NewBattleEvent_Sleep(_ticks uint32) *PU_BattleEvent_Sleep {
	return &PU_BattleEvent_Sleep{ticks: _ticks}
}

func (e *PU_BattleEvent_Sleep) Execute() {
	if g_game.battle != nil {
		g_game.battle.Wait(e.ticks)
	}
}
