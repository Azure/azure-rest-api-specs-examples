package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2022-01-31-preview/examples/iotconnectors/iotconnector_Patch.json
func ExampleIotConnectorsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhealthcareapis.NewIotConnectorsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"testRG",
		"blue",
		"workspace1",
		armhealthcareapis.IotConnectorPatchResource{
			Tags: map[string]*string{
				"additionalProp1": to.Ptr("string"),
				"additionalProp2": to.Ptr("string"),
				"additionalProp3": to.Ptr("string"),
			},
			Identity: &armhealthcareapis.ServiceManagedIdentityIdentity{
				Type: to.Ptr(armhealthcareapis.ServiceManagedIdentityTypeSystemAssigned),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
