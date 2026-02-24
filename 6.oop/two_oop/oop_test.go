package main

import "testing"

// Тесты для BasicUser

func TestBasicUser_GetUsername(t *testing.T) {
	user := NewBasicUser("alice")
	expected := "alice"
	got := user.GetUsername()
	if got != expected {
		t.Errorf("GetUsername() = %q; want %q", got, expected)
	}
}

func TestBasicUser_GetRole(t *testing.T) {
	user := NewBasicUser("alice")
	expected := "basic"
	got := user.GetRole()
	if got != expected {
		t.Errorf("GetRole() = %q; want %q", got, expected)
	}
}

func TestBasicUser_HasPermission_Read(t *testing.T) {
	user := NewBasicUser("alice")
	if !user.HasPermission("read") {
		t.Error("BasicUser should have 'read' permission")
	}
}

func TestBasicUser_HasPermission_Edit(t *testing.T) {
	user := NewBasicUser("alice")
	if user.HasPermission("edit") {
		t.Error("BasicUser should NOT have 'edit' permission")
	}
}

func TestBasicUser_HasPermission_Delete(t *testing.T) {
	user := NewBasicUser("alice")
	if user.HasPermission("delete") {
		t.Error("BasicUser should NOT have 'delete' permission")
	}
}

// Тесты для Moderator

func TestModerator_GetRole(t *testing.T) {
	user := NewModerator("bob")
	expected := "moderator"
	got := user.GetRole()
	if got != expected {
		t.Errorf("GetRole() = %q; want %q", got, expected)
	}
}

func TestModerator_HasPermission_Inherited(t *testing.T) {
	user := NewModerator("bob")
	if !user.HasPermission("read") {
		t.Error("Moderator should inherit 'read' permission")
	}
}

func TestModerator_HasPermission_Edit(t *testing.T) {
	user := NewModerator("bob")
	if !user.HasPermission("edit") {
		t.Error("Moderator should have 'edit' permission")
	}
}

func TestModerator_HasPermission_BanUser(t *testing.T) {
	user := NewModerator("bob")
	if !user.HasPermission("ban_user") {
		t.Error("Moderator should have 'ban_user' permission")
	}
}

func TestModerator_HasPermission_Delete(t *testing.T) {
	user := NewModerator("bob")
	if user.HasPermission("delete") {
		t.Error("Moderator should NOT have 'delete' permission")
	}
}

func TestModerator_CanBan_BasicUser(t *testing.T) {
	mod := NewModerator("bob")
	basic := NewBasicUser("alice")
	if !mod.CanBan(basic) {
		t.Error("Moderator should be able to ban BasicUser")
	}
}

func TestModerator_CanBan_Admin(t *testing.T) {
	mod := NewModerator("bob")
	admin := NewAdmin("charlie")
	if mod.CanBan(admin) {
		t.Error("Moderator should NOT be able to ban Admin")
	}
}

func TestModerator_CanBan_AnotherModerator(t *testing.T) {
	mod1 := NewModerator("bob")
	mod2 := NewModerator("dave")
	if mod1.CanBan(mod2) {
		t.Error("Moderator should NOT be able to ban another Moderator")
	}
}

// Тесты для Admin

func TestAdmin_GetRole(t *testing.T) {
	user := NewAdmin("charlie")
	expected := "admin"
	got := user.GetRole()
	if got != expected {
		t.Errorf("GetRole() = %q; want %q", got, expected)
	}
}

func TestAdmin_HasPermission_AllInherited(t *testing.T) {
	user := NewAdmin("charlie")
	permissions := []string{"read", "edit", "ban_user"}
	for _, perm := range permissions {
		if !user.HasPermission(perm) {
			t.Errorf("Admin should have inherited permission %q", perm)
		}
	}
}

func TestAdmin_HasPermission_Delete(t *testing.T) {
	user := NewAdmin("charlie")
	if !user.HasPermission("delete") {
		t.Error("Admin should have 'delete' permission")
	}
}

func TestAdmin_HasPermission_ManageRoles(t *testing.T) {
	user := NewAdmin("charlie")
	if !user.HasPermission("manage_roles") {
		t.Error("Admin should have 'manage_roles' permission")
	}
}

func TestAdmin_CanDelete(t *testing.T) {
	admin := NewAdmin("charlie")
	if !admin.CanDelete() {
		t.Error("Admin.CanDelete() should return true")
	}
}

func TestAdmin_CanManageRole_Valid(t *testing.T) {
	admin := NewAdmin("charlie")
	basic := NewBasicUser("alice")
	if !admin.CanManageRole(basic, "moderator") {
		t.Error("Admin should be able to change BasicUser role to moderator")
	}
}

func TestAdmin_CanManageRole_CannotPromoteToAdmin(t *testing.T) {
	admin := NewAdmin("charlie")
	basic := NewBasicUser("alice")
	if admin.CanManageRole(basic, "admin") {
		t.Error("Admin should NOT be able to promote user to admin via this method")
	}
}

func TestAdmin_CanManageRole_Moderator(t *testing.T) {
	admin := NewAdmin("charlie")
	mod := NewModerator("bob")
	if !admin.CanManageRole(mod, "basic") {
		t.Error("Admin should be able to demote Moderator to basic")
	}
}

//Тесты на полиморфизм

func TestPolymorphism_UserInterface(t *testing.T) {
	users := []User{
		NewBasicUser("alice"),
		NewModerator("bob"),
		NewAdmin("charlie"),
	}

	for i, user := range users {
		if user == nil {
			t.Errorf("User %d is nil", i)
			continue
		}
		if user.GetUsername() == "" {
			t.Errorf("User %d has empty username", i)
		}
		if user.GetRole() == "" {
			t.Errorf("User %d has empty role", i)
		}
		// Все пользователи должны иметь право read
		if !user.HasPermission("read") {
			t.Errorf("User %d should have 'read' permission", i)
		}
	}
}

func TestPolymorphism_PermissionHierarchy(t *testing.T) {
	basic := NewBasicUser("alice")
	mod := NewModerator("bob")
	admin := NewAdmin("charlie")

	// BasicUser имеет только read
	if basic.HasPermission("edit") || basic.HasPermission("delete") {
		t.Error("BasicUser should only have 'read' permission")
	}

	// Moderator имеет read + edit + ban_user, но не delete
	if !mod.HasPermission("edit") || !mod.HasPermission("ban_user") {
		t.Error("Moderator should have 'edit' and 'ban_user'")
	}
	if mod.HasPermission("delete") {
		t.Error("Moderator should NOT have 'delete'")
	}

	// Admin имеет все права
	adminPerms := []string{"read", "edit", "ban_user", "delete", "manage_roles"}
	for _, perm := range adminPerms {
		if !admin.HasPermission(perm) {
			t.Errorf("Admin should have permission %q", perm)
		}
	}
}

// Тесты на инкапсуляцию

func TestEncapsulation_PermissionsNotDirectlyAccessible(t *testing.T) {
	user := NewBasicUser("alice")

	// Проверяем, что права доступны только через HasPermission
	if !user.HasPermission("read") {
		t.Error("Should be able to check 'read' via HasPermission")
	}

	// Попытка получить несуществующее право
	if user.HasPermission("nonexistent") {
		t.Error("Should return false for nonexistent permission")
	}
}

// Бенчмарки

func BenchmarkBasicUser_HasPermission(b *testing.B) {
	user := NewBasicUser("alice")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = user.HasPermission("read")
	}
}

func BenchmarkModerator_HasPermission(b *testing.B) {
	user := NewModerator("bob")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = user.HasPermission("edit")
	}
}

func BenchmarkAdmin_HasPermission(b *testing.B) {
	user := NewAdmin("charlie")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = user.HasPermission("delete")
	}
}
