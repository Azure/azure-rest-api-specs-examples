package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/workloadmonitor/Db2ProviderInstances_Create_Root_Certificate.json
func ExampleProviderInstancesClient_BeginCreate_createADb2ProviderWithRootCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armworkloads.NewProviderInstancesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "myResourceGroup", "mySapMonitor", "myProviderInstance", armworkloads.ProviderInstance{
		Properties: &armworkloads.ProviderInstanceProperties{
			ProviderSettings: &armworkloads.DB2ProviderInstanceProperties{
				ProviderType:  to.Ptr("Db2"),
				DbName:        to.Ptr("dbName"),
				DbPassword:    to.Ptr("password"),
				DbPasswordURI: to.Ptr(""),
				DbPort:        to.Ptr("dbPort"),
				DbUsername:    to.Ptr("username"),
				Hostname:      to.Ptr("hostname"),
				SapSid:        to.Ptr("SID"),
				SSLPreference: to.Ptr(armworkloads.SSLPreferenceRootCertificate),
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
	// 	Type: to.Ptr("Microsoft.Workloads/workloads/providerInstances"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Workloads/workloads/mySapMonitor/providerInstances/myProviderInstance"),
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
	// 			SSLPreference: to.Ptr(armworkloads.SSLPreferenceRootCertificate),
	// 		},
	// 		ProvisioningState: to.Ptr(armworkloads.WorkloadMonitorProvisioningStateSucceeded),
	// 	},
	// }
}
