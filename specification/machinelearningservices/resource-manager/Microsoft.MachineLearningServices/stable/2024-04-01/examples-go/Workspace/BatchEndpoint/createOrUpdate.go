package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/BatchEndpoint/createOrUpdate.json
func ExampleBatchEndpointsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBatchEndpointsClient().BeginCreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "testEndpointName", armmachinelearning.BatchEndpoint{
		Location: to.Ptr("string"),
		Tags:     map[string]*string{},
		Identity: &armmachinelearning.ManagedServiceIdentity{
			Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeSystemAssigned),
			UserAssignedIdentities: map[string]*armmachinelearning.UserAssignedIdentity{
				"string": {},
			},
		},
		Kind: to.Ptr("string"),
		Properties: &armmachinelearning.BatchEndpointProperties{
			Description: to.Ptr("string"),
			AuthMode:    to.Ptr(armmachinelearning.EndpointAuthModeAMLToken),
			Properties: map[string]*string{
				"string": to.Ptr("string"),
			},
			Defaults: &armmachinelearning.BatchEndpointDefaults{
				DeploymentName: to.Ptr("string"),
			},
		},
		SKU: &armmachinelearning.SKU{
			Name:     to.Ptr("string"),
			Capacity: to.Ptr[int32](1),
			Family:   to.Ptr("string"),
			Size:     to.Ptr("string"),
			Tier:     to.Ptr(armmachinelearning.SKUTierFree),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BatchEndpoint = armmachinelearning.BatchEndpoint{
	// 	Name: to.Ptr("string"),
	// 	Type: to.Ptr("string"),
	// 	ID: to.Ptr("string"),
	// 	SystemData: &armmachinelearning.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("string"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armmachinelearning.ManagedServiceIdentity{
	// 		Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		UserAssignedIdentities: map[string]*armmachinelearning.UserAssignedIdentity{
	// 			"string": &armmachinelearning.UserAssignedIdentity{
	// 				ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 				PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			},
	// 		},
	// 	},
	// 	Kind: to.Ptr("string"),
	// 	Properties: &armmachinelearning.BatchEndpointProperties{
	// 		Description: to.Ptr("string"),
	// 		AuthMode: to.Ptr(armmachinelearning.EndpointAuthModeAMLToken),
	// 		Properties: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		ScoringURI: to.Ptr("https://www.contoso.com/example"),
	// 		SwaggerURI: to.Ptr("https://www.contoso.com/example"),
	// 		Defaults: &armmachinelearning.BatchEndpointDefaults{
	// 			DeploymentName: to.Ptr("string"),
	// 		},
	// 		ProvisioningState: to.Ptr(armmachinelearning.EndpointProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armmachinelearning.SKU{
	// 		Name: to.Ptr("string"),
	// 		Capacity: to.Ptr[int32](1),
	// 		Family: to.Ptr("string"),
	// 		Size: to.Ptr("string"),
	// 		Tier: to.Ptr(armmachinelearning.SKUTierFree),
	// 	},
	// }
}
