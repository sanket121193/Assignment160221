package main

import (
	
	"encoding/json"
	"fmt"
	

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Patient struct {
	PatientId          string `json:"PatientId"`
	FirstName    		string `json:"firstname"`
	LastName         	string `json:"lastname"`
	Address    		string `json:"address"`
	History      		string `json:"history"`
	
}

func (s *SmartContract) AddPatient(ctx contractapi.TransactionContextInterface, patientData string) (string, error) {

	if len(patientData) == 0 {
		return "", fmt.Errorf("Please pass the correct patient data")
	}

	var patient Patient
	err := json.Unmarshal([]byte(patientData), &patient)
	if err != nil {
		return "", fmt.Errorf("Failed while unmarshling patient. %s", err.Error())
	}

	patientAsBytes, err := json.Marshal(patient)
	if err != nil {
		return "", fmt.Errorf("Failed while marshling patient. %s", err.Error())
	}

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(patient.PatientId, patientAsBytes)
}

func (s *SmartContract) GetPatientById(ctx contractapi.TransactionContextInterface, patientId string) (*Patient, error) {
	if len(patientId) == 0 {
		return nil, fmt.Errorf("Please provide correct patient Id")
	}

	patientAsBytes, err := ctx.GetStub().GetState(patientId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if patientAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", patientId)
	}

	patient := new(Patient)
	_ = json.Unmarshal(patientAsBytes, patient)

	return patient, nil

}

func (s *SmartContract) UpdatePatient(ctx contractapi.TransactionContextInterface, patientId string, history string) (string, error) {
	if len(patientId) == 0 {
		return "", fmt.Errorf("Please provide correct Patient Id")
		// return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	patientAsBytes, err := ctx.GetStub().GetState(patientId)

	if err != nil {
		return "", fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if patientAsBytes == nil {
		return "", fmt.Errorf("%s does not exist", patientId)
	}

	patient := new(Patient)
	_ = json.Unmarshal(patientAsBytes, patient)

	patient.History = history
	

	patientAsBytes, err = json.Marshal(patient)
	if err != nil {
		return "", fmt.Errorf("Failed while marshling patient. %s", err.Error())
	}

	return ctx.GetStub().GetTxID(), ctx.GetStub().PutState(patient.PatientId, patientAsBytes)

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
