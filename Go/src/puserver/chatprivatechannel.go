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
	"fmt"
	
	pnet "network"
)

type ChatInvitedMap map[uint64]*Player

type ChatPrivateChannel struct {
	ChatChannel
	
	Invited	ChatInvitedMap
	Owner	uint64
}

func NewChatPrivateChannel(_id int, _name string) *ChatPrivateChannel {
	channel := &ChatPrivateChannel { Invited: make(ChatInvitedMap),
								 Owner: 0 }								 
	channel.Id = _id
	channel.Name = _name
	channel.Users = make(ChatUsersMap)
	
	return channel
}

func (c *ChatPrivateChannel) ChatChannelType() int {
	return 2;
}

func (c *ChatPrivateChannel) GetOwner() uint64 {
	return c.Owner
}

func (c *ChatPrivateChannel) IsInvited(_player *Player) bool {
	if _player == nil {
		return false
	}
	
	if _player.GetUID() == c.Owner {
		return true
	}
	
	_, found := c.Invited[_player.GetUID()]
	return found
}

func (c *ChatPrivateChannel) AddInvited(_player *Player) bool {
	if _, found := c.Invited[_player.GetUID()]; found {
		return false
	}
	
	c.Invited[_player.GetUID()] = _player
	
	return true
}

func (c *ChatPrivateChannel) RemoveInvited(_player *Player) bool {
	_, found := c.Invited[_player.GetUID()]
	if found {
		delete(c.Invited, _player.GetUID())
	}
	
	return found
}

func (c *ChatPrivateChannel) InvitePlayer(_player *Player, _invitedPlayer *Player) {
	if _player.GetUID() != _invitedPlayer.GetUID() && c.AddInvited(_invitedPlayer) {
		msg := fmt.Sprintf("%v invites you to %v private chat channel.", _player.GetName(), "his") // TODO: His/Her
		_invitedPlayer.sendTextMessage(pnet.MSG_INFO_DESCR, msg)
		
		msg = fmt.Sprintf("%v has been invited.", _invitedPlayer.GetName())
		_player.sendTextMessage(pnet.MSG_INFO_DESCR, msg)
	}
}

func (c *ChatPrivateChannel) ExcludePlayer(_player *Player, _excludePlayer *Player) {
	if _player.GetUID() != _excludePlayer.GetUID() && c.RemoveInvited(_excludePlayer) {
		msg := fmt.Sprintf("%v has been excluded.", _excludePlayer.GetName())
		_player.sendTextMessage(pnet.MSG_INFO_DESCR, msg)
		
		c.RemoveUser(_excludePlayer, true)
	}
}

func  (c *ChatPrivateChannel) CloseChannel() {
	for _, value := range(c.Users) {
		value.sendClosePrivateChat(c.Id)
	}
}