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

// @Summary Get user by ID
// @Description Get user details by user ID
// @Tags user
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 404 {object} map[string]string
// @Router /user/{id} [get]
func (h *Handler) getUser(c *gin.Context) {
	userId := c.Param("id")
	userInfo, err := h.svc.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userInfo)
}

// @Summary Get all users
// @Description Retrieve a list of all registered users
// @Tags user
// @Produce json
// @Success 200 {array} model.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	allUsersInfo, err := h.svc.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, allUsersInfo)
}

// @Summary Create a new user
// @Description Register a new user with name and ID
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.User true "User to create"
// @Success 201 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user [post]
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
// @Summary Get event by ID
// @Description Retrieve event details by event ID
// @Tags event
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 {object} model.Event
// @Failure 404 {object} map[string]string
// @Router /event/{id} [get]
func (h *Handler) getEvent(c *gin.Context) {
	eventId := c.Param("id")
	event, err := h.svc.GetEvent(eventId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
}

// @Summary Create a new event
// @Description Create an event with title, duration, and time slots
// @Tags event
// @Accept json
// @Produce json
// @Param event body model.Event true "Event to create"
// @Success 201 {object} model.Event
// @Failure 400 {object} map[string]string
// @Router /event [post]
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

// @Summary Update an event
// @Description Update event details (slots, title, duration, etc.)
// @Tags event
// @Accept json
// @Produce json
// @Param event body model.Event true "Event to update"
// @Success 200 {object} model.Event
// @Failure 400 {object} map[string]string
// @Router /event [put]
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

// @Summary Delete an event
// @Description Delete event by ID
// @Tags event
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /event/{id} [delete]
func (h *Handler) deleteEvent(c *gin.Context) {
	id := c.Param("id")

	if err := h.svc.DeleteEvent(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

// ========== Availability Handlers ==========

// @Summary Add user availability
// @Description Add time slots a user is available for an event
// @Tags availability
// @Accept json
// @Produce json
// @Param availability body model.Availability true "Availability to add"
// @Success 201 {object} model.Availability
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /event/availability [post]
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

// @Summary Get user availability
// @Description Get a user's availability for a given event
// @Tags availability
// @Produce json
// @Param id path string true "Event ID"
// @Param user_id path string true "User ID"
// @Success 200 {object} model.Availability
// @Failure 404 {object} map[string]string
// @Router /event/{id}/availability/{user_id} [get]
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

// @Summary Update user availability
// @Description Update time slots a user is available for an event
// @Tags availability
// @Accept json
// @Produce json
// @Param availability body model.Availability true "Availability to update"
// @Success 200 {object} model.Availability
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /event/availability [put]
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

// @Summary Remove user availability
// @Description Remove a user's availability for a specific event
// @Tags availability
// @Produce json
// @Param id path string true "Event ID"
// @Param user_id path string true "User ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /event/{id}/availability/{user_id} [delete]
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

// @Summary Suggest meeting slots
// @Description Suggest best time slots for a meeting based on availability
// @Tags suggestion
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 {object} map[string][]model.SlotSuggestion
// @Router /event/{id}/suggestions [get]
func (h *Handler) suggestSlots(c *gin.Context) {
	id := c.Param("id")
	slots, _ := h.svc.SuggestSlots(id)
	c.JSON(http.StatusOK, gin.H{"suggested_slots": slots})
}
