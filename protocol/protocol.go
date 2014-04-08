package protocol

import (
	"encoding/json"
	"fmt"

	"golang-refactoring.org/go-doctor/doctor"
)

type Reply struct {
	Params map[string]interface{}
}

func (r Reply) String() string {
	replyJson, _ := json.Marshal(r.Params)
	return string(replyJson)
}

type State struct {
	State      int
	Mode       string
	Dir        string
	Filesystem doctor.FileSystem
}

func Run(args []string) {

	// single command console
	if len(args) == 0 {
		runSingle()
		return
	}
	cmdList := setup()
	// list of commands
	var argJson []map[string]interface{}
	err := json.Unmarshal([]byte(args[0]), &argJson)
	if err != nil {
		printReply(Reply{map[string]interface{}{"reply": "Error", "message": err.Error()}})
		return
	}
	var state = State{1, "", "", nil}
	for i, cmdObj := range argJson {
		// has command?
		cmd, found := cmdObj["command"]
		if !found { // no command
			printReply(Reply{map[string]interface{}{"reply": "Error", "message": "Invalid JSON command"}})
			return
		}
		// valid command?
		if _, found := cmdList[cmd.(string)]; found {
			resultReply, err := cmdList[cmd.(string)].Run(&state, cmdObj)
			if err != nil {
				printReply(resultReply)
				return
			}
			// last command?
			if i == len(argJson)-1 {
				printReply(resultReply)
			}
		} else {
			printReply(Reply{map[string]interface{}{"reply": "Error", "message": "Invalid JSON command"}})
			return
		}
	}

}

func runSingle() {
	cmdList := setup()
	var state = State{0, "", "", nil}
	var input []byte
	var inputJson map[string]interface{}
	for {
		fmt.Scan(&input)
		err := json.Unmarshal(input, &inputJson)
		if err != nil {
			printReply(Reply{map[string]interface{}{"reply": "Error", "message": err.Error()}})
			continue
		}
		// check command key exists
		cmd, found := inputJson["command"]
		if !found {
			printReply(Reply{map[string]interface{}{"reply": "Error", "message": "Invalid JSON command"}})
			continue
		}
		// if close command, just exit
		if cmd == "close" {
			break
		}
		// check command is one we support
		if _, found := cmdList[cmd.(string)]; !found {
			printReply(Reply{map[string]interface{}{"reply": "Error", "message": "Invalid JSON command"}})
			continue
		}
		// everything good to run command
		result, _ := cmdList[cmd.(string)].Run(&state, inputJson) // run the command
		printReply(result)
	}
}

// little helpers
func setup() map[string]Command {
	cmds := make(map[string]Command)
	cmds["about"] = &About{}
	cmds["open"] = &Open{}
	cmds["list"] = &List{}
	cmds["setdir"] = &Setdir{}
	cmds["params"] = &Params{}
	cmds["xrun"] = &XRun{}
	return cmds
}

func printReply(reply Reply) {
	fmt.Printf("%s\n", reply)
}