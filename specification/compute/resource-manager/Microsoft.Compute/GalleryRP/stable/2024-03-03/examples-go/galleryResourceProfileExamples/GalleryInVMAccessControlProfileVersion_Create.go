package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d6d0798c6f5eb196fba7bd1924db2b145a94f58c/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryResourceProfileExamples/GalleryInVMAccessControlProfileVersion_Create.json
func ExampleGalleryInVMAccessControlProfileVersionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleryInVMAccessControlProfileVersionsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myGalleryName", "myInVMAccessControlProfileName", "1.0.0", armcompute.GalleryInVMAccessControlProfileVersion{
		Location: to.Ptr("West US"),
		Properties: &armcompute.GalleryInVMAccessControlProfileVersionProperties{
			ExcludeFromLatest: to.Ptr(false),
			TargetLocations: []*armcompute.TargetRegion{
				{
					Name: to.Ptr("West US"),
				},
				{
					Name: to.Ptr("South Central US"),
				}},
			DefaultAccess: to.Ptr(armcompute.EndpointAccessAllow),
			Mode:          to.Ptr(armcompute.AccessControlRulesModeAudit),
			Rules: &armcompute.AccessControlRules{
				Identities: []*armcompute.AccessControlRulesIdentity{
					{
						Name:        to.Ptr("WinPA"),
						ExePath:     to.Ptr("C:\\Windows\\System32\\cscript.exe"),
						GroupName:   to.Ptr("Administrators"),
						ProcessName: to.Ptr("cscript"),
						UserName:    to.Ptr("SYSTEM"),
					}},
				Privileges: []*armcompute.AccessControlRulesPrivilege{
					{
						Name: to.Ptr("GoalState"),
						Path: to.Ptr("/machine"),
						QueryParameters: map[string]*string{
							"comp": to.Ptr("goalstate"),
						},
					}},
				RoleAssignments: []*armcompute.AccessControlRulesRoleAssignment{
					{
						Identities: []*string{
							to.Ptr("WinPA")},
						Role: to.Ptr("Provisioning"),
					}},
				Roles: []*armcompute.AccessControlRulesRole{
					{
						Name: to.Ptr("Provisioning"),
						Privileges: []*string{
							to.Ptr("GoalState")},
					}},
			},
		},
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
	// res.GalleryInVMAccessControlProfileVersion = armcompute.GalleryInVMAccessControlProfileVersion{
	// 	Name: to.Ptr("1.0.0"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGallery/inVMAccessControlProfiles/myInVMAccessControlProfileName/versions/1.0.0"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryInVMAccessControlProfileVersionProperties{
	// 		ExcludeFromLatest: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-01T00:00:00.000Z"); return t}()),
	// 		TargetLocations: []*armcompute.TargetRegion{
	// 			{
	// 				Name: to.Ptr("West US"),
	// 			},
	// 			{
	// 				Name: to.Ptr("South Central US"),
	// 		}},
	// 		DefaultAccess: to.Ptr(armcompute.EndpointAccessAllow),
	// 		Mode: to.Ptr(armcompute.AccessControlRulesModeAudit),
	// 		Rules: &armcompute.AccessControlRules{
	// 			Identities: []*armcompute.AccessControlRulesIdentity{
	// 				{
	// 					Name: to.Ptr("WinPA"),
	// 					ExePath: to.Ptr("C:\\Windows\\System32\\cscript.exe"),
	// 					GroupName: to.Ptr("Administrators"),
	// 					ProcessName: to.Ptr("cscript"),
	// 					UserName: to.Ptr("SYSTEM"),
	// 			}},
	// 			Privileges: []*armcompute.AccessControlRulesPrivilege{
	// 				{
	// 					Name: to.Ptr("GoalState"),
	// 					Path: to.Ptr("/machine"),
	// 					QueryParameters: map[string]*string{
	// 						"comp": to.Ptr("goalstate"),
	// 					},
	// 			}},
	// 			RoleAssignments: []*armcompute.AccessControlRulesRoleAssignment{
	// 				{
	// 					Identities: []*string{
	// 						to.Ptr("WinPA")},
	// 						Role: to.Ptr("Provisioning"),
	// 				}},
	// 				Roles: []*armcompute.AccessControlRulesRole{
	// 					{
	// 						Name: to.Ptr("Provisioning"),
	// 						Privileges: []*string{
	// 							to.Ptr("GoalState")},
	// 					}},
	// 				},
	// 			},
	// 		}
}
