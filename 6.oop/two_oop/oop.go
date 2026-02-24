package main

import (
	"fmt"
	"strings"
)

type User interface {
	GetUsername() string
	HasPermission(permission string) bool
	GetRole() string
}

type BaseUser struct {
	username    string
	permissions map[string]bool
	role        string
}

type BasicUser struct {
	BaseUser
}

type Moderator struct {
	BasicUser
}

type Admin struct {
	Moderator
}

func (b *BaseUser) GetUsername() string {
	return b.username
}

func (b *BaseUser) HasPermission(permission string) bool {
	return b.permissions[permission]
}

func (b *BaseUser) GetRole() string {
	return b.role
}

func (b *BaseUser) addPermission(permission string) {
	b.permissions[permission] = true

}

func NewBasicUser(username string) *BasicUser {
	u := &BasicUser{
		BaseUser: BaseUser{
			username:    username,
			permissions: make(map[string]bool),
			role:        "basic",
		},
	}
	u.addPermission("read")
	return u
}

func NewModerator(username string) *Moderator {
	m := &Moderator{
		BasicUser: *NewBasicUser(username),
	}
	m.role = "moderator"
	m.addPermission("edit")
	m.addPermission("ban_user")
	return m
}

func (m *Moderator) CanBan(target User) bool {
	return m.HasPermission("ban_user") && target.GetRole() == "basic"
}

func NewAdmin(username string) *Admin {
	a := &Admin{
		Moderator: *NewModerator(username),
	}
	a.role = "admin"
	a.addPermission("delete")
	a.addPermission("manage_roles")
	return a
}

func (a *Admin) CanManageRole(target User, newRole string) bool {
	return a.HasPermission("manage_roles") && newRole != "admin"
}

func (a *Admin) CanDelete() bool {
	return a.HasPermission("delete")
}

func CheckAccess(user User, resource string, action string) bool {
	permission := fmt.Sprintf("%s_%s", action, resource)
	return user.HasPermission(permission) || user.HasPermission(action)
}

func PrintUserPermissions(user User) {
	fmt.Printf("Пользователь: %s (роль: %s)\n", user.GetUsername(), user.GetRole())
	fmt.Println("Доступные действия:")

	actions := []string{"read", "edit", "delete", "ban_user", "manage_roles"}
	for _, action := range actions {
		if user.HasPermission(action) {
			fmt.Printf("Готово %s\n", action)
		}
	}
}

func main() {
	fmt.Println("Система управления пользователями и ролями")
	fmt.Println("=" + strings.Repeat("=", 60))

	basic := NewBasicUser("alice")
	moderator := NewModerator("bob")
	admin := NewAdmin("charlie")

	// Демонстрация информации
	fmt.Println("Информация о пользователях:")
	PrintUserPermissions(basic)
	PrintUserPermissions(moderator)
	PrintUserPermissions(admin)

	// Тестирование прав доступа
	fmt.Println("Проверка прав доступа:")

	resources := []struct {
		resource string
		action   string
	}{
		{"post", "read"},
		{"post", "edit"},
		{"post", "delete"},
		{"user", "ban_user"},
		{"role", "manage_roles"},
	}

	users := []User{basic, moderator, admin}
	userNames := []string{"BasicUser", "Moderator", "Admin"}

	for i, user := range users {
		fmt.Printf("\n%s (%s):\n", userNames[i], user.GetUsername())
		for _, r := range resources {
			can := CheckAccess(user, r.resource, r.action)
			status := "X"
			if can {
				status = "Готово!"
			}
			fmt.Printf("   %s %s.%s\n", status, r.resource, r.action)
		}
	}

	fmt.Println("Уникальные возможности:")

	fmt.Printf("\n1. %s пытается забанить %s:\n", moderator.GetUsername(), basic.GetUsername())
	if moderator.CanBan(basic) {
		fmt.Printf("Готово! %s может забанить %s\n", moderator.GetUsername(), basic.GetUsername())
	} else {
		fmt.Printf("X Недостаточно прав\n")
	}

	fmt.Printf("\n2. %s пытается забанить %s:\n", moderator.GetUsername(), admin.GetUsername())
	if moderator.CanBan(admin) {
		fmt.Printf("Готово! %s может забанить %s\n", moderator.GetUsername(), admin.GetUsername())
	} else {
		fmt.Printf("X Нельзя банить пользователей с высшей ролью\n")
	}

	fmt.Printf("\n3. %s пытается изменить роль %s на 'moderator':\n", admin.GetUsername(), basic.GetUsername())
	if admin.CanManageRole(basic, "moderator") {
		fmt.Printf("Готово! %s может управлять ролями\n", admin.GetUsername())
	} else {
		fmt.Printf("X Недостаточно прав\n")
	}

	fmt.Printf("\n4. %s пытается удалить данные:\n", admin.GetUsername())
	if admin.CanDelete() {
		fmt.Printf("Готово! %s может удалять данные\n", admin.GetUsername())
	} else {
		fmt.Printf("X Недостаточно прав\n")
	}

	// Полиморфизм: работа с интерфейсом
	fmt.Println("Полиморфизм через интерфейс User:")
	allUsers := []User{basic, moderator, admin}
	for _, u := range allUsers {
		fmt.Printf("• %s: роль=%s, есть право 'read'=%v\n",
			u.GetUsername(), u.GetRole(), u.HasPermission("read"))
	}
}
