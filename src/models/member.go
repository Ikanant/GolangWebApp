package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
)

type Member struct {
	email string
	id int
	password string
	firstName string
}

func (this *Member) Email() string{
	return this.email
}
func (this *Member) SetEmail(value string) {
	this.email = value
}
func (this *Member) Id() int {
	return this.id
}
func (this *Member) SetId(value int) {
	this.id = value
}
func (this *Member) Password() string {
	return this.password
}
func (this *Member) SetPassword (value string) {
	this.password = value
}
func (this *Member) FirstName() string {
	return this.firstName
}
func (this *Member) SetFirstName(value string) {
	this.firstName = value
}

func GetMember(email string, password string) (Member, error) {
	db, err := getDBConnection()
	result := Member{}
	
	if err == nil {
		defer db.Close()
		pwd := sha256.Sum256([]byte(password))
		
		row := db.QueryRow(`SELECT id, email, first_name FROM member WHERE email = $1 AND password = $2`, email, hex.EncodeToString(pwd[:]))
		
		err = row.Scan(&result.id, &result.email, &result.firstName)
		
		if err == nil {
			return result, nil
		} else {
			return result, errors.New("Unable to find Member with info provided")
		}
	} else {
		return result, errors.New("Couldn't connect to the database")
	}
}

func GetMemberById(id int) (Member, error) {
	db, err := getDBConnection()
	result := Member{}
	
	if err == nil {
		defer db.Close()
		
		row := db.QueryRow(`SELECT id, email, first_name FROM member WHERE id = $1;`, id)
		
		err = row.Scan(&result.id, &result.email, &result.firstName)
		
		if err == nil {
			return result, nil
		} else {
			return result, errors.New("Unable to find Member with info provided")
		}
	} else {
		return result, errors.New("Couldn't connect to the database")
	}
}

func InsertMember(firstName string, email string, password string) error{
	db, err := getDBConnection()
	
	if err == nil {
		defer db.Close()
		pwd := sha256.Sum256([]byte(password))
		
		db.QueryRow(`INSERT INTO member( id, email, password, first_name) VALUES (DEFAULT, $1, $2, $3);`, email, hex.EncodeToString(pwd[:]), firstName)
		
		return nil
	} else {
		return errors.New("Couldn't connect to the database")
	}
}





type Session struct {
	id int
	memberId int
	sessionId string
}

func (this *Session) Id() int{
	return this.id
}
func (this *Session) SetId(value int){
	this.id = value
}
func (this *Session) MemberId() int{
	return this.memberId
}
func (this *Session) SetMemberId (value int){
	this.memberId = value
}
func (this *Session) SessionId () string{
	return this.sessionId
}
func (this *Session) SetSessionId (value string){
	this.sessionId = value
}

func CreateSession (member Member) (Session, error){
	result := Session{}
	result.memberId = member.Id()
	
	sessionId := sha256.Sum256([]byte(member.Email() + time.Now().Format("12:00:00")))
	result.sessionId = hex.EncodeToString(sessionId[:])
	
	db, err := getDBConnection()
	
	if err == nil {
		defer db.Close()
		
		err := db.QueryRow(`INSERT INTO session (member_id, session_id) VALUES ($1, $2) RETURNING id`, member.Id(), result.sessionId).Scan(&result.id)
		
		if err == nil {
			return result, nil
		} else {
			return Session{}, errors.New("Unable to save session into the database")
		}
	} else {
		return Session{}, errors.New("Couldn't connect to the database")
	}
	
}

func RemoveSession (member_id string) bool{
	db, err := getDBConnection()
	
	if err == nil {
		defer db.Close()
		
		nerr := db.QueryRow(`DELETE FROM session WHERE member_id = $1;`, member_id);
			
		if nerr == nil {
			return true
		}
		return true
	}
	
	return false
}