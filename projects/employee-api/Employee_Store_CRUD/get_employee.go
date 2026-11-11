package employee_store_crud

import (
	"encoding/json"
	module_employee "home-3-4-5/Module_employee"
	"net/http"
)

func (store *Store) GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sqlQuery := `
	SELECT id, full_name, position
	FROM employees;
	`
	rows, err := store.Conn.Query(r.Context(), sqlQuery)
	if err != nil {
		http.Error(w, "Не удалось получить данные!", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	employees := []module_employee.Employee{}
	for rows.Next() {
		var e module_employee.Employee
		err := rows.Scan(
			&e.ID,
			&e.FullName,
			&e.Position,
		)
		if err != nil {
			http.Error(w, "Не удалось прочитать работника", http.StatusInternalServerError)
			return
		}
		employees = append(employees, e)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Не удалось получить данные", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(employees)
}
