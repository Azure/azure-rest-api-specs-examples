package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/update.json
func ExampleServiceClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcommunication.NewServiceClient("12345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"MyResourceGroup",
		"MyCommunicationResource",
		&armcommunication.ServiceClientUpdateOptions{Parameters: &armcommunication.ServiceResource{
			Tags: map[string]*string{
				"newTag": to.Ptr("newVal"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
