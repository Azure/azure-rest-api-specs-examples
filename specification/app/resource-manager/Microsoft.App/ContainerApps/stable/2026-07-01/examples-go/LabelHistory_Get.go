package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2026-07-01/LabelHistory_Get.json
func ExampleContainerAppsLabelHistoryClient_GetLabelHistory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsLabelHistoryClient().GetLabelHistory(ctx, "rg", "testContainerApp", "dev", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappcontainers.ContainerAppsLabelHistoryClientGetLabelHistoryResponse{
	// 	LabelHistory: armappcontainers.LabelHistory{
	// 		Name: to.Ptr("dev"),
	// 		Type: to.Ptr("Microsoft.App/containerApps/labelHistories"),
	// 		ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/testContainerApp/labelHistories/dev"),
	// 		Properties: &armappcontainers.LabelHistoryProperties{
	// 			Records: []*armappcontainers.LabelHistoryRecordItem{
	// 				{
	// 					Revision: to.Ptr("testContainerApp--2"),
	// 					Start: to.Ptr(time.Date(2024, time.October, 15, 5, 49, 47, 0, time.UTC)),
	// 					Status: to.Ptr(armappcontainers.StatusSucceeded),
	// 				},
	// 				{
	// 					Revision: to.Ptr("testContainerApp--1"),
	// 					Start: to.Ptr(time.Date(2024, time.October, 15, 1, 12, 56, 0, time.UTC)),
	// 					Status: to.Ptr(armappcontainers.StatusSucceeded),
	// 					Stop: to.Ptr(time.Date(2024, time.October, 15, 5, 45, 8, 0, time.UTC)),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
