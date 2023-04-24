package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/ProviderInstances_Create.json
func ExampleProviderInstancesClient_BeginCreate_createASapMonitorHanaProvider() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloads.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProviderInstancesClient().BeginCreate(ctx, "myResourceGroup", "mySapMonitor", "myProviderInstance", armworkloads.ProviderInstance{
		Properties: &armworkloads.ProviderInstanceProperties{
			ProviderSettings: &armworkloads.HanaDbProviderInstanceProperties{
				ProviderType:             to.Ptr("SapHana"),
				DbName:                   to.Ptr("db"),
				DbPassword:               to.Ptr("****"),
				DbPasswordURI:            to.Ptr(""),
				DbUsername:               to.Ptr("user"),
				Hostname:                 to.Ptr("name"),
				InstanceNumber:           to.Ptr("00"),
				SapSid:                   to.Ptr("SID"),
				SQLPort:                  to.Ptr("0000"),
				SSLCertificateURI:        to.Ptr("https://storageaccount.blob.core.windows.net/containername/filename"),
				SSLHostNameInCertificate: to.Ptr("xyz.domain.com"),
				SSLPreference:            to.Ptr(armworkloads.SSLPreferenceServerCertificate),
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
	// 		ProviderSettings: &armworkloads.HanaDbProviderInstanceProperties{
	// 			ProviderType: to.Ptr("SapHana"),
	// 			DbName: to.Ptr("db"),
	// 			DbPasswordURI: to.Ptr(""),
	// 			DbUsername: to.Ptr("user"),
	// 			Hostname: to.Ptr("name"),
	// 			InstanceNumber: to.Ptr("00"),
	// 			SapSid: to.Ptr("SID"),
	// 			SQLPort: to.Ptr("0000"),
	// 			SSLCertificateURI: to.Ptr("https://storageaccount.blob.core.windows.net/containername/filename"),
	// 			SSLHostNameInCertificate: to.Ptr("xyz.domain.com"),
	// 			SSLPreference: to.Ptr(armworkloads.SSLPreferenceServerCertificate),
	// 		},
	// 		ProvisioningState: to.Ptr(armworkloads.WorkloadMonitorProvisioningStateSucceeded),
	// 	},
	// }
}
