package handler

import (
	"encoding/json"
	"jeevan/jobportal/internal/models"
	"net/http"
	"strconv"

	"jeevan/jobportal/internal/middleware"

	"jeevan/jobportal/internal/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (h *Handler) ViewJobByID(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	id := c.Param("id")

	jid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	jobData, err := h.Service.ViewJobById(ctx, jid)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobData)

}

func (h *Handler) ViewAllJobs(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	jobDatas, err := h.Service.ViewAllJobs(ctx)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobDatas)

}

func (h *Handler) ViewJob(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	id := c.Param("id")

	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	jobData, err := h.Service.ViewJobByCid(ctx, cid)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobData)

}
func (h *Handler) AddJob(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	id := c.Param("cid")

	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	var jobData models.Jobs

	err = json.NewDecoder(c.Request.Body).Decode(&jobData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid name, location and field",
		})
		return
	}

	jobData, err = h.Service.AddJobDetails(ctx, jobData, cid)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobData)

}
