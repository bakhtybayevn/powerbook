package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/middleware"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/response"
	"github.com/bakhtybayevn/powerbook/internal/core"
	"github.com/bakhtybayevn/powerbook/internal/ports"
)

// GetGiftExchanges returns all gift pairings for a competition (public).
func GetGiftExchanges(repo ports.CompetitionRepository, userRepo ports.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		compID := c.Param("id")
		if compID == "" {
			c.Error(core.New(core.ValidationError, "competition id is required"))
			return
		}

		gifts, err := repo.GetGiftExchanges(compID)
		if err != nil {
			c.Error(err)
			return
		}

		out := make([]gin.H, 0, len(gifts))
		for _, g := range gifts {
			giverName := "Unknown"
			receiverName := "Unknown"
			if u, err := userRepo.Get(g.GiverID); err == nil {
				giverName = u.DisplayName
			}
			if u, err := userRepo.Get(g.ReceiverID); err == nil {
				receiverName = u.DisplayName
			}

			out = append(out, gin.H{
				"id":                 g.ID,
				"competition_id":     g.CompetitionID,
				"giver_id":           g.GiverID,
				"giver_name":         giverName,
				"receiver_id":        g.ReceiverID,
				"receiver_name":      receiverName,
				"gift_description":   g.GiftDescription,
				"giver_confirmed":    g.GiverConfirmed,
				"receiver_confirmed": g.ReceiverConfirmed,
			})
		}

		response.JSON(c, gin.H{"gifts": out})
	}
}

// ConfirmGift allows giver to confirm giving (with description) or receiver to confirm receiving.
func ConfirmGift(repo ports.CompetitionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserID(c)
		if userID == "" {
			c.Error(core.New(core.AuthError, "unauthorized"))
			return
		}

		giftID := c.Param("giftId")
		if giftID == "" {
			c.Error(core.New(core.ValidationError, "gift id is required"))
			return
		}

		var req struct {
			GiftDescription string `json:"gift_description"`
		}
		c.ShouldBindJSON(&req)

		g, err := repo.GetGiftExchange(giftID)
		if err != nil {
			c.Error(err)
			return
		}

		if userID == g.GiverID {
			g.GiverConfirmed = true
		} else if userID == g.ReceiverID {
			g.ReceiverConfirmed = true
		} else {
			c.Error(core.New(core.AuthError, "you are not part of this gift exchange"))
			return
		}
		if req.GiftDescription != "" {
			g.GiftDescription = req.GiftDescription
		}

		if err := repo.UpdateGiftExchange(g); err != nil {
			c.Error(err)
			return
		}

		response.JSON(c, gin.H{
			"id":                 g.ID,
			"giver_confirmed":    g.GiverConfirmed,
			"receiver_confirmed": g.ReceiverConfirmed,
			"gift_description":   g.GiftDescription,
		})
	}
}
