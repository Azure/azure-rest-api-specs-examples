package armappcomplianceautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82ea406b73d671269217053d7ef336450d860345/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Report_Get.json
func ExampleReportClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReportClient().Get(ctx, "testReport", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReportResource = armappcomplianceautomation.ReportResource{
	// 	Name: to.Ptr("testReportName"),
	// 	Type: to.Ptr("Microsfot.AppComplianceAutomation/reports"),
	// 	ID: to.Ptr("/provider/Microsfot.AppComplianceAutomation/reports/testReportName"),
	// 	SystemData: &armappcomplianceautomation.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
	// 		CreatedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		CreatedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		LastModifiedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
	// 	},
	// 	Properties: &armappcomplianceautomation.ReportProperties{
	// 		ComplianceStatus: &armappcomplianceautomation.ReportComplianceStatus{
	// 			M365: &armappcomplianceautomation.OverviewStatus{
	// 				FailedCount: to.Ptr[int32](0),
	// 				ManualCount: to.Ptr[int32](0),
	// 				PassedCount: to.Ptr[int32](0),
	// 			},
	// 		},
	// 		ID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		LastTriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:17:23.922Z"); return t}()),
	// 		NextTriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:17:23.922Z"); return t}()),
	// 		OfferGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		ProvisioningState: to.Ptr(armappcomplianceautomation.ProvisioningStateSucceeded),
	// 		ReportName: to.Ptr("testReportName"),
	// 		Resources: []*armappcomplianceautomation.ResourceMetadata{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService"),
	// 				Tags: map[string]*string{
	// 					"key1": to.Ptr("value1"),
	// 				},
	// 		}},
	// 		Status: to.Ptr(armappcomplianceautomation.ReportStatusActive),
	// 		Subscriptions: []*string{
	// 			to.Ptr("00000000-0000-0000-0000-000000000000")},
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TimeZone: to.Ptr("GMT Standard Time"),
	// 			TriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-02T05:17:23.922Z"); return t}()),
	// 		},
	// 	}
}
