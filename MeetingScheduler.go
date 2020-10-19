package main

import "net/http"

//MeetingScheduler : This is a function to handle "/meetings" and for calling CreateMeeting and GetMeetingID 
func MeetingScheduler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		CreateMeeting(response, request)
	}
	if request.Method == "GET" {
		GetMeetingwithTime(response, request)
	}
}
