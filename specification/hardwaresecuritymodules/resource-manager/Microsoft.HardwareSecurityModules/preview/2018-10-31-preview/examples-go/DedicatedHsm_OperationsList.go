package armhardwaresecuritymodules_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules"
)

// x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2018-10-31-preview/examples/DedicatedHsm_OperationsList.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhardwaresecuritymodules.NewOperationsClient(cred, nil)
	_, err = client.List(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
