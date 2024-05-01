package fabric

import (
	"fmt"
	"os"
	"os/exec"
)

func RegisterUser(username, password string) error {
	// Set the necessary environment variables
	os.Setenv("PATH", "/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/bin:$PATH")
	os.Setenv("FABRIC_CFG_PATH", "/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/config/")
	os.Setenv("FABRIC_CA_CLIENT_HOME", "/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/")

	// Register the user using the fabric-ca-client command
	registerCmd := exec.Command(
		"fabric-ca-client",
		"register",
		"--caname", "ca-org1",
		"--id.name", username,
		"--id.secret", password,
		"--id.type", "client",
		"--tls.certfiles", "/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/fabric-ca/org1/tls-cert.pem",
	)
	registerOutput, err := registerCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to register user: %v\n%s", err, string(registerOutput))
	}

	// Enroll the user using the fabric-ca-client command
	enrollCmd := exec.Command(
		"fabric-ca-client",
		"enroll",
		"-u", fmt.Sprintf("https://%s:%s@localhost:7054", username, password),
		"--caname", "ca-org1",
		"-M", fmt.Sprintf("/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/%s@org1.example.com/msp", username),
		"--tls.certfiles", "/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/fabric-ca/org1/tls-cert.pem",
	)
	enrollOutput, err := enrollCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to enroll user: %v\n%s", err, string(enrollOutput))
	}

	// Copy the config.yaml file to the user's MSP directory
	srcFile := "/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/msp/config.yaml"
	dstFile := fmt.Sprintf("/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/%s@org1.example.com/msp/config.yaml", username)

	srcBytes, err := os.ReadFile(srcFile)
	if err != nil {
		return fmt.Errorf("failed to read config.yaml: %v", err)
	}

	err = os.WriteFile(dstFile, srcBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config.yaml: %v", err)
	}

	return nil
}
