package armcontainerservicefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservicefleet/armcontainerservicefleet/v2"
)

// Generated from example definition: 2025-03-01/Fleets_ListCredentialsResult.json
func ExampleFleetsClient_ListCredentials_listsTheUserCredentialsOfAFleet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservicefleet.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFleetsClient().ListCredentials(ctx, "rg1", "fleet", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerservicefleet.FleetsClientListCredentialsResponse{
	// 	FleetCredentialResults: &armcontainerservicefleet.FleetCredentialResults{
	// 		Kubeconfigs: []*armcontainerservicefleet.FleetCredentialResult{
	// 			{
	// 				Name: to.Ptr("credentialName1"),
	// 				Value: []byte("Y3JlZGVudGlhbFZhbHVlMQ=="),
	// 			},
	// 		},
	// 	},
	// }
}
