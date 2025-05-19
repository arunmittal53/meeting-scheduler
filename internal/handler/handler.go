package handler

import (
	"meeting-scheduler/internal/model"
	"meeting-scheduler/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *service.SchedulerService
}

func NewHandler(svc *service.SchedulerService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {

	// Health Check
	r.GET("/ping", h.healthCheck)

	// User routes
	r.GET("/user/:id", h.getUser)
	r.GET("/users", h.getAllUsers)
	r.POST("/user", h.createUser)

	// Event routes
	r.GET("/event/:id", h.getEvent)
	r.POST("/event", h.createEvent)
	r.PUT("/event", h.updateEvent)
	r.DELETE("/event/:id", h.deleteEvent)

	// Availability routes
	r.GET("/event/:id/availability/:user_id", h.getAvailability)
	r.POST("/event/availability", h.addAvailability)
	r.PUT("/event/availability", h.updateAvailability)
	r.DELETE("/event/:id/availability/:user_id", h.removeAvailability)

	// Suggestions
	r.GET("/event/:id/suggestions", h.suggestSlots)
}

// ========== Health Check ==========

func (h *Handler) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// ========== User Handlers ==========

func (h *Handler) getUser(c *gin.Context) {
	userId := c.Param("id")
	userInfo, err := h.svc.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userInfo)
}
func (h *Handler) getAllUsers(c *gin.Context) {
	allUsersInfo, err := h.svc.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, allUsersInfo)
}
func (h *Handler) createUser(c *gin.Context) {
	var u model.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}
	if err := h.svc.CreateUser(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, u)
}

// ========== Event Handlers ==========

func (h *Handler) getEvent(c *gin.Context) {
	eventId := c.Param("id")
	event, err := h.svc.GetEvent(eventId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
}
func (h *Handler) createEvent(c *gin.Context) {
	var e model.Event
	if err := c.BindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.CreateEvent(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, e)
}
func (h *Handler) updateEvent(c *gin.Context) {
	var e model.Event
	if err := c.BindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.UpdateEvent(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}
func (h *Handler) deleteEvent(c *gin.Context) {
	id := c.Param("id")

	if err := h.svc.DeleteEvent(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

// ========== Availability Handlers ==========

func (h *Handler) addAvailability(c *gin.Context) {
	var av model.Availability
	if err := c.BindJSON(&av); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.svc.AddAvailability(av)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, av)
}
func (h *Handler) getAvailability(c *gin.Context) {
	eid := c.Param("id")
	uid := c.Param("user_id")
	av, err := h.svc.GetAvailability(eid, uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "availability not found"})
		return
	}
	c.JSON(http.StatusOK, av)
}
func (h *Handler) updateAvailability(c *gin.Context) {
	var av model.Availability
	if err := c.BindJSON(&av); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.svc.UpdateAvailability(av)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, av)
}
func (h *Handler) removeAvailability(c *gin.Context) {
	eid := c.Param("id")
	uid := c.Param("user_id")
	if err := h.svc.DeleteAvailability(eid, uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

// ========== Suggestion Handler ==========

func (h *Handler) suggestSlots(c *gin.Context) {
	id := c.Param("id")
	slots, _ := h.svc.SuggestSlots(id)
	c.JSON(http.StatusOK, gin.H{"suggested_slots": slots})
}
