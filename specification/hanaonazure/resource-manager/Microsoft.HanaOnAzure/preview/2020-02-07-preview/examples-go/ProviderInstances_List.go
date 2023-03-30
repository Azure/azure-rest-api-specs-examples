package armhanaonazure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/ProviderInstances_List.json
func ExampleProviderInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhanaonazure.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProviderInstancesClient().NewListPager("myResourceGroup", "mySapMonitor", nil)
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
		// page.ProviderInstanceListResult = armhanaonazure.ProviderInstanceListResult{
		// 	Value: []*armhanaonazure.ProviderInstance{
		// 		{
		// 			Name: to.Ptr("myProviderInstance1"),
		// 			Type: to.Ptr("Microsoft.HanaOnAzure/sapMonitors/providerInstances"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.HanaOnAzure/sapMonitors/mySapMonitor/providerInstances/myProviderInstance1"),
		// 			Properties: &armhanaonazure.ProviderInstanceProperties{
		// 				Type: to.Ptr("SapHana"),
		// 				Metadata: to.Ptr("{\"key\":\"value\"}"),
		// 				Properties: to.Ptr("{\"hostname\":\"10.0.0.6\",\"dbName\":\"SYSTEMDB\",\"sqlPort\":30013,\"dbUsername\":\"SYSTEM\"}"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myProviderInstance2"),
		// 			Type: to.Ptr("Microsoft.HanaOnAzure/sapMonitors/providerInstances"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.HanaOnAzure/sapMonitors/mySapMonitor/providerInstances/myProviderInstance2"),
		// 			Properties: &armhanaonazure.ProviderInstanceProperties{
		// 				Type: to.Ptr("PrometheusHaCluster"),
		// 				Metadata: to.Ptr("{\"key\":\"value\"}"),
		// 				Properties: to.Ptr("{\"prometheusUrl\":\"http://10.0.0.21:9664/metrics\"}"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myProviderInstance2"),
		// 			Type: to.Ptr("Microsoft.HanaOnAzure/sapMonitors/providerInstances"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.HanaOnAzure/sapMonitors/mySapMonitor/providerInstances/myProviderInstance2"),
		// 			Properties: &armhanaonazure.ProviderInstanceProperties{
		// 				Type: to.Ptr("MsSqlServer"),
		// 				Metadata: to.Ptr("{\"key\":\"value\"}"),
		// 				Properties: to.Ptr("{\"sqlHostname\":\"10.0.0.6\",\"sqlPort\":1433,\"sqlUsername\":\"sqladmin\"}"),
		// 			},
		// 	}},
		// }
	}
}
