package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"
)

// x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-03-01-preview/examples/CheckNameAvailable.json
func ExampleOperationsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappconfiguration.NewOperationsClient("<subscription-id>", cred, nil)
	res, err := client.CheckNameAvailability(ctx,
		armappconfiguration.CheckNameAvailabilityParameters{
			Name: to.StringPtr("<name>"),
			Type: armappconfiguration.ConfigurationResourceType("Microsoft.AppConfiguration/configurationStores").ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.OperationsClientCheckNameAvailabilityResult)
}
