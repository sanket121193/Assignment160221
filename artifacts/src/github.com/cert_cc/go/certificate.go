package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Certificate struct {
	CertificateId          string `json:"CertificateId"`
	FileHash    		string `json:"fileHash"`
	FilePath         	string `json:"filePath"`
	CertType    		string `json:"certType"`
	Status      		string `json:"status"`
	Certifier1   		string `json:"certifier1"`
	Certifier2   		string `json:"certifier2"`
	Certifier3   		string `json:"certifier3"`
	StudentAck  		string `json:"studentAck"`
	AllValues   		string `json:"allValues"`
	CurrentOwner		string `json:"currentOwner"`
	TransferTo		string `json:"transferTo"`
	ExpiryDate  		uint64 `json:"expiryDate"`
	IssueDate   		uint64 `json:"issueDate"`
	RecordDate  		uint64 `json:"recordDate"`
	ReissueDate 		uint64 `json:"reissueDate"`
}

func (s *SmartContract) AddActive(ctx contractapi.TransactionContextInterface, certificateData string) (string, error) {

	if len(certificateData) == 0 {
		return "", fmt.Errorf("Please pass the correct certificate data")
	}

	var certificate Certificate
	err := json.Unmarshal([]byte(certificateData), &certificate)
	if err != nil {
		return "", fmt.Errorf("Failed while unmarshling certificate. %s", err.Error())
	}

	certificateAsBytes, err := json.Marshal(certificate)
	if err != nil {
		return "", fmt.Errorf("Failed while marshling certificate. %s", err.Error())
	}

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(certificate.CertificateId, certificateAsBytes)
}

func (s *SmartContract) GetCertificateById(ctx contractapi.TransactionContextInterface, certificateId string) (*Certificate, error) {
	if len(certificateId) == 0 {
		return nil, fmt.Errorf("Please provide correct certificate Id")
	}

	certificateAsBytes, err := ctx.GetStub().GetState(certificateId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if certificateAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", certificateId)
	}

	certificate := new(Certificate)
	_ = json.Unmarshal(certificateAsBytes, certificate)

	return certificate, nil

}

func (s *SmartContract) UpdateCertificate(ctx contractapi.TransactionContextInterface, certificateId string, filePath string, fileHash string) (string, error) {
	if len(certificateId) == 0 {
		return "", fmt.Errorf("Please provide correct Certificate Id")
		// return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	certificateAsBytes, err := ctx.GetStub().GetState(certificateId)

	if err != nil {
		return "", fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if certificateAsBytes == nil {
		return "", fmt.Errorf("%s does not exist", certificateId)
	}

	certificate := new(Certificate)
	_ = json.Unmarshal(certificateAsBytes, certificate)

	certificate.FilePath = filePath
	certificate.FileHash = fileHash
	certificate.Status = "Active"

	certificateAsBytes, err = json.Marshal(certificate)
	if err != nil {
		return "", fmt.Errorf("Failed while marshling certificate. %s", err.Error())
	}

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(certificate.CertificateId, certificateAsBytes)

}

func (s *SmartContract) ReissueCertificate(ctx contractapi.TransactionContextInterface, certificateId string, expiryDate uint64, reissueDate uint64) (string, error) {
	if len(certificateId) == 0 {
		return "", fmt.Errorf("Please provide correct Certificate Id")
	}

	certificateAsBytes, err := ctx.GetStub().GetState(certificateId)

	if err != nil {
		return "", fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if certificateAsBytes == nil {
		return "", fmt.Errorf("%s does not exist", certificateId)
	}

	certificate := new(Certificate)
	_ = json.Unmarshal(certificateAsBytes, certificate)

	certificate.ExpiryDate = expiryDate
	certificate.ReissueDate = reissueDate
	certificate.Status = "Active"

	certificateAsBytes, err = json.Marshal(certificate)
	if err != nil {
		return "", fmt.Errorf("Failed while marshling certificate. %s", err.Error())
	}

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(certificate.CertificateId, certificateAsBytes)

}

func (s *SmartContract) RevokeCertificate(ctx contractapi.TransactionContextInterface, certificateId string) (string, error) {
	if len(certificateId) == 0 {
		return "", fmt.Errorf("Please provide correct Certificate Id")
		// return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	certificateAsBytes, err := ctx.GetStub().GetState(certificateId)

	if err != nil {
		return "", fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if certificateAsBytes == nil {
		return "", fmt.Errorf("%s does not exist", certificateId)
	}

	certificate := new(Certificate)
	_ = json.Unmarshal(certificateAsBytes, certificate)

	certificate.Status = "Revoked"

	certificateAsBytes, err = json.Marshal(certificate)
	if err != nil {
		return "", fmt.Errorf("Failed while marshling certificate. %s", err.Error())
	}

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(certificate.CertificateId, certificateAsBytes)

}

func (s *SmartContract) query(ctx contractapi.TransactionContextInterface, queryString string) ([]Certificate, error) {
	queryResults, err := s.getQueryResultForQueryString(ctx, queryString)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	return queryResults, nil

}

func (s *SmartContract) StudentAcknowledgement(ctx contractapi.TransactionContextInterface, certificateId string,studentAck string ) error {
        certificate := Certificate{

		CertificateId : certificateId,
		StudentAck : studentAck,
	}

	certificateJSON, err := json.Marshal(certificate)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(certificateId, certificateJSON)
}

func (s *SmartContract) getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]Certificate, error) {

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []Certificate{}

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		newCertificate := new(Certificate)

		err = json.Unmarshal(response.Value, newCertificate)
		if err != nil {
			return nil, err
		}

		results = append(results, *newCertificate)
	}
	return results, nil
}
func (s *SmartContract) GetHistoryForAsset(ctx contractapi.TransactionContextInterface, certificateId string) (string, error) {

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(certificateId)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return "", fmt.Errorf(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return string(buffer.Bytes()), nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
}
