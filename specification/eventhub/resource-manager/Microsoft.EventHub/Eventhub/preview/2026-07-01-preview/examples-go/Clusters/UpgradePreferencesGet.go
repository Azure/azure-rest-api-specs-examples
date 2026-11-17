package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub/v2"
)

// Generated from example definition: 2026-07-01-preview/Clusters/UpgradePreferencesGet.json
func ExampleUpgradePreferencesOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUpgradePreferencesOperationsClient().Get(ctx, "contoso-rg", "contoso-cluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armeventhub.UpgradePreferencesOperationsClientGetResponse{
	// 	UpgradePreferences: armeventhub.UpgradePreferences{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Microsoft.EventHub/clusters/contoso-cluster/upgradePreferences/default"),
	// 		Name: to.Ptr("default"),
	// 		Properties: &armeventhub.UpgradePreferencesProperties{
	// 			ExceptionWindows: []*armeventhub.ExceptionWindow{
	// 				{
	// 					Action: to.Ptr(armeventhub.ExceptionWindowActionBlock),
	// 					Date: to.Ptr(time.Date(2026, time.August, 15, 0, 0, 0, 0, time.UTC)),
	// 					DurationMinutes: to.Ptr[int32](1440),
	// 					StartTimeOfDay: to.Ptr("PT0S"),
	// 				},
	// 			},
	// 			MaintenanceWindows: []*armeventhub.MaintenanceWindow{
	// 				{
	// 					DayOfWeek: to.Ptr(armeventhub.UpgradePreferenceDayOfWeekSaturday),
	// 					DurationMinutes: to.Ptr[int32](960),
	// 					StartTimeOfDay: to.Ptr("PT2H"),
	// 				},
	// 			},
	// 			UpgradeStatus: &armeventhub.UpgradeStatus{
	// 				CompletesAt: to.Ptr(time.Date(2026, time.July, 2, 20, 0, 0, 0, time.UTC)),
	// 				InProgress: to.Ptr(true),
	// 				PendingUpgrade: to.Ptr(true),
	// 			},
	// 		},
	// 		Type: to.Ptr("Microsoft.EventHub/clusters/upgradePreferences"),
	// 	},
	// }
}
