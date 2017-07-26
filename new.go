func (t *SimpleChainecode) changeOwner(stub shim.ChaincodeStubInterface, args []string)([]byte,error){
	var asset, newowner string
	var err error
	fmt.Println("running write()")
	
	if len(args) != 2 {
		return nil, error.New("Incorrect number of arguments.")
	}
	
	asset = args[0]
	newowner = arg[1]
	err = stub.PutState(asset, []byte(newowner))
	if err != nill {
		return nil, err
	}
	return nil, err
}