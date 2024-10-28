package armiotoperations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotoperations/armiotoperations"
)

// Generated from example definition: 2024-11-01/Instance_Get_MaximumSet_Gen.json
func ExampleInstanceClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotoperations.NewClientFactory("F8C729F9-DF9C-4743-848F-96EE433D8E53", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInstanceClient().Get(ctx, "rgiotoperations", "aio-instance", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armiotoperations.InstanceClientGetResponse{
	// 	InstanceResource: &armiotoperations.InstanceResource{
	// 		Properties: &armiotoperations.InstanceProperties{
	// 			SchemaRegistryRef: &armiotoperations.SchemaRegistryRef{
	// 				ResourceID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.DeviceRegistry/schemaRegistries/resource-name123"),
	// 			},
	// 			Description: to.Ptr("rlfvvnnhcypp"),
	// 			ProvisioningState: to.Ptr(armiotoperations.ProvisioningStateSucceeded),
	// 			Version: to.Ptr("vjjbmunthiphfmekvxgxcxkzdwjti"),
	// 		},
	// 		ExtendedLocation: &armiotoperations.ExtendedLocation{
	// 			Name: to.Ptr("qmbrfwcpwwhggszhrdjv"),
	// 			Type: to.Ptr(armiotoperations.ExtendedLocationTypeCustomLocation),
	// 		},
	// 		Identity: &armiotoperations.ManagedServiceIdentity{
	// 			Type: to.Ptr(armiotoperations.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armiotoperations.UserAssignedIdentity{
	// 			},
	// 			PrincipalID: to.Ptr("4a6e4195-75b8-4685-aa0c-0b5704779327"),
	// 			TenantID: to.Ptr("ed060aa2-71ff-4d3f-99c4-a9138356fdec"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("xvewadyhycrjpu"),
	// 		ID: to.Ptr("/subscriptions/0000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup123/providers/Microsoft.IoTOperations/instances/resource-name123"),
	// 		Name: to.Ptr("llptmlifnqqwairx"),
	// 		Type: to.Ptr("qwrfzxjfxvismlqvigot"),
	// 		SystemData: &armiotoperations.SystemData{
	// 			CreatedBy: to.Ptr("ssvaslsmudloholronopqyxjcu"),
	// 			CreatedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("gnicpuszwd"),
	// 			LastModifiedByType: to.Ptr(armiotoperations.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-09T18:13:29.389Z"); return t}()),
	// 		},
	// 	},
	// }
}
