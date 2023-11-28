package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getDscNodeReport.json
func ExampleNodeReportsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNodeReportsClient().Get(ctx, "rg", "myAutomationAccount33", "nodeId", "903a5ead-140c-11e7-a943-000d3a6140c9", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DscNodeReport = armautomation.DscNodeReport{
	// 	Type: to.Ptr("Consistency"),
	// 	ConfigurationVersion: to.Ptr("2.0.0"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:16:27.587Z"); return t}()),
	// 	Errors: []*armautomation.DscReportError{
	// 	},
	// 	HostName: to.Ptr("DSCCOMP"),
	// 	IPV4Addresses: []*string{
	// 		to.Ptr("172.16.2.5"),
	// 		to.Ptr("127.0.0.1")},
	// 		IPV6Addresses: []*string{
	// 			to.Ptr("fe80::4c51:9518:aa3c:256a%5"),
	// 			to.Ptr("::2000:0:0:0"),
	// 			to.Ptr("::1"),
	// 			to.Ptr("::2000:0:0:0"),
	// 			to.Ptr("2001:0:9d38:78cf:106b:130a:53ef:fdfa"),
	// 			to.Ptr("fe80::106b:130a:53ef:fdfa%7")},
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/nodes/nodeId/reports/903a5ead-140c-11e7-a943-000d3a6140c9"),
	// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:16:29.444Z"); return t}()),
	// 			MetaConfiguration: &armautomation.DscMetaConfiguration{
	// 				ActionAfterReboot: to.Ptr("ContinueConfiguration"),
	// 				AllowModuleOverwrite: to.Ptr(false),
	// 				CertificateID: to.Ptr("certId"),
	// 				ConfigurationMode: to.Ptr("ApplyAndMonitor"),
	// 				ConfigurationModeFrequencyMins: to.Ptr[int32](15),
	// 				RebootNodeIfNeeded: to.Ptr(false),
	// 				RefreshFrequencyMins: to.Ptr[int32](30),
	// 			},
	// 			NumberOfResources: to.Ptr[int32](1),
	// 			RebootRequested: to.Ptr("False"),
	// 			RefreshMode: to.Ptr("Pull"),
	// 			ReportFormatVersion: to.Ptr("2.0"),
	// 			ReportID: to.Ptr("903a5ead-140c-11e7-a943-000d3a6140c9"),
	// 			Resources: []*armautomation.DscReportResource{
	// 				{
	// 					DependsOn: []*armautomation.DscReportResourceNavigation{
	// 					},
	// 					DurationInSeconds: to.Ptr[float64](0.25),
	// 					ModuleName: to.Ptr("PsDesiredStateConfiguration"),
	// 					ModuleVersion: to.Ptr("1.1"),
	// 					ResourceID: to.Ptr("[WindowsFeature]IIS"),
	// 					ResourceName: to.Ptr("WindowsFeature"),
	// 					SourceInfo: to.Ptr("::4::32::WindowsFeature"),
	// 					StartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:16:28.181Z"); return t}()),
	// 					Status: to.Ptr("Compliant"),
	// 			}},
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:16:27.587Z"); return t}()),
	// 			Status: to.Ptr("Compliant"),
	// 		}
}
