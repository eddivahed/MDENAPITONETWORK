package fabric

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	// "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func InitializeSDK(configPath string) (*fabsdk.FabricSDK, error) {
	sdk, err := fabsdk.New(config.FromFile(configPath))
	if err != nil {
		return nil, fmt.Errorf("failed to create Fabric SDK: %v", err)
	}
	return sdk, nil
}

func InvokeChaincode(channelID, chaincodeName, functionName string, args []string, username string) ([]byte, error) {
	// Set the necessary environment variables
	os.Setenv("PATH", "/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/bin:$PATH")
	os.Setenv("FABRIC_CFG_PATH", "/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/config/")
	os.Setenv("FABRIC_CA_CLIENT_HOME", "/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/")
	os.Setenv("CORE_PEER_TLS_ENABLED", "true")
	os.Setenv("CORE_PEER_LOCALMSPID", "Org1MSP")
	os.Setenv("CORE_PEER_MSPCONFIGPATH", fmt.Sprintf("/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/%s@org1.example.com/msp", username))
	os.Setenv("CORE_PEER_TLS_ROOTCERT_FILE", "/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt")
	os.Setenv("CORE_PEER_ADDRESS", "localhost:7051")
	os.Setenv("TARGET_TLS_OPTIONS", "-o localhost:7050 --ordererTLSHostnameOverride localhost --tls --cafile \"/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem\" --peerAddresses localhost:7051 --tlsRootCertFiles \"/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt\" --peerAddresses localhost:9051 --tlsRootCertFiles \"/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt\"")	
	// Construct the peer chaincode invoke command
	cmd := exec.Command("peer", "chaincode", "invoke",
		"--channelID", channelID,
		"--name", chaincodeName,
		"--ctor", fmt.Sprintf("{\"Args\":[\"%s\",\"%s\"]}", functionName, strings.Join(args, "\",\"")),
		"--tls",
		"--cafile", "/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem",
	)


	// Execute the command and capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to invoke chaincode: %v\nOutput: %s", err, string(output))
	}

	return output, nil
}

func QueryChaincode(channelID, chaincodeName, functionName string, args []string, username string) ([]byte, error) {
	// Set the necessary environment variables
	os.Setenv("PATH", "/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/bin:$PATH")
	os.Setenv("FABRIC_CFG_PATH", "/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/config/")
	os.Setenv("FABRIC_CA_CLIENT_HOME", "/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/")
	os.Setenv("CORE_PEER_TLS_ENABLED", "true")
	os.Setenv("CORE_PEER_LOCALMSPID", "Org1MSP")
	os.Setenv("CORE_PEER_MSPCONFIGPATH", fmt.Sprintf("/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/users/%s@org1.example.com/msp", username))
	os.Setenv("CORE_PEER_TLS_ROOTCERT_FILE", "/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt")
	os.Setenv("CORE_PEER_ADDRESS", "localhost:7051")
	os.Setenv("TARGET_TLS_OPTIONS", "-o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile \"/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem\" --peerAddresses localhost:7051 --tlsRootCertFiles \"/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt\" --peerAddresses localhost:9051 --tlsRootCertFiles \"/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt\"")

	// Construct the peer chaincode query command
	cmd := exec.Command("peer", "chaincode", "query",
		"--channelID", channelID,
		"--name", chaincodeName,
		"--ctor", fmt.Sprintf("{\"Args\":[\"%s\",\"%s\"]}", functionName, strings.Join(args, "\",\"")),
		"--tls",
		"--cafile", "/Users/amirh/Documents/BlogsContents/MDENBlockchain/fabric-samples/test-network/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem",
	)

	// Execute the command and capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to invoke chaincode: %v\nOutput: %s", err, string(output))
	}

	return output, nil
}
