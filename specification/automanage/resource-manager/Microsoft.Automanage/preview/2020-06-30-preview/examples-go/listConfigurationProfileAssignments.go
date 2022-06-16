package armautomanage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
)

// x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/preview/2020-06-30-preview/examples/listConfigurationProfileAssignments.json
func ExampleConfigurationProfileAssignmentsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomanage.NewConfigurationProfileAssignmentsClient("<subscription-id>", cred, nil)
	_, err = client.List(ctx,
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
