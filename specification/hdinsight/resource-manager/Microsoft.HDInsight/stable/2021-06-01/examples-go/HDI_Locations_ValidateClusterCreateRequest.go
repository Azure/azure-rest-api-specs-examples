package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/HDI_Locations_ValidateClusterCreateRequest.json
func ExampleLocationsClient_ValidateClusterCreateRequest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhdinsight.NewLocationsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ValidateClusterCreateRequest(ctx,
		"southcentralus",
		armhdinsight.ClusterCreateRequestValidationParameters{
			Location: to.Ptr("southcentralus"),
			Properties: &armhdinsight.ClusterCreateProperties{
				ClusterDefinition: &armhdinsight.ClusterDefinition{
					ComponentVersion: map[string]*string{
						"Spark": to.Ptr("2.4"),
					},
					Configurations: map[string]interface{}{
						"gateway": map[string]interface{}{
							"restAuthCredential.isEnabled": true,
							"restAuthCredential.password":  "**********",
							"restAuthCredential.username":  "admin",
						},
					},
					Kind: to.Ptr("spark"),
				},
				ClusterVersion: to.Ptr("4.0"),
				ComputeProfile: &armhdinsight.ComputeProfile{
					Roles: []*armhdinsight.Role{
						{
							Name: to.Ptr("headnode"),
							HardwareProfile: &armhdinsight.HardwareProfile{
								VMSize: to.Ptr("Standard_E8_V3"),
							},
							MinInstanceCount: to.Ptr[int32](1),
							OSProfile: &armhdinsight.OsProfile{
								LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
									Password: to.Ptr("********"),
									Username: to.Ptr("sshuser"),
								},
							},
							ScriptActions:       []*armhdinsight.ScriptAction{},
							TargetInstanceCount: to.Ptr[int32](2),
						},
						{
							Name: to.Ptr("workernode"),
							HardwareProfile: &armhdinsight.HardwareProfile{
								VMSize: to.Ptr("Standard_E8_V3"),
							},
							OSProfile: &armhdinsight.OsProfile{
								LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
									Password: to.Ptr("********"),
									Username: to.Ptr("sshuser"),
								},
							},
							ScriptActions:       []*armhdinsight.ScriptAction{},
							TargetInstanceCount: to.Ptr[int32](4),
						},
						{
							Name: to.Ptr("zookeepernode"),
							HardwareProfile: &armhdinsight.HardwareProfile{
								VMSize: to.Ptr("Standard_D13_V2"),
							},
							MinInstanceCount: to.Ptr[int32](1),
							OSProfile: &armhdinsight.OsProfile{
								LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
									Password: to.Ptr("**********"),
									Username: to.Ptr("sshuser"),
								},
							},
							ScriptActions:       []*armhdinsight.ScriptAction{},
							TargetInstanceCount: to.Ptr[int32](3),
						}},
				},
				MinSupportedTLSVersion: to.Ptr("1.2"),
				OSType:                 to.Ptr(armhdinsight.OSTypeLinux),
				StorageProfile: &armhdinsight.StorageProfile{
					Storageaccounts: []*armhdinsight.StorageAccount{
						{
							Name:       to.Ptr("storagename.blob.core.windows.net"),
							Container:  to.Ptr("contianername"),
							IsDefault:  to.Ptr(true),
							Key:        to.Ptr("*******"),
							ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/storagename"),
						}},
				},
				Tier: to.Ptr(armhdinsight.TierStandard),
			},
			Tags:               map[string]*string{},
			Name:               to.Ptr("testclustername"),
			Type:               to.Ptr("Microsoft.HDInsight/clusters"),
			FetchAaddsResource: to.Ptr(false),
			TenantID:           to.Ptr("00000000-0000-0000-0000-000000000000"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
