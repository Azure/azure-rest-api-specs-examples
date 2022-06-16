package armazurearcdata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-11-01/examples/GetDataController.json
func ExampleDataControllersClient_GetDataController() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armazurearcdata.NewDataControllersClient("<subscription-id>", cred, nil)
	res, err := client.GetDataController(ctx,
		"<resource-group-name>",
		"<data-controller-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DataControllersClientGetDataControllerResult)
}
