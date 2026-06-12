package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Workspace/InferencePool/list.json
func ExampleInferencePoolsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInferencePoolsClient().NewListPager("test-rg", "my-aml-workspace", &armmachinelearning.InferencePoolsClientListOptions{
		Count:      to.Ptr[int32](1),
		OrderBy:    to.Ptr(armmachinelearning.OrderStringUpdatedAtAsc),
		Properties: to.Ptr("string"),
		Tags:       to.Ptr("string")})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armmachinelearning.InferencePoolsClientListResponse{
		// 	InferencePoolTrackedResourceArmPaginatedResult: armmachinelearning.InferencePoolTrackedResourceArmPaginatedResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.MachineLearningServices/workspaces/my-aml-workspace/inferencePools?api-version=2025-07-01-preview&$skip=2"),
		// 		Value: []*armmachinelearning.InferencePool{
		// 			{
		// 				Name: to.Ptr("string"),
		// 				Type: to.Ptr("string"),
		// 				ID: to.Ptr("string"),
		// 				Identity: &armmachinelearning.ManagedServiceIdentity{
		// 					Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeUserAssigned),
		// 					PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 					TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 					UserAssignedIdentities: map[string]*armmachinelearning.UserAssignedIdentity{
		// 						"string": &armmachinelearning.UserAssignedIdentity{
		// 							ClientID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 							PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 						},
		// 					},
		// 				},
		// 				Kind: to.Ptr("string"),
		// 				Location: to.Ptr("string"),
		// 				Properties: &armmachinelearning.InferencePoolProperties{
		// 					Description: to.Ptr("string"),
		// 					Properties: []*armmachinelearning.StringStringKeyValuePair{
		// 						{
		// 							Key: to.Ptr("string"),
		// 							Value: to.Ptr("string"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armmachinelearning.PoolProvisioningStateFailed),
		// 					ScaleUnitConfiguration: &armmachinelearning.ScaleUnitConfiguration{
		// 						DisablePublicEgress: to.Ptr(false),
		// 						Registries: []*string{
		// 							to.Ptr("string"),
		// 							to.Ptr("registryName1"),
		// 						},
		// 					},
		// 				},
		// 				SKU: &armmachinelearning.SKU{
		// 					Name: to.Ptr("string"),
		// 					Capacity: to.Ptr[int32](1),
		// 					Family: to.Ptr("string"),
		// 					Size: to.Ptr("string"),
		// 					Tier: to.Ptr(armmachinelearning.SKUTierPremium),
		// 				},
		// 				SystemData: &armmachinelearning.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999+00:22"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999+00:22"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeApplication),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
