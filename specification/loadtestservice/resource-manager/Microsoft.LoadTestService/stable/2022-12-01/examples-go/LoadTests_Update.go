package armloadtesting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/LoadTests_Update.json
func ExampleLoadTestsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armloadtesting.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLoadTestsClient().BeginUpdate(ctx, "dummyrg", "myLoadTest", armloadtesting.LoadTestResourcePatchRequestBody{
		Identity: &armloadtesting.ManagedServiceIdentity{
			Type: to.Ptr(armloadtesting.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
			UserAssignedIdentities: map[string]*armloadtesting.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dummyrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
			},
		},
		Properties: &armloadtesting.LoadTestResourcePatchRequestBodyProperties{
			Description: to.Ptr("This is new load test resource"),
			Encryption: &armloadtesting.EncryptionProperties{
				Identity: &armloadtesting.EncryptionPropertiesIdentity{
					Type: to.Ptr(armloadtesting.TypeSystemAssigned),
				},
				KeyURL: to.Ptr("https://dummy.vault.azure.net/keys/dummykey1"),
			},
		},
		Tags: map[string]*string{
			"Division": to.Ptr("LT"),
			"Team":     to.Ptr("Dev Exp"),
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
	// res.LoadTestResource = armloadtesting.LoadTestResource{
	// 	Name: to.Ptr("myLoadTest"),
	// 	Type: to.Ptr("Microsoft.LoadTestService/loadTests"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dummyrg/providers/Microsoft.LoadTestService/loadTests/myLoadTest"),
	// 	SystemData: &armloadtesting.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-28T12:32:33.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("userId1001"),
	// 		CreatedByType: to.Ptr(armloadtesting.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-28T12:32:33.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("userId1001"),
	// 		LastModifiedByType: to.Ptr(armloadtesting.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"Division": to.Ptr("LT"),
	// 		"Team": to.Ptr("Dev Exp"),
	// 	},
	// 	Identity: &armloadtesting.ManagedServiceIdentity{
	// 		Type: to.Ptr(armloadtesting.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
	// 		PrincipalID: to.Ptr("12345678-1234-1234-1234-123456789012"),
	// 		TenantID: to.Ptr("12345678-1234-1234-4321-123456789012"),
	// 		UserAssignedIdentities: map[string]*armloadtesting.UserAssignedIdentity{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dummyrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": &armloadtesting.UserAssignedIdentity{
	// 				ClientID: to.Ptr("12345678-4321-1234-1234-123456789012"),
	// 				PrincipalID: to.Ptr("12345678-1234-4321-1234-123456789012"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armloadtesting.LoadTestProperties{
	// 		Description: to.Ptr("This is new load test resource"),
	// 		DataPlaneURI: to.Ptr("https://myLoadTest.00000000-0000-0000-0000-000000000000.cnt-dp.domain.com"),
	// 		Encryption: &armloadtesting.EncryptionProperties{
	// 			Identity: &armloadtesting.EncryptionPropertiesIdentity{
	// 				Type: to.Ptr(armloadtesting.TypeSystemAssigned),
	// 			},
	// 			KeyURL: to.Ptr("https://dummy.vault.azure.net/keys/dummykey1"),
	// 		},
	// 		ProvisioningState: to.Ptr(armloadtesting.ResourceStateSucceeded),
	// 	},
	// }
}
