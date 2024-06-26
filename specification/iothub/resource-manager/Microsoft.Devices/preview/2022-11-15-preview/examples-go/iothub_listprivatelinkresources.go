package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/iothub/resource-manager/Microsoft.Devices/preview/2022-11-15-preview/examples/iothub_listprivatelinkresources.json
func ExamplePrivateLinkResourcesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().List(ctx, "myResourceGroup", "testHub", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResources = armiothub.PrivateLinkResources{
	// 	Value: []*armiothub.GroupIDInformation{
	// 		{
	// 			Name: to.Ptr("iotHub"),
	// 			Type: to.Ptr("Microsoft.Devices/IotHubs/PrivateLinkResources"),
	// 			ID: to.Ptr("/subscriptions/91d12660-3dec-467a-be2a-213b5544ddc0/resourceGroups/myResourceGroup/providers/Microsoft.Devices/IotHubs/testHub/PrivateLinkResources/iotHub"),
	// 			Properties: &armiothub.GroupIDInformationProperties{
	// 				GroupID: to.Ptr("iotHub"),
	// 				RequiredMembers: []*string{
	// 					to.Ptr("iotHub")},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.azure-devices.net")},
	// 					},
	// 			}},
	// 		}
}
