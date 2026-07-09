package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/Pools_Get.json
func ExamplePoolsClient_Get_poolsGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoolsClient().Get(ctx, "rg1", "DevProject", "DevPool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevcenter.PoolsClientGetResponse{
	// 	Pool: armdevcenter.Pool{
	// 		Name: to.Ptr("DevPool"),
	// 		Type: to.Ptr("Microsoft.DevCenter/pools"),
	// 		ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/DevProject/pools/DevPool"),
	// 		Location: to.Ptr("centralus"),
	// 		Properties: &armdevcenter.PoolProperties{
	// 			ActiveHoursConfiguration: &armdevcenter.ActiveHoursConfiguration{
	// 				AutoStartEnableStatus: to.Ptr(armdevcenter.AutoStartEnableStatusEnabled),
	// 				DaysOfWeekLimit: to.Ptr[int32](5),
	// 				DefaultDaysOfWeek: []*armdevcenter.DayOfWeek{
	// 					to.Ptr(armdevcenter.DayOfWeekMonday),
	// 					to.Ptr(armdevcenter.DayOfWeekTuesday),
	// 					to.Ptr(armdevcenter.DayOfWeekWednesday),
	// 					to.Ptr(armdevcenter.DayOfWeekThursday),
	// 					to.Ptr(armdevcenter.DayOfWeekFriday),
	// 				},
	// 				DefaultEndTimeHour: to.Ptr[int32](17),
	// 				DefaultStartTimeHour: to.Ptr[int32](9),
	// 				DefaultTimeZone: to.Ptr("America/Los_Angeles"),
	// 				KeepAwakeEnableStatus: to.Ptr(armdevcenter.KeepAwakeEnableStatusEnabled),
	// 			},
	// 			DevBoxCount: to.Ptr[int32](1),
	// 			DevBoxDefinitionName: to.Ptr("WebDevBox"),
	// 			DisplayName: to.Ptr("Developer Pool"),
	// 			HealthStatus: to.Ptr(armdevcenter.HealthStatusHealthy),
	// 			LicenseType: to.Ptr(armdevcenter.LicenseTypeWindowsClient),
	// 			LocalAdministrator: to.Ptr(armdevcenter.LocalAdminStatusEnabled),
	// 			ManagedVirtualNetworkRegions: []*string{
	// 				to.Ptr("centralus"),
	// 			},
	// 			NetworkConnectionName: to.Ptr("Network1-westus2"),
	// 			ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
	// 			SingleSignOnStatus: to.Ptr(armdevcenter.SingleSignOnStatusDisabled),
	// 			StopOnDisconnect: &armdevcenter.StopOnDisconnectConfiguration{
	// 				GracePeriodMinutes: to.Ptr[int32](60),
	// 				Status: to.Ptr(armdevcenter.StopOnDisconnectEnableStatusEnabled),
	// 			},
	// 			StopOnNoConnect: &armdevcenter.StopOnNoConnectConfiguration{
	// 				GracePeriodMinutes: to.Ptr[int32](120),
	// 				Status: to.Ptr(armdevcenter.StopOnNoConnectEnableStatusEnabled),
	// 			},
	// 			VirtualNetworkType: to.Ptr(armdevcenter.VirtualNetworkTypeManaged),
	// 		},
	// 		SystemData: &armdevcenter.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user1"),
	// 			LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
