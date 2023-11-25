package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/PackageUpdate.json
func ExamplePackagesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtestbase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPackagesClient().BeginUpdate(ctx, "contoso-rg1", "contoso-testBaseAccount1", "contoso-package2", armtestbase.PackageUpdateParameters{
		Properties: &armtestbase.PackageUpdateParameterProperties{
			BlobPath:      to.Ptr("storageAccountPath/package.zip"),
			FlightingRing: to.Ptr("Insider Beta Channel"),
			IsEnabled:     to.Ptr(false),
			TargetOSList: []*armtestbase.TargetOSInfo{
				{
					OSUpdateType: to.Ptr("Security updates"),
					TargetOSs: []*string{
						to.Ptr("Windows 10 2004"),
						to.Ptr("Windows 10 1903")},
				}},
			Tests: []*armtestbase.Test{
				{
					IsActive: to.Ptr(true),
					TestType: to.Ptr(armtestbase.TestTypeOutOfBoxTest),
					Commands: []*armtestbase.Command{
						{
							Name:              to.Ptr("Install"),
							Action:            to.Ptr(armtestbase.ActionInstall),
							AlwaysRun:         to.Ptr(true),
							ApplyUpdateBefore: to.Ptr(false),
							Content:           to.Ptr("app/scripts/install/job.ps1"),
							ContentType:       to.Ptr(armtestbase.ContentTypePath),
							MaxRunTime:        to.Ptr[int32](1800),
							RestartAfter:      to.Ptr(true),
							RunAsInteractive:  to.Ptr(true),
							RunElevated:       to.Ptr(true),
						},
						{
							Name:              to.Ptr("Launch"),
							Action:            to.Ptr(armtestbase.ActionLaunch),
							AlwaysRun:         to.Ptr(false),
							ApplyUpdateBefore: to.Ptr(true),
							Content:           to.Ptr("app/scripts/launch/job.ps1"),
							ContentType:       to.Ptr(armtestbase.ContentTypePath),
							MaxRunTime:        to.Ptr[int32](1800),
							RestartAfter:      to.Ptr(false),
							RunAsInteractive:  to.Ptr(true),
							RunElevated:       to.Ptr(true),
						},
						{
							Name:              to.Ptr("Close"),
							Action:            to.Ptr(armtestbase.ActionClose),
							AlwaysRun:         to.Ptr(false),
							ApplyUpdateBefore: to.Ptr(false),
							Content:           to.Ptr("app/scripts/close/job.ps1"),
							ContentType:       to.Ptr(armtestbase.ContentTypePath),
							MaxRunTime:        to.Ptr[int32](1800),
							RestartAfter:      to.Ptr(false),
							RunAsInteractive:  to.Ptr(true),
							RunElevated:       to.Ptr(true),
						},
						{
							Name:              to.Ptr("Uninstall"),
							Action:            to.Ptr(armtestbase.ActionUninstall),
							AlwaysRun:         to.Ptr(true),
							ApplyUpdateBefore: to.Ptr(false),
							Content:           to.Ptr("app/scripts/uninstall/job.ps1"),
							ContentType:       to.Ptr(armtestbase.ContentTypePath),
							MaxRunTime:        to.Ptr[int32](1800),
							RestartAfter:      to.Ptr(false),
							RunAsInteractive:  to.Ptr(true),
							RunElevated:       to.Ptr(true),
						}},
				}},
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PackageResource = armtestbase.PackageResource{
	// 	Name: to.Ptr("contoso-package2"),
	// 	Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/packages"),
	// 	ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/packages/contoso-package2"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armtestbase.PackageProperties{
	// 		ApplicationName: to.Ptr("contoso-package2"),
	// 		BlobPath: to.Ptr("storageAccountPath/package.zip"),
	// 		FlightingRing: to.Ptr("Insider Beta Channel"),
	// 		IsEnabled: to.Ptr(false),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-28T17:30:00.000Z"); return t}()),
	// 		PackageStatus: to.Ptr(armtestbase.PackageStatusReady),
	// 		ProvisioningState: to.Ptr(armtestbase.ProvisioningStateSucceeded),
	// 		TargetOSList: []*armtestbase.TargetOSInfo{
	// 			{
	// 				OSUpdateType: to.Ptr("Security updates"),
	// 				TargetOSs: []*string{
	// 					to.Ptr("Windows 10 2004"),
	// 					to.Ptr("Windows 10 1903")},
	// 			}},
	// 			TestTypes: []*armtestbase.TestType{
	// 				to.Ptr(armtestbase.TestTypeOutOfBoxTest)},
	// 				Tests: []*armtestbase.Test{
	// 					{
	// 						IsActive: to.Ptr(true),
	// 						TestType: to.Ptr(armtestbase.TestTypeOutOfBoxTest),
	// 						ValidationRunStatus: to.Ptr(armtestbase.ValidationRunStatusPassed),
	// 						Commands: []*armtestbase.Command{
	// 							{
	// 								Name: to.Ptr("Install"),
	// 								Action: to.Ptr(armtestbase.ActionInstall),
	// 								AlwaysRun: to.Ptr(true),
	// 								ApplyUpdateBefore: to.Ptr(false),
	// 								Content: to.Ptr("app/scripts/install/job.ps1"),
	// 								ContentType: to.Ptr(armtestbase.ContentTypePath),
	// 								MaxRunTime: to.Ptr[int32](1800),
	// 								RestartAfter: to.Ptr(true),
	// 								RunAsInteractive: to.Ptr(true),
	// 								RunElevated: to.Ptr(true),
	// 							},
	// 							{
	// 								Name: to.Ptr("Launch"),
	// 								Action: to.Ptr(armtestbase.ActionLaunch),
	// 								AlwaysRun: to.Ptr(false),
	// 								ApplyUpdateBefore: to.Ptr(true),
	// 								Content: to.Ptr("app/scripts/launch/job.ps1"),
	// 								ContentType: to.Ptr(armtestbase.ContentTypePath),
	// 								MaxRunTime: to.Ptr[int32](1800),
	// 								RestartAfter: to.Ptr(false),
	// 								RunAsInteractive: to.Ptr(true),
	// 								RunElevated: to.Ptr(true),
	// 							},
	// 							{
	// 								Name: to.Ptr("Close"),
	// 								Action: to.Ptr(armtestbase.ActionClose),
	// 								AlwaysRun: to.Ptr(false),
	// 								ApplyUpdateBefore: to.Ptr(false),
	// 								Content: to.Ptr("app/scripts/close/job.ps1"),
	// 								ContentType: to.Ptr(armtestbase.ContentTypePath),
	// 								MaxRunTime: to.Ptr[int32](1800),
	// 								RestartAfter: to.Ptr(false),
	// 								RunAsInteractive: to.Ptr(true),
	// 								RunElevated: to.Ptr(true),
	// 							},
	// 							{
	// 								Name: to.Ptr("Uninstall"),
	// 								Action: to.Ptr(armtestbase.ActionUninstall),
	// 								AlwaysRun: to.Ptr(true),
	// 								ApplyUpdateBefore: to.Ptr(false),
	// 								Content: to.Ptr("app/scripts/uninstall/job.ps1"),
	// 								ContentType: to.Ptr(armtestbase.ContentTypePath),
	// 								MaxRunTime: to.Ptr[int32](1800),
	// 								RestartAfter: to.Ptr(false),
	// 								RunAsInteractive: to.Ptr(true),
	// 								RunElevated: to.Ptr(true),
	// 						}},
	// 				}},
	// 				ValidationResults: []*armtestbase.PackageValidationResult{
	// 					{
	// 						IsValid: to.Ptr(true),
	// 						ValidationName: to.Ptr("Syntax Validation"),
	// 					},
	// 					{
	// 						IsValid: to.Ptr(true),
	// 						ValidationName: to.Ptr("Package Run Validation"),
	// 				}},
	// 				Version: to.Ptr("1.0.0"),
	// 			},
	// 		}
}
