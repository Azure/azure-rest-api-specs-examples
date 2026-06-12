package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Workspace/ServerlessEndpoint/update.json
func ExampleServerlessEndpointsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerlessEndpointsClient().BeginUpdate(ctx, "test-rg", "my-aml-workspace", "string", armmachinelearning.PartialMinimalTrackedResourceWithSKUAndIdentity{
		Identity: &armmachinelearning.PartialManagedServiceIdentity{
			Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]any{
				"string": map[string]any{},
			},
		},
		SKU: &armmachinelearning.PartialSKU{
			Name:     to.Ptr("string"),
			Capacity: to.Ptr[int32](1),
			Family:   to.Ptr("string"),
			Size:     to.Ptr("string"),
			Tier:     to.Ptr(armmachinelearning.SKUTierPremium),
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.ServerlessEndpointsClientUpdateResponse{
	// 	ServerlessEndpoint: armmachinelearning.ServerlessEndpoint{
	// 		Name: to.Ptr("string"),
	// 		Type: to.Ptr("string"),
	// 		ID: to.Ptr("string"),
	// 		Identity: &armmachinelearning.ManagedServiceIdentity{
	// 			Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			UserAssignedIdentities: map[string]*armmachinelearning.UserAssignedIdentity{
	// 				"string": &armmachinelearning.UserAssignedIdentity{
	// 					ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 					PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 				},
	// 			},
	// 		},
	// 		Kind: to.Ptr("string"),
	// 		Location: to.Ptr("string"),
	// 		Properties: &armmachinelearning.ServerlessEndpointProperties{
	// 			AuthMode: to.Ptr(armmachinelearning.ServerlessInferenceEndpointAuthModeKey),
	// 			ContentSafety: &armmachinelearning.ContentSafety{
	// 				ContentSafetyLevel: to.Ptr(armmachinelearning.ContentSafetyLevelBlocking),
	// 				ContentSafetyStatus: to.Ptr(armmachinelearning.ContentSafetyStatusDisabled),
	// 			},
	// 			EndpointState: to.Ptr(armmachinelearning.ServerlessEndpointStateSuspended),
	// 			InferenceEndpoint: &armmachinelearning.ServerlessInferenceEndpoint{
	// 				Headers: map[string]*string{
	// 					"string": to.Ptr("string"),
	// 				},
	// 				URI: to.Ptr("https://www.contoso.com/example"),
	// 			},
	// 			MarketplaceSubscriptionID: to.Ptr("string"),
	// 			ModelSettings: &armmachinelearning.ModelSettings{
	// 				ModelID: to.Ptr("string"),
	// 			},
	// 			ProvisioningState: to.Ptr(armmachinelearning.EndpointProvisioningStateCreating),
	// 		},
	// 		SKU: &armmachinelearning.SKU{
	// 			Name: to.Ptr("string"),
	// 			Capacity: to.Ptr[int32](1),
	// 			Family: to.Ptr("string"),
	// 			Size: to.Ptr("string"),
	// 			Tier: to.Ptr(armmachinelearning.SKUTierFree),
	// 		},
	// 		SystemData: &armmachinelearning.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999+00:13"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999+00:13"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeApplication),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
