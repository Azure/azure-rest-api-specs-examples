package armdeviceupdate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80065490402157d0df0dd37ab347c651b22eb576/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/PrivateLinkResources/PrivateLinkResources_Get.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceupdate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "test-rg", "contoso", "adu", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GroupInformation = armdeviceupdate.GroupInformation{
	// 	Name: to.Ptr("adu"),
	// 	Type: to.Ptr("Microsoft.DeviceUpdate/accounts/privateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DeviceUpdate/accounts/contoso/privateLinkResources/adu"),
	// 	Properties: &armdeviceupdate.GroupInformationProperties{
	// 		GroupID: to.Ptr("adu"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("adu")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.adu.microsoft.com")},
	// 			},
	// 		}
}
