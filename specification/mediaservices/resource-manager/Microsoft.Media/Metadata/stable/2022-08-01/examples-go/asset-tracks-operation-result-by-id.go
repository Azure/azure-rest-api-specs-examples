package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/asset-tracks-operation-result-by-id.json
func ExampleAssetTrackOperationResultsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssetTrackOperationResultsClient().Get(ctx, "contoso", "contosomedia", "ClimbingMountRainer", "text1", "e78f8d40-7aaa-4f2f-8ae6-73987e7c5a08", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssetTrack = armmediaservices.AssetTrack{
	// 	Name: to.Ptr("text1"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/assets/tracks/operationResults"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/assets/ClimbingMountRainer/tracks/text1/operationResults/e78f8d40-7aaa-4f2f-8ae6-73987e7c5a08"),
	// 	Properties: &armmediaservices.AssetTrackProperties{
	// 		ProvisioningState: to.Ptr(armmediaservices.ProvisioningStateSucceeded),
	// 		Track: &armmediaservices.TextTrack{
	// 			ODataType: to.Ptr("#Microsoft.Media.TextTrack"),
	// 			DisplayName: to.Ptr("Auto generated"),
	// 			LanguageCode: to.Ptr("en-us"),
	// 			PlayerVisibility: to.Ptr(armmediaservices.VisibilityVisible),
	// 		},
	// 	},
	// }
}
