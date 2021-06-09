package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/404th/testit/genproto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createContact(c *gin.Context) {
	var input genproto.ContactCreateReq

	if err := c.BindJSON(&input); err != nil {
		log.Fatalf("BAD REQUEST failed to creating contact: %s", err)
		return
	}

	id, err := h.services.ContactServiceI.CreateContact(&input)
	if err != nil {
		log.Fatalf("INTERNAL SERVER ERROR failed to creating contact: %s", err)
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) getContact(c *gin.Context) {
	var input genproto.WithID
	i64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalf("Failed to parse string to int32 while getting contact")
	}
	i := int32(i64)
	input.Id = i

	ctc, err := h.services.ContactServiceI.GetContact(&input)
	if err != nil {
		log.Fatalf("INTERNAL SERVER ERROR failed to getting contacts: %s", err)
		return
	}

	c.JSON(http.StatusOK, ctc)
}
func (h *Handler) getAllContacts(c *gin.Context) {
	var input genproto.EmptyResponse

	ctcs, err := h.services.ContactServiceI.GetAllContacts(&input)
	if err != nil {
		log.Fatalf("INTERNAL SERVER ERROR failed to getting all contacts: %s", err)
		return
	}

	c.JSON(http.StatusOK, ctcs)
}

func (h *Handler) updateContact(c *gin.Context) {
	var input genproto.Contact

	if err := c.BindJSON(&input); err != nil {
		log.Fatalf("BAD REQUEST failed to updating contact: %s", err)
		return
	}

	ctc, err := h.services.ContactServiceI.UpdateContact(&input)
	if err != nil {
		log.Fatalf("INTERNAL SERVER ERROR failed to updating contact: %s", err)
		return
	}

	c.JSON(http.StatusOK, ctc)
}
func (h *Handler) deleteContact(c *gin.Context) {
	var input genproto.WithID
	i64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		log.Fatalf("Failed to parse string to int32 while updating contact")
	}
	i := int32(i64)
	input.Id = i

	if err := c.BindJSON(&input); err != nil {
		log.Fatalf("BAD REQUEST failed to deleting contact: %s", err)
		return
	}

	id, err := h.services.ContactServiceI.DeleteContact(&input)
	if err != nil {
		log.Fatalf("INTERNAL SERVER ERROR failed to deleting contact: %s", err)
		return
	}

	c.JSON(http.StatusOK, id)
}
