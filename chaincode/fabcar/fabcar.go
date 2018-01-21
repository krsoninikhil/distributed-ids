/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"reflect"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the nid structure, with 4 properties.  Structure tags are used by encoding/json library
type NID struct {
	NIDN string `json:"nidn"`
	Name string `json:"name"`
	Age string `json:"age"`
	DOB string `json:"dob"`
	Contact string `json:"contact"`
	Owner string `json:"owner"`
}

func getField(v *NID, field string) string {
    r := reflect.ValueOf(v)
    f := reflect.Indirect(r).FieldByName(field)
    return f.Interface().(string)
}

/*
 * The Init method is called when the Smart Contract "fabnid" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabnid"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryNID" {
		return s.queryNID(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createNID" {
		return s.createNID(APIstub, args)
	} else if function == "queryAllNIDs" {
		return s.queryAllNIDs(APIstub)
	} else if function == "changeNIDOwner" {
		return s.changeNIDOwner(APIstub, args)
	} else if function == "authorize" {
		return s.authorize(APIstub, args)
	} else if function == "queryNID" {
		return s.queryNID(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryNID(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	nidAsBytes, _ := APIstub.GetState(args[1])
	nid := NID{}
        json.Unmarshal(nidAsBytes, &nid)
	fmt.Println("queryNID", nid.Name)

	key := fmt.Sprintf("auth_%s_%s", args[1], args[0])
	allowedBytes, _ := APIstub.GetState(key)
	fmt.Println("queryNID2", allowedBytes)

	var allowed []string
	json.Unmarshal(allowedBytes, &allowed)

	fmt.Println("queryNID3", allowed)
	var res []string
	for _, field := range allowed {
		val := getField(&nid, field)
		res = append(res, val)
	}
	fmt.Println("queryNID4", res)
	resBytes, _ := json.Marshal(res)
	return shim.Success(resBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	NIDs := []NID{
		NID{NIDN: "123123", Name: "Barina", Age: "brown", DOB: "123", Contact: "1234234345", Owner: "Shotaro"},
	}

	i := 0
	for i < len(NIDs) {
		fmt.Println("i is ", i)
		nidAsBytes, _ := json.Marshal(NIDs[i])
		APIstub.PutState("NID"+strconv.Itoa(i), nidAsBytes)
		fmt.Println("Added", NIDs[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createNID(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	var nid = NID{NIDN: args[1], Name: args[2], Age: args[3], DOB: args[4], Contact: args[5], Owner: args[6]}

	nidAsBytes, _ := json.Marshal(nid)
	APIstub.PutState(args[0], nidAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) authorize(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting at least 3")
	}

	// 0: user, 1: merchant, 2: allowed field
	var auth []string
	i := 2
	n := len(args)
	for i < n {
		auth = append(auth, args[i])
		i++
	}
	key := fmt.Sprintf("auth_%s_%s", args[0], args[1]) 
	nidAsBytes, _ := json.Marshal(auth)
	APIstub.PutState(key, nidAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllNIDs(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "user0"
	endKey := "user999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllNIDs:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changeNIDOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	nidAsBytes, _ := APIstub.GetState(args[0])
	nid := NID{}

	json.Unmarshal(nidAsBytes, &nid)
	nid.Owner = args[1]

	nidAsBytes, _ = json.Marshal(nid)
	APIstub.PutState(args[0], nidAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
