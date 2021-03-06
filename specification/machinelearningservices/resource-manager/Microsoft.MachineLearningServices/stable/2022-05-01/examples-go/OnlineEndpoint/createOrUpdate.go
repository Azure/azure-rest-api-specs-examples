package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/OnlineEndpoint/createOrUpdate.json
func ExampleOnlineEndpointsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewOnlineEndpointsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"test-rg",
		"my-aml-workspace",
		"testEndpointName",
		armmachinelearning.OnlineEndpoint{
			Location: to.Ptr("string"),
			Tags:     map[string]*string{},
			Identity: &armmachinelearning.ManagedServiceIdentity{
				Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeSystemAssigned),
				UserAssignedIdentities: map[string]*armmachinelearning.UserAssignedIdentity{
					"string": {},
				},
			},
			Kind: to.Ptr("string"),
			Properties: &armmachinelearning.OnlineEndpointProperties{
				Description: to.Ptr("string"),
				AuthMode:    to.Ptr(armmachinelearning.EndpointAuthModeAMLToken),
				Properties: map[string]*string{
					"string": to.Ptr("string"),
				},
				Compute: to.Ptr("string"),
				Traffic: map[string]*int32{
					"string": to.Ptr[int32](1),
				},
			},
			SKU: &armmachinelearning.SKU{
				Name:     to.Ptr("string"),
				Capacity: to.Ptr[int32](1),
				Family:   to.Ptr("string"),
				Size:     to.Ptr("string"),
				Tier:     to.Ptr(armmachinelearning.SKUTierFree),
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
