package structs

import (
	"fmt"
	"github.com/uberswe/beubo/pkg/utility"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

// User is a user who can authenticate with Beubo
type User struct {
	gorm.Model
	Email       string `gorm:"size:255"`
	Password    string `gorm:"size:255"`
	Activations []UserActivation
	Roles       []*Role `gorm:"many2many:user_roles;"`
	Sites       []*Site `gorm:"many2many:user_sites;"`
}

// UserActivation is used to verify a user when signing up
type UserActivation struct {
	gorm.Model
	UserID uint
	Type   string // Email, SMS, PushNotification
	Active bool
	Code   string
}

// CanAccess checks if a user has access to the specified feature
func (u User) CanAccess(db *gorm.DB, featureKey string) bool {
	var count int64
	db.Model(&Feature{}).
		Select("features.id").
		Joins("left join role_features on role_features.feature_id = features.id").
		Joins("left join user_roles on user_roles.role_id = role_features.role_id").
		Where("features.key = ?", featureKey).
		Where("user_roles.user_id = ?", u.ID).
		Count(&count)
	return count > 0
}

// CanAccessSite checks if a user is allowed to access the specified site
func (u User) CanAccessSite(db *gorm.DB, site Site) bool {
	var count int64
	db.Model(&Site{}).
		Select("sites.id").
		Joins("left join user_sites on user_sites.site_id = sites.id").
		Where("user_sites.user_id = ?", u.ID).
		Where("sites.id = ?", site.ID).
		Count(&count)
	return count > 0
}

// HasRole checks if a user has the specified role
func (u User) HasRole(db *gorm.DB, r Role) bool {
	roles := []Role{}
	_ = db.Model(&u).Where("roles.name = ?", r.Name).Association("Roles").Find(&roles)
	if len(roles) >= 1 {
		return true
	}
	return false
}

// CreateUser is a method which creates a user using gorm
func CreateUser(db *gorm.DB, email string, password string, roles []*Role, sites []*Site) bool {
	checkUser := FetchUserByEmail(db, email)
	if checkUser.ID != 0 {
		return false
	}
	user := User{
		Email: email,
		Roles: roles,
		Sites: sites,
	}
	if err := db.Create(&user).Error; err != nil {
		fmt.Println("Could not create user")
		return false
	}
	// User password is hashed after the response is returned to improve performance
	go hashUserPassword(db, user, password)
	return true
}

// FetchUser retrieves a user from the database using the provided id
func FetchUser(db *gorm.DB, id int) User {
	user := User{}
	db.First(&user, id)
	return user
}

// FetchUserByEmail retrieves a user from the database using the provided email
func FetchUserByEmail(db *gorm.DB, email string) User {
	user := User{}
	db.Where("email = ?", email).First(&user)
	return user
}

// hashUserPassword hashes the user password using bcrypt
func hashUserPassword(db *gorm.DB, user User, password string) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		fmt.Println("Password hashing failed for user")
		return
	}

	user.Password = string(hashedPassword)

	if err := db.Save(&user).Error; err != nil {
		fmt.Println("Could not update hashed password for user")
	}
}

// AuthUser authenticates the user by verifying a username and password
func AuthUser(db *gorm.DB, email string, password string) *User {
	user := User{}

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		log.Println(err)
		return nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return nil
	} else if err != nil {
		return nil
	}

	return &user
}

// UpdateUser updates the user struct with the provided details
func UpdateUser(db *gorm.DB, id int, email string, password string, roles []*Role, sites []*Site) bool {
	user := FetchUser(db, id)
	checkUser := FetchUserByEmail(db, email)
	if checkUser.ID != 0 && checkUser.ID != user.ID {
		return false
	}
	err := db.Model(&user).Association("Roles").Clear()
	utility.ErrorHandler(err, false)
	err = db.Model(&user).Association("Sites").Clear()
	utility.ErrorHandler(err, false)
	user.Email = email
	user.Roles = roles
	user.Sites = sites
	if err := db.Save(&user).Error; err != nil {
		fmt.Println("Could not create user")
		return false
	}
	// User password is hashed after the response is returned to improve performance
	// An empty password updates the user without changing the password
	if len(password) > 8 {
		go hashUserPassword(db, user, password)
	}
	return true
}

// DeleteUser deletes a user by id
func DeleteUser(db *gorm.DB, id int) (user User) {
	db.First(&user, id)
	db.Delete(&user)
	return user
}
