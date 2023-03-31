package armdynatrace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3751f321654db00858e70649291af5c81e94611e/specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/Monitors_ListAppServices_MaximumSet_Gen.json
func ExampleMonitorsClient_NewListAppServicesPager_monitorsListAppServicesMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdynatrace.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListAppServicesPager("myResourceGroup", "myMonitor", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AppServiceListResponse = armdynatrace.AppServiceListResponse{
		// 	Value: []*armdynatrace.AppServiceInfo{
		// 		{
		// 			AutoUpdateSetting: to.Ptr(armdynatrace.AutoUpdateSettingENABLED),
		// 			AvailabilityState: to.Ptr(armdynatrace.AvailabilityStateCRASHED),
		// 			HostGroup: to.Ptr("myGroup"),
		// 			HostName: to.Ptr("myName"),
		// 			LogModule: to.Ptr(armdynatrace.LogModuleENABLED),
		// 			MonitoringType: to.Ptr(armdynatrace.MonitoringTypeCLOUDINFRASTRUCTURE),
		// 			ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/virtual1"),
		// 			UpdateStatus: to.Ptr(armdynatrace.UpdateStatusINCOMPATIBLE),
		// 			Version: to.Ptr("1.2.0"),
		// 	}},
		// }
	}
}
