package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase/v3"
)

// Generated from example definition: 2026-06-01/GoldenGateDeployments_ListBySubscription_MaximumSet_Gen.json
func ExampleGoldenGateDeploymentsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGoldenGateDeploymentsClient().NewListBySubscriptionPager(nil)
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
		// page = armoracledatabase.GoldenGateDeploymentsClientListBySubscriptionResponse{
		// 	GoldenGateDeploymentListResult: armoracledatabase.GoldenGateDeploymentListResult{
		// 		Value: []*armoracledatabase.GoldenGateDeployment{
		// 			{
		// 				Properties: &armoracledatabase.DeploymentProperties{
		// 					BackupSchedule: &armoracledatabase.BackupScheduleType{
		// 						BucketName: to.Ptr("resource1"),
		// 						CompartmentID: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 						FrequencyBackupScheduled: to.Ptr(armoracledatabase.FrequencyTypeDaily),
		// 						IsMetadataOnly: to.Ptr(true),
		// 						NamespaceName: to.Ptr("qif"),
		// 						TimeBackupScheduled: to.Ptr("2026-06-01T00:00:00Z"),
		// 					},
		// 					Compartment: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 					Ocid: to.Ptr("ocid1.autonomousdatabase.oc1..aaaaa3klq"),
		// 					CPUCoreCount: to.Ptr[int32](18),
		// 					DisplayName: to.Ptr("vo"),
		// 					Category: to.Ptr(armoracledatabase.CategoryTypeDataReplication),
		// 					DeploymentType: to.Ptr(armoracledatabase.DeploymentTypeOgg),
		// 					DeploymentURL: to.Ptr("https://microsoft.com/a"),
		// 					EnvironmentType: to.Ptr(armoracledatabase.SetupTypeProduction),
		// 					IngressIPs: []*string{
		// 						to.Ptr("example"),
		// 					},
		// 					LicenseModel: to.Ptr(armoracledatabase.LicenseModelLicenseIncluded),
		// 					LifecycleDetails: to.Ptr("example"),
		// 					LifecycleState: to.Ptr(armoracledatabase.DeploymentLifecycleStateCreating),
		// 					TimeCreated: to.Ptr("2026-06-01T00:00:00Z"),
		// 					TimeUpdated: to.Ptr("2026-06-01T00:00:00Z"),
		// 					MaintenanceConfiguration: &armoracledatabase.MaintenanceConfigurationType{
		// 						BundleReleaseUpgradePeriodInDays: to.Ptr[int32](21),
		// 						InterimReleaseUpgradePeriodInDays: to.Ptr[int32](28),
		// 						IsInterimReleaseAutoUpgradeEnabled: to.Ptr(true),
		// 						MajorReleaseUpgradePeriodInDays: to.Ptr[int32](5),
		// 						SecurityPatchUpgradePeriodInDays: to.Ptr[int32](10),
		// 					},
		// 					MaintenanceWindow: &armoracledatabase.MaintenanceWindowType{
		// 						Day: to.Ptr(armoracledatabase.DayOfWeekName("resource1")),
		// 						StartHour: to.Ptr[int32](25),
		// 					},
		// 					NetworkAnchorID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/networkAnchors/networkanchor1"),
		// 					OggData: &armoracledatabase.OggDeploymentDetails{
		// 						AdminUsername: to.Ptr("resource1"),
		// 						Certificate: to.Ptr("example"),
		// 						CredentialStore: to.Ptr(armoracledatabase.CredentialTypeGoldenGate),
		// 						DeploymentName: to.Ptr("uzyxba"),
		// 						GroupToRolesMapping: &armoracledatabase.GroupToRolesMappingDetails{
		// 							AdministratorGroupID: to.Ptr("example"),
		// 							OperatorGroupID: to.Ptr("example"),
		// 							SecurityGroupID: to.Ptr("example"),
		// 							UserGroupID: to.Ptr("example"),
		// 							IdentityDomainID: to.Ptr("example"),
		// 						},
		// 						OggVersion: to.Ptr("example"),
		// 						PasswordSecretID: to.Ptr("example"),
		// 					},
		// 					PrivateIPAddress: to.Ptr("wdzfjc"),
		// 					ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
		// 					ResourceAnchorID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/resourceAnchors/anchor1"),
		// 					StorageUtilizationInBytes: to.Ptr[int32](29),
		// 					TimeZone: to.Ptr("2026-06-01T00:00:00Z"),
		// 					Version: to.Ptr("zhvgen"),
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("example"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key4445": to.Ptr("zasewagrwgc"),
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/resources/resource1"),
		// 				Name: to.Ptr("resource1"),
		// 				Type: to.Ptr("dfxma"),
		// 				SystemData: &armoracledatabase.SystemData{
		// 					CreatedBy: to.Ptr("ns"),
		// 					CreatedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2026, time.July, 28, 21, 47, 8, 141000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("example"),
		// 					LastModifiedByType: to.Ptr(armoracledatabase.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2026, time.July, 28, 21, 47, 8, 141000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
