package resouces

func GetDatabaseConstants()(string){
	//todo:find a way to obfuscate the username and password
	Text := "{\"Host\": \"localhost\",\"Port\": \"5432\",\"User\": \"postgres\",\"Password\": \"postgres\",\"Database\": \"light\"}"
	return Text;
}


func GetKafkaProducerConstants()(string){
	Text := "{\"bootstrap.servers\":\"localhost\",\"Port\": \"9092\",\"Topic\": \"find-device\"}"
	return Text;
}

func GetKafkaConsumerConstants()(string) {
	Text := "{\"bootstrap.servers\": \"localhost\",\"group.id\":\"myGroup\",\"auto.offset.reset\":\"earliest\"}"
	return Text;
}

