package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/Db2ProviderInstances_Get.json
func ExampleProviderInstancesClient_Get_getPropertiesOfADb2Provider() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloads.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProviderInstancesClient().Get(ctx, "myResourceGroup", "mySapMonitor", "myProviderInstance", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProviderInstance = armworkloads.ProviderInstance{
	// 	Name: to.Ptr("myProviderInstance"),
	// 	Type: to.Ptr("Microsoft.Workloads/monitors/providerInstances"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Workloads/monitors/mySapMonitor/providerInstances/myProviderInstance"),
	// 	SystemData: &armworkloads.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@xyz.com"),
	// 		CreatedByType: to.Ptr(armworkloads.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@xyz.com"),
	// 		LastModifiedByType: to.Ptr(armworkloads.CreatedByTypeUser),
	// 	},
	// 	Properties: &armworkloads.ProviderInstanceProperties{
	// 		ProviderSettings: &armworkloads.DB2ProviderInstanceProperties{
	// 			ProviderType: to.Ptr("Db2"),
	// 			DbName: to.Ptr("OPA"),
	// 			DbPasswordURI: to.Ptr(""),
	// 			DbPort: to.Ptr("5912"),
	// 			DbUsername: to.Ptr("Db2OPA"),
	// 			Hostname: to.Ptr("vmname.azure.com"),
	// 			SapSid: to.Ptr("SID"),
	// 			SSLCertificateURI: to.Ptr("https://storageaccount.blob.core.windows.net/containername/filename"),
	// 			SSLPreference: to.Ptr(armworkloads.SSLPreferenceServerCertificate),
	// 		},
	// 		ProvisioningState: to.Ptr(armworkloads.WorkloadMonitorProvisioningStateSucceeded),
	// 	},
	// }
}
