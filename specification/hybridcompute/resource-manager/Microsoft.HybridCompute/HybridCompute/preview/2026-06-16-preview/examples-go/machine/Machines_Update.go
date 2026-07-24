package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v3"
)

// Generated from example definition: 2026-06-16-preview/machine/Machines_Update.json
func ExampleMachinesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMachinesClient().Update(ctx, "myResourceGroup", "myMachine", armhybridcompute.MachineUpdate{
		Identity: &armhybridcompute.ManagedServiceIdentity{
			Type: to.Ptr(armhybridcompute.ManagedServiceIdentityTypeSystemAssigned),
		},
		Properties: &armhybridcompute.MachineUpdateProperties{
			IdentityKeyStore: to.Ptr("TPM"),
			LocationData: &armhybridcompute.LocationData{
				Name: to.Ptr("Redmond"),
			},
			OSProfile: &armhybridcompute.OSProfile{
				LinuxConfiguration: &armhybridcompute.OSProfileLinuxConfiguration{
					PatchSettings: &armhybridcompute.PatchSettings{
						AssessmentMode: to.Ptr(armhybridcompute.AssessmentModeTypesImageDefault),
						PatchMode:      to.Ptr(armhybridcompute.PatchModeTypesManual),
					},
				},
				WindowsConfiguration: &armhybridcompute.OSProfileWindowsConfiguration{
					PatchSettings: &armhybridcompute.PatchSettings{
						AssessmentMode:    to.Ptr(armhybridcompute.AssessmentModeTypesImageDefault),
						EnableHotpatching: to.Ptr(true),
						PatchMode:         to.Ptr(armhybridcompute.PatchModeTypesAutomaticByPlatform),
					},
				},
			},
			ParentClusterResourceID:    to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AzureStackHCI/clusters/myAzureStackHCICluster"),
			PrivateLinkScopeResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
			TpmEkCertificate:           to.Ptr("string"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridcompute.MachinesClientUpdateResponse{
	// 	Machine: armhybridcompute.Machine{
	// 		Name: to.Ptr("myMachine"),
	// 		Type: to.Ptr("Microsoft.HybridCompute/machines"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/machines/myMachine"),
	// 		Identity: &armhybridcompute.ManagedServiceIdentity{
	// 			Type: to.Ptr(armhybridcompute.ManagedServiceIdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		},
	// 		Location: to.Ptr("eastus2euap"),
	// 		Properties: &armhybridcompute.MachineProperties{
	// 			ClientPublicKey: to.Ptr("string"),
	// 			IdentityKeyStore: to.Ptr(armhybridcompute.IdentityKeyStoreTPM),
	// 			LicenseProfile: &armhybridcompute.LicenseProfileMachineInstanceView{
	// 				EsuProfile: &armhybridcompute.LicenseProfileMachineInstanceViewEsuProperties{
	// 					EsuEligibility: to.Ptr(armhybridcompute.EsuEligibilityIneligible),
	// 					EsuKeyState: to.Ptr(armhybridcompute.EsuKeyStateInactive),
	// 					EsuKeys: []*armhybridcompute.EsuKey{
	// 						{
	// 							LicenseStatus: to.Ptr[int32](1),
	// 							SKU: to.Ptr("skuNumber1"),
	// 						},
	// 						{
	// 							LicenseStatus: to.Ptr[int32](1),
	// 							SKU: to.Ptr("skuNumber2"),
	// 						},
	// 					},
	// 					LicenseAssignmentState: to.Ptr(armhybridcompute.LicenseAssignmentStateAssigned),
	// 					ServerType: to.Ptr(armhybridcompute.EsuServerTypeStandard),
	// 				},
	// 				LicenseChannel: to.Ptr("PSG"),
	// 				LicenseStatus: to.Ptr(armhybridcompute.LicenseStatusLicensed),
	// 				ProductProfile: &armhybridcompute.LicenseProfileArmProductProfileProperties{
	// 					BillingEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-05T20:36:49.745Z"); return t}()),
	// 					BillingStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T20:36:49.745Z"); return t}()),
	// 					DisenrollmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-05T20:36:49.745Z"); return t}()),
	// 					EnrollmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T20:36:49.745Z"); return t}()),
	// 					ProductFeatures: []*armhybridcompute.ProductFeature{
	// 						{
	// 							Name: to.Ptr("Hotpatch"),
	// 							BillingEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-05T20:36:49.745Z"); return t}()),
	// 							BillingStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T20:36:49.745Z"); return t}()),
	// 							DisenrollmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-05T20:36:49.745Z"); return t}()),
	// 							EnrollmentDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T20:36:49.745Z"); return t}()),
	// 							SubscriptionStatus: to.Ptr(armhybridcompute.LicenseProfileSubscriptionStatusEnabled),
	// 						},
	// 					},
	// 					ProductType: to.Ptr(armhybridcompute.LicenseProfileProductTypeWindowsServer),
	// 					SubscriptionStatus: to.Ptr(armhybridcompute.LicenseProfileSubscriptionStatusEnabled),
	// 				},
	// 				SoftwareAssurance: &armhybridcompute.LicenseProfileMachineInstanceViewSoftwareAssurance{
	// 					SoftwareAssuranceCustomer: to.Ptr(true),
	// 				},
	// 			},
	// 			LocationData: &armhybridcompute.LocationData{
	// 				Name: to.Ptr("Redmond"),
	// 			},
	// 			OSEdition: to.Ptr("Standard"),
	// 			OSProfile: &armhybridcompute.OSProfile{
	// 				LinuxConfiguration: &armhybridcompute.OSProfileLinuxConfiguration{
	// 					PatchSettings: &armhybridcompute.PatchSettings{
	// 						AssessmentMode: to.Ptr(armhybridcompute.AssessmentModeTypesImageDefault),
	// 						PatchMode: to.Ptr(armhybridcompute.PatchModeTypesManual),
	// 					},
	// 				},
	// 				WindowsConfiguration: &armhybridcompute.OSProfileWindowsConfiguration{
	// 					PatchSettings: &armhybridcompute.PatchSettings{
	// 						AssessmentMode: to.Ptr(armhybridcompute.AssessmentModeTypesImageDefault),
	// 						EnableHotpatching: to.Ptr(true),
	// 						PatchMode: to.Ptr(armhybridcompute.PatchModeTypesAutomaticByPlatform),
	// 						Status: &armhybridcompute.PatchSettingsStatus{
	// 							HotpatchEnablementStatus: to.Ptr(armhybridcompute.HotpatchEnablementStatusPendingEvaluation),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ParentClusterResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AzureStackHCI/clusters/myAzureStackHCICluster"),
	// 			PrivateLinkScopeResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			TpmEkCertificate: to.Ptr("string"),
	// 			VMID: to.Ptr("b7a098cc-b0b8-46e8-a205-62f301a62a8f"),
	// 		},
	// 	},
	// }
}
