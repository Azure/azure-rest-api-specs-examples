package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0219f9da9de5833fc380f57d6b026f68b5df6f93/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/CreateRangerCluster.json
func ExampleClustersClient_BeginCreate_hdInsightRangerClusterPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginCreate(ctx, "hiloResourcegroup", "clusterpool1", "cluster1", armhdinsightcontainers.Cluster{
		Location: to.Ptr("West US 2"),
		Properties: &armhdinsightcontainers.ClusterResourceProperties{
			ClusterProfile: &armhdinsightcontainers.ClusterProfile{
				AuthorizationProfile: &armhdinsightcontainers.AuthorizationProfile{
					UserIDs: []*string{
						to.Ptr("testuser1"),
						to.Ptr("testuser2")},
				},
				ClusterVersion: to.Ptr("0.0.1"),
				ManagedIdentityProfile: &armhdinsightcontainers.ManagedIdentityProfile{
					IdentityList: []*armhdinsightcontainers.ManagedIdentitySpec{
						{
							Type:       to.Ptr(armhdinsightcontainers.ManagedIdentityTypeCluster),
							ClientID:   to.Ptr("de91f1d8-767f-460a-ac11-3cf103f74b34"),
							ObjectID:   to.Ptr("40491351-c240-4042-91e0-f644a1d2b441"),
							ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-msi"),
						}},
				},
				OssVersion: to.Ptr("2.2.3"),
				RangerProfile: &armhdinsightcontainers.RangerProfile{
					RangerAdmin: &armhdinsightcontainers.RangerAdminSpec{
						Admins: []*string{
							to.Ptr("testuser1@contoso.com"),
							to.Ptr("testuser2@contoso.com")},
						Database: &armhdinsightcontainers.RangerAdminSpecDatabase{
							Name:              to.Ptr("testdb"),
							Host:              to.Ptr("testsqlserver.database.windows.net"),
							PasswordSecretRef: to.Ptr("https://testkv.vault.azure.net/secrets/mysecret/5df6584d9c25418c8d900240aa6c3452"),
							Username:          to.Ptr("admin"),
						},
					},
					RangerAudit: &armhdinsightcontainers.RangerAuditSpec{
						StorageAccount: to.Ptr("https://teststorage.blob.core.windows.net/testblob"),
					},
					RangerUsersync: &armhdinsightcontainers.RangerUsersyncSpec{
						Enabled: to.Ptr(true),
						Groups: []*string{
							to.Ptr("0a53828f-36c9-44c3-be3d-99a7fce977ad"),
							to.Ptr("13be6971-79db-4f33-9d41-b25589ca25ac")},
						Mode: to.Ptr(armhdinsightcontainers.RangerUsersyncModeAutomatic),
						Users: []*string{
							to.Ptr("testuser1@contoso.com"),
							to.Ptr("testuser2@contoso.com")},
					},
				},
			},
			ClusterType: to.Ptr("ranger"),
			ComputeProfile: &armhdinsightcontainers.ComputeProfile{
				AvailabilityZones: []*string{
					to.Ptr("1"),
					to.Ptr("2"),
					to.Ptr("3")},
				Nodes: []*armhdinsightcontainers.NodeProfile{
					{
						Type:   to.Ptr("head"),
						Count:  to.Ptr[int32](2),
						VMSize: to.Ptr("Standard_D3_v2"),
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
	// res.Cluster = armhdinsightcontainers.Cluster{
	// 	Name: to.Ptr("cluster1"),
	// 	Type: to.Ptr("Microsoft.HDInsight/clusterPools/clusters"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.HDInsight/clusterPools/clusterpool1/clusters/cluster1"),
	// 	SystemData: &armhdinsightcontainers.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armhdinsightcontainers.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.197Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armhdinsightcontainers.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("West US 2"),
	// 	Properties: &armhdinsightcontainers.ClusterResourceProperties{
	// 		ClusterProfile: &armhdinsightcontainers.ClusterProfile{
	// 			AuthorizationProfile: &armhdinsightcontainers.AuthorizationProfile{
	// 				UserIDs: []*string{
	// 					to.Ptr("testuser1"),
	// 					to.Ptr("testuser2")},
	// 				},
	// 				ClusterVersion: to.Ptr("0.0.1"),
	// 				Components: []*armhdinsightcontainers.ClusterComponentsItem{
	// 					{
	// 						Name: to.Ptr("HDFS"),
	// 						Version: to.Ptr("2.2.3"),
	// 				}},
	// 				ConnectivityProfile: &armhdinsightcontainers.ConnectivityProfile{
	// 					Web: &armhdinsightcontainers.ConnectivityProfileWeb{
	// 						Fqdn: to.Ptr("cluster1.clusterpool1.westus2.projecthilo.net"),
	// 					},
	// 				},
	// 				ManagedIdentityProfile: &armhdinsightcontainers.ManagedIdentityProfile{
	// 					IdentityList: []*armhdinsightcontainers.ManagedIdentitySpec{
	// 						{
	// 							Type: to.Ptr(armhdinsightcontainers.ManagedIdentityTypeCluster),
	// 							ClientID: to.Ptr("de91f1d8-767f-460a-ac11-3cf103f74b34"),
	// 							ObjectID: to.Ptr("40491351-c240-4042-91e0-f644a1d2b441"),
	// 							ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/hiloResourcegroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-msi"),
	// 					}},
	// 				},
	// 				OssVersion: to.Ptr("2.2.3"),
	// 				RangerProfile: &armhdinsightcontainers.RangerProfile{
	// 					RangerAdmin: &armhdinsightcontainers.RangerAdminSpec{
	// 						Admins: []*string{
	// 							to.Ptr("testuser1@contoso.com"),
	// 							to.Ptr("testuser2@contoso.com")},
	// 							Database: &armhdinsightcontainers.RangerAdminSpecDatabase{
	// 								Name: to.Ptr("testdb"),
	// 								Host: to.Ptr("testsqlserver.database.windows.net"),
	// 								PasswordSecretRef: to.Ptr("https://testkv.vault.azure.net/secrets/mysecret/5df6584d9c25418c8d900240aa6c3452"),
	// 								Username: to.Ptr("admin"),
	// 							},
	// 						},
	// 						RangerAudit: &armhdinsightcontainers.RangerAuditSpec{
	// 							StorageAccount: to.Ptr("https://teststorage.blob.core.windows.net/testblob"),
	// 						},
	// 						RangerUsersync: &armhdinsightcontainers.RangerUsersyncSpec{
	// 							Enabled: to.Ptr(true),
	// 							Groups: []*string{
	// 								to.Ptr("0a53828f-36c9-44c3-be3d-99a7fce977ad"),
	// 								to.Ptr("13be6971-79db-4f33-9d41-b25589ca25ac")},
	// 								Mode: to.Ptr(armhdinsightcontainers.RangerUsersyncModeAutomatic),
	// 								Users: []*string{
	// 									to.Ptr("testuser1@contoso.com"),
	// 									to.Ptr("testuser2@contoso.com")},
	// 								},
	// 							},
	// 						},
	// 						ClusterType: to.Ptr("ranger"),
	// 						ComputeProfile: &armhdinsightcontainers.ComputeProfile{
	// 							AvailabilityZones: []*string{
	// 								to.Ptr("1"),
	// 								to.Ptr("2"),
	// 								to.Ptr("3")},
	// 								Nodes: []*armhdinsightcontainers.NodeProfile{
	// 									{
	// 										Type: to.Ptr("head"),
	// 										Count: to.Ptr[int32](2),
	// 										VMSize: to.Ptr("Standard_D3_v2"),
	// 								}},
	// 							},
	// 							DeploymentID: to.Ptr("45cd32aead6e4a91b079a0cdbfac8c36"),
	// 							ProvisioningState: to.Ptr(armhdinsightcontainers.ProvisioningStatusSucceeded),
	// 						},
	// 					}
}
