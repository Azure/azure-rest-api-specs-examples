package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/PrivateLinkResources_Get.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "examplerg", "topics", "exampletopic1", "topic", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResource = armeventgrid.PrivateLinkResource{
	// 	Name: to.Ptr("topic"),
	// 	Type: to.Ptr("Microsoft.EventGrid/topics/privateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/5B4B650E-28B9-4790-B3AB-DDBD88D727C4/resourceGroups/amh/providers/Microsoft.EventGrid/topics/exampletopic1/privateLinkResources/topic"),
	// 	Properties: &armeventgrid.PrivateLinkResourceProperties{
	// 		DisplayName: to.Ptr("Event Grid topic"),
	// 		GroupID: to.Ptr("topic"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("topic")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.eventgrid.azure.net")},
	// 			},
	// 		}
}
