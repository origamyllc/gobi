package resouces

func GetDatabaseConstants()(string){
	//todo:find a way to obfuscate the username and password
	Text := "{\"Host\": \"localhost\",\"Port\": \"5432\",\"User\": \"postgres\",\"Password\": \"postgres\",\"Database\": \"light\"}"
	return Text;
}