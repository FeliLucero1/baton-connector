package client

// User representa un usuario en el sistema.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Project representa un proyecto con su propietario.
type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerUser   User   `json:"ownerUser"`
}

// UserWithRole representa a un usuario con un rol dentro de un proyecto.
type UserWithRole struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserRole string `json:"userRole"`
}
