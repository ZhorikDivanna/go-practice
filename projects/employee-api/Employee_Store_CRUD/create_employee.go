package employee_store_crud

import (
	"encoding/json"
	module_employee "home-3-4-5/Module_employee"
	"net/http"
)

func (store *Store) CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var e module_employee.Employee

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	sqlQuery := `
	INSERT INTO employees (id, full_name, position)
	VALUES ($1, $2, $3);
	`
	_, err := store.Conn.Exec(r.Context(), sqlQuery, e.ID, e.FullName, e.Position)
	if err != nil {
		http.Error(w, "Не удалось сохранить работника", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
