package models

type SecretRDSJson struct {
	Username            string `json:"username"` //Alt izq + 96
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                string `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

type SignUp struct {
	UserEmail string `json:"UserEmail"` //Alt izq + 96
	UserUUID  string `json:"UserUUID"`
}
