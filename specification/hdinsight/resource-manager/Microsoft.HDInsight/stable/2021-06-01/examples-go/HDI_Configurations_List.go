package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/HDI_Configurations_List.json
func ExampleConfigurationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationsClient().List(ctx, "rg1", "cluster1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ClusterConfigurations = armhdinsight.ClusterConfigurations{
	// 	Configurations: map[string]map[string]*string{
	// 		"core-site": map[string]*string{
	// 			"fs.azure.account.key.test.blob.core.windows.net": to.Ptr("accountKey"),
	// 			"fs.defaultFS": to.Ptr("wasb://test.blob.core.windows.net"),
	// 		},
	// 		"gateway": map[string]*string{
	// 			"restAuthCredential.isEnabled": to.Ptr("true"),
	// 			"restAuthCredential.password": to.Ptr("**********"),
	// 			"restAuthCredential.username": to.Ptr("admin"),
	// 		},
	// 	},
	// }
}
