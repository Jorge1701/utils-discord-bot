package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GetMembersIdWithRole(s *discordgo.Session, guildId string, rolId string) []string {
	var members []string
	if guild, err := s.State.Guild(guildId); err != nil {
		fmt.Printf("%s\n", err)
	} else {
		for _, member := range guild.Members {
			if SliceContains(member.Roles, rolId) {
				if !SliceContains(members, member.User.ID) {
					members = append(members, member.User.ID)
				}
			}
		}
	}
	return members
}
