package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-12-03/galleryResourceProfileExamples/GalleryInVMAccessControlProfileVersion_Update.json
func ExampleGalleryInVMAccessControlProfileVersionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleryInVMAccessControlProfileVersionsClient().BeginUpdate(ctx, "myResourceGroup", "myGalleryName", "myInVMAccessControlProfileName", "1.0.0", armcompute.GalleryInVMAccessControlProfileVersionUpdate{
		Properties: &armcompute.GalleryInVMAccessControlProfileVersionProperties{
			Mode:          to.Ptr(armcompute.AccessControlRulesModeAudit),
			DefaultAccess: to.Ptr(armcompute.EndpointAccessAllow),
			TargetLocations: []*armcompute.TargetRegion{
				{
					Name: to.Ptr("West US"),
				},
				{
					Name: to.Ptr("South Central US"),
				},
				{
					Name: to.Ptr("East US"),
				},
			},
			ExcludeFromLatest: to.Ptr(false),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.GalleryInVMAccessControlProfileVersionsClientUpdateResponse{
	// 	GalleryInVMAccessControlProfileVersion: armcompute.GalleryInVMAccessControlProfileVersion{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGallery/inVMAccessControlProfiles/myInVMAccessControlProfileName/versions/1.0.0"),
	// 		Properties: &armcompute.GalleryInVMAccessControlProfileVersionProperties{
	// 			Mode: to.Ptr(armcompute.AccessControlRulesModeAudit),
	// 			DefaultAccess: to.Ptr(armcompute.EndpointAccessAllow),
	// 			Rules: &armcompute.AccessControlRules{
	// 				Privileges: []*armcompute.AccessControlRulesPrivilege{
	// 					{
	// 						Name: to.Ptr("GoalState"),
	// 						Path: to.Ptr("/machine"),
	// 						QueryParameters: map[string]*string{
	// 							"comp": to.Ptr("goalstate"),
	// 						},
	// 					},
	// 				},
	// 				Roles: []*armcompute.AccessControlRulesRole{
	// 					{
	// 						Name: to.Ptr("Provisioning"),
	// 						Privileges: []*string{
	// 							to.Ptr("GoalState"),
	// 						},
	// 					},
	// 				},
	// 				Identities: []*armcompute.AccessControlRulesIdentity{
	// 					{
	// 						Name: to.Ptr("WinPA"),
	// 						UserName: to.Ptr("SYSTEM"),
	// 						GroupName: to.Ptr("Administrators"),
	// 						ExePath: to.Ptr("C:\\Windows\\System32\\cscript.exe"),
	// 						ProcessName: to.Ptr("cscript"),
	// 					},
	// 				},
	// 				RoleAssignments: []*armcompute.AccessControlRulesRoleAssignment{
	// 					{
	// 						Role: to.Ptr("Provisioning"),
	// 						Identities: []*string{
	// 							to.Ptr("WinPA"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			TargetLocations: []*armcompute.TargetRegion{
	// 				{
	// 					Name: to.Ptr("West US"),
	// 				},
	// 				{
	// 					Name: to.Ptr("South Central US"),
	// 				},
	// 				{
	// 					Name: to.Ptr("East US"),
	// 				},
	// 			},
	// 			ExcludeFromLatest: to.Ptr(false),
	// 			PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00Z"); return t}()),
	// 			ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateUpdating),
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		Name: to.Ptr("1.0.0"),
	// 	},
	// }
}
