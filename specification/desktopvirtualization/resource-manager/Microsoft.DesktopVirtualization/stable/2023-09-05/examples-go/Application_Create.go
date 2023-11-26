package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/Application_Create.json
func ExampleApplicationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationsClient().CreateOrUpdate(ctx, "resourceGroup1", "applicationGroup1", "application1", armdesktopvirtualization.Application{
		Properties: &armdesktopvirtualization.ApplicationProperties{
			Description:          to.Ptr("des1"),
			CommandLineArguments: to.Ptr("arguments"),
			CommandLineSetting:   to.Ptr(armdesktopvirtualization.CommandLineSettingAllow),
			FilePath:             to.Ptr("path"),
			FriendlyName:         to.Ptr("friendly"),
			IconIndex:            to.Ptr[int32](1),
			IconPath:             to.Ptr("icon"),
			ShowInPortal:         to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Application = armdesktopvirtualization.Application{
	// 	Name: to.Ptr("applicationGroup1/application1"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/applicationGroups/applications"),
	// 	ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/applicationGroups/applicationGroup1/applications/application1"),
	// 	Properties: &armdesktopvirtualization.ApplicationProperties{
	// 		Description: to.Ptr("des1"),
	// 		ApplicationType: to.Ptr(armdesktopvirtualization.RemoteApplicationTypeInBuilt),
	// 		CommandLineArguments: to.Ptr("arguments"),
	// 		CommandLineSetting: to.Ptr(armdesktopvirtualization.CommandLineSettingAllow),
	// 		FilePath: to.Ptr("path"),
	// 		FriendlyName: to.Ptr("friendly"),
	// 		IconContent: []byte("VGhpcyBpcyBhIHN0cmluZyB0byBoYXNo"),
	// 		IconHash: to.Ptr("bEQ1n2HysrGxCDvgt4bfOtkET8ydzh5SXGM0KGTBPVc"),
	// 		IconIndex: to.Ptr[int32](1),
	// 		IconPath: to.Ptr("icon"),
	// 		ObjectID: to.Ptr("7877fb31-4bde-49fd-9df3-c046e0ec5325"),
	// 		ShowInPortal: to.Ptr(true),
	// 	},
	// 	SystemData: &armdesktopvirtualization.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 	},
	// }
}
