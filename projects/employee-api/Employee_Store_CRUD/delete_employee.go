package employee_store_crud

import (
	"net/http"
	"strconv"
)

func (store *Store) DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Укажи числовой id", http.StatusBadRequest)
		return
	}

	sqlQuery := `
	DELETE FROM employees 
	WHERE id = $1;
	`

	_, err = store.Conn.Exec(r.Context(), sqlQuery, id)
	if err != nil {
		http.Error(w, "Не удалось удалить работника!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
