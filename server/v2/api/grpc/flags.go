package grpc

import "fmt"

// start flags are prefixed with the server name
// as the config in prefixed with the server name
// this allows viper to properly bind the flags
func prefix(f string) string {
	return fmt.Sprintf("%s.%s", ServerName, f)
}

var FlagAddress = prefix("address")
