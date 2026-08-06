package armcontainerserviceaimanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerserviceaimanager/armcontainerserviceaimanager"
)

// Generated from example definition: 2026-05-02-preview/AIManagers_CreateOrUpdate.json
func ExampleAIManagersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerserviceaimanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAIManagersClient().BeginCreateOrUpdate(ctx, "rg1", "aimanager1", armcontainerserviceaimanager.AIManager{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Identity: &armcontainerserviceaimanager.ManagedServiceIdentity{
			Type: to.Ptr(armcontainerserviceaimanager.ManagedServiceIdentityTypeSystemAssigned),
		},
		Properties: &armcontainerserviceaimanager.AIManagerProperties{
			DeletePolicy: to.Ptr(armcontainerserviceaimanager.DeletePolicyKeep),
		},
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
	// res = armcontainerserviceaimanager.AIManagersClientCreateOrUpdateResponse{
	// 	AIManager: armcontainerserviceaimanager.AIManager{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/aiManagers/aimanager1"),
	// 		Name: to.Ptr("aimanager1"),
	// 		Type: to.Ptr("Microsoft.ContainerService/aiManagers"),
	// 		Location: to.Ptr("eastus"),
	// 		Tags: map[string]*string{
	// 			"key1": to.Ptr("value1"),
	// 		},
	// 		SystemData: &armcontainerserviceaimanager.SystemData{
	// 			CreatedBy: to.Ptr("user@example.com"),
	// 			CreatedByType: to.Ptr(armcontainerserviceaimanager.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00.000Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@example.com"),
	// 			LastModifiedByType: to.Ptr(armcontainerserviceaimanager.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00.000Z"); return t}()),
	// 		},
	// 		ETag: to.Ptr("\"00000000-0000-0000-0000-000000000000\""),
	// 		Identity: &armcontainerserviceaimanager.ManagedServiceIdentity{
	// 			Type: to.Ptr(armcontainerserviceaimanager.ManagedServiceIdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 		Properties: &armcontainerserviceaimanager.AIManagerProperties{
	// 			ProvisioningState: to.Ptr(armcontainerserviceaimanager.AIManagerProvisioningStateSucceeded),
	// 			DeletePolicy: to.Ptr(armcontainerserviceaimanager.DeletePolicyKeep),
	// 			ManagedResourceGroupName: to.Ptr("AIM_rg1_aimanager1_eastus"),
	// 		},
	// 	},
	// }
}
