package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0219f9da9de5833fc380f57d6b026f68b5df6f93/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2024-05-01-preview/examples/ListPredefinedClusterLibraries.json
func ExampleClusterLibrariesClient_NewListPager_listPredefinedClusterLibraries() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClusterLibrariesClient().NewListPager("hiloResourceGroup", "clusterPool", "cluster", armhdinsightcontainers.CategoryPredefined, nil)
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
		// page.ClusterLibraryList = armhdinsightcontainers.ClusterLibraryList{
		// 	Value: []*armhdinsightcontainers.ClusterLibrary{
		// 		{
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hiloResourceGroup/providers/Microsoft.HDInsight/clusterPools/clusterPool/clusters/cluster/libraries/pyspark"),
		// 			Properties: &armhdinsightcontainers.PyPiLibraryProperties{
		// 				Type: to.Ptr(armhdinsightcontainers.TypePypi),
		// 				Name: to.Ptr("pyspark"),
		// 				Version: to.Ptr("3.5.0"),
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hiloResourceGroup/providers/Microsoft.HDInsight/clusterPools/clusterPool/clusters/cluster/libraries/spark-sql"),
		// 			Properties: &armhdinsightcontainers.MavenLibraryProperties{
		// 				Type: to.Ptr(armhdinsightcontainers.TypeMaven),
		// 				Name: to.Ptr("spark-sql"),
		// 				GroupID: to.Ptr("org.apache.spark"),
		// 				Version: to.Ptr("3.5.0_2.12"),
		// 			},
		// 	}},
		// }
	}
}
