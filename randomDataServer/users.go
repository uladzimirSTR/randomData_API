package randomDataServer

import (
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/uladzimirSTR/randomData_API/dbase"
	obj "github.com/uladzimirSTR/randomData_API/objects"
)

func GetUsers(w http.ResponseWriter, r *http.Request, pool *pgxpool.Pool) {
	query := r.URL.Query()

	// defaults
	page := 0
	limit := 100

	// page
	if rawPage := query.Get("page"); rawPage != "" {
		p, err := strconv.Atoi(rawPage)
		if err != nil || p < 0 {
			writeJSON(w, http.StatusBadRequest, obj.GetUsersResponse{
				Error: "invalid page: must be integer >= 0",
			})
			return
		}
		page = p
	}

	// limit
	if rawLimit := query.Get("limit"); rawLimit != "" {
		l, err := strconv.Atoi(rawLimit)
		if err != nil || l <= 0 {
			writeJSON(w, http.StatusBadRequest, obj.GetUsersResponse{
				Error: "invalid limit: must be integer > 0",
			})
			return
		}
		limit = l
	}

	offset := (page - 1) * limit
	if offset > 0 {
		offset += 1
	}

	// date params as strings
	dateCol := query.Get("dateCol")
	start := query.Get("start")
	end := query.Get("end")

	// whitelist for date column
	if dateCol == "" {
		dateCol = "updated_at"
	}
	if dateCol != "created_at" && dateCol != "updated_at" {
		writeJSON(w, http.StatusBadRequest, obj.GetUsersResponse{
			Error: "invalid dateCol: allowed values are created_at or updated_at",
		})
		return
	}

	params := map[string]any{
		"limit":  limit,
		"offset": offset,
	}

	if dateCol != "" {
		params["dateCol"] = dateCol
	}
	if start != "" {
		params["start"] = start
	}
	if end != "" {
		params["end"] = end
	}

	data, err := db.GetUsers(pool, "random_data", "users", params)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, obj.GetUsersResponse{
			Page:   page,
			Limit:  limit,
			Offset: offset,
			Error:  err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, obj.GetUsersResponse{
		Page:   page,
		Limit:  limit,
		Offset: offset,
		Data:   data,
	})
}
