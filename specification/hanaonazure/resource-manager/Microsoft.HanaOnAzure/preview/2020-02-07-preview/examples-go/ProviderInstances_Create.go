package armhanaonazure_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2020-02-07-preview/examples/ProviderInstances_Create.json
func ExampleProviderInstancesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhanaonazure.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProviderInstancesClient().BeginCreate(ctx, "myResourceGroup", "mySapMonitor", "myProviderInstance", armhanaonazure.ProviderInstance{
		Properties: &armhanaonazure.ProviderInstanceProperties{
			Type:       to.Ptr("hana"),
			Metadata:   to.Ptr("{\"key\":\"value\"}"),
			Properties: to.Ptr("{\"hostname\":\"10.0.0.6\",\"dbName\":\"SYSTEMDB\",\"sqlPort\":30013,\"dbUsername\":\"SYSTEM\",\"dbPassword\":\"PASSWORD\"}"),
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
	// res.ProviderInstance = armhanaonazure.ProviderInstance{
	// 	Name: to.Ptr("myProviderInstance"),
	// 	Type: to.Ptr("Microsoft.HanaOnAzure/sapMonitors/providerInstances"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.HanaOnAzure/sapMonitors/mySapMonitor/providerInstances/myProviderInstance"),
	// 	Properties: &armhanaonazure.ProviderInstanceProperties{
	// 		Type: to.Ptr("SapHana"),
	// 		Metadata: to.Ptr("{\"key\":\"value\"}"),
	// 		Properties: to.Ptr("{\"hostname\":\"10.0.0.6\",\"dbName\":\"SYSTEMDB\",\"sqlPort\":30013,\"dbUsername\":\"SYSTEM\"}"),
	// 		ProvisioningState: to.Ptr(armhanaonazure.HanaProvisioningStatesEnumSucceeded),
	// 	},
	// }
}
