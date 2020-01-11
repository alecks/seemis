package main

type command struct {
	Caller string `json:"caller"`
	Function func(string) string `json:"function"`
}

var commands = []command{
	{
		Caller: "createpupil",
		Function: createPupil,
	},
	{
		Caller: "createclass",
		Function: createClass,
	},
	{
		Caller: "deletepupil",
		Function: deletePupil,
	},
	{
		Caller: "deleteclass",
		Function: deleteClass,
	},
	{
		Caller: "setattendance",
		Function: setAttendance,
	},
	{
		Caller: "getattendance",
		Function: getAttendance,
	},
	{
		Caller: "savedb",
		Function: func(text string) string {
			DB.Dump()
			return "DB dumped."
		},
	},
}
