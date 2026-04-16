package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/middleware"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

func requireAdmin(c *gin.Context, userRepo ports.UserRepository) bool {
	userID := middleware.GetUserID(c)
	if userID == "" {
		c.Error(core.New(core.AuthError, "unauthorized"))
		return false
	}
	u, err := userRepo.Get(userID)
	if err != nil || !u.IsAdmin {
		c.Error(core.New(core.AuthError, "admin access required"))
		return false
	}
	return true
}

func AdminListUsers(userRepo ports.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !requireAdmin(c, userRepo) {
			return
		}
		users, err := userRepo.ListAll()
		if err != nil {
			c.Error(err)
			return
		}
		out := make([]gin.H, 0, len(users))
		for _, u := range users {
			out = append(out, gin.H{
				"id":              u.ID,
				"email":           u.Email,
				"display_name":    u.DisplayName,
				"xp":              u.XP,
				"level":           u.Level(),
				"level_name":      u.LevelName(),
				"streak_current":  u.StreakCurrentDays,
				"total_minutes":   u.TotalMinutes,
				"telegram_handle": u.TelegramHandle,
				"is_admin":        u.IsAdmin,
			})
		}
		response.JSON(c, gin.H{"users": out})
	}
}

func AdminDeleteUser(userRepo ports.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !requireAdmin(c, userRepo) {
			return
		}
		uid := c.Param("id")
		if uid == "" {
			c.Error(core.New(core.ValidationError, "user id required"))
			return
		}
		if err := userRepo.Delete(uid); err != nil {
			c.Error(err)
			return
		}
		response.JSON(c, gin.H{"deleted": uid})
	}
}
