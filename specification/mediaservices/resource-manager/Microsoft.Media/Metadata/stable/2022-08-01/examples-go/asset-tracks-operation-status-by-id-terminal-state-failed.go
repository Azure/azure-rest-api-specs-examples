package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/asset-tracks-operation-status-by-id-terminal-state-failed.json
func ExampleAssetTrackOperationStatusesClient_Get_getStatusOfAsynchronousOperationWhenItIsCompletedWithError() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssetTrackOperationStatusesClient().Get(ctx, "contoso", "contosomedia", "ClimbingMountRainer", "text1", "86835197-3b47-402e-b313-70b82eaba296", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssetTrackOperationStatus = armmediaservices.AssetTrackOperationStatus{
	// 	Name: to.Ptr("86835197-3b47-402e-b313-70b82eaba296"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-01T00:00:30.000Z"); return t}()),
	// 	Error: &armmediaservices.ErrorDetail{
	// 		Code: to.Ptr("ClientError"),
	// 		Message: to.Ptr("Error while parsing WEBVTT file and creating CMFT header."),
	// 	},
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/assets/ClimbingMountRainer/tracks/text1/operationStatuses/86835197-3b47-402e-b313-70b82eaba296"),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-01T00:00:00.000Z"); return t}()),
	// 	Status: to.Ptr("Failed"),
	// }
}
