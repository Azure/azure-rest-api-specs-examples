package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0219f9da9de5833fc380f57d6b026f68b5df6f93/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/InstallNewClusterLibraries.json
func ExampleClusterLibrariesClient_BeginManageLibraries_installNewClusterLibraries() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClusterLibrariesClient().BeginManageLibraries(ctx, "hiloResourceGroup", "clusterPool", "cluster", armhdinsightcontainers.ClusterLibraryManagementOperation{
		Properties: &armhdinsightcontainers.ClusterLibraryManagementOperationProperties{
			Action: to.Ptr(armhdinsightcontainers.LibraryManagementActionInstall),
			Libraries: []*armhdinsightcontainers.ClusterLibrary{
				{
					Properties: &armhdinsightcontainers.PyPiLibraryProperties{
						Type:    to.Ptr(armhdinsightcontainers.TypePypi),
						Remarks: to.Ptr("PyPi packages."),
						Name:    to.Ptr("requests"),
						Version: to.Ptr("2.31.0"),
					},
				},
				{
					Properties: &armhdinsightcontainers.MavenLibraryProperties{
						Type:    to.Ptr(armhdinsightcontainers.TypeMaven),
						Remarks: to.Ptr("Maven packages."),
						Name:    to.Ptr("flink-connector-kafka"),
						GroupID: to.Ptr("org.apache.flink"),
						Version: to.Ptr("3.0.2-1.18"),
					},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
