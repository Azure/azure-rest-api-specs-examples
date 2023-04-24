package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/ProviderInstances_List.json
func ExampleProviderInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloads.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.ProviderInstanceListResult = armworkloads.ProviderInstanceListResult{
		// 	Value: []*armworkloads.ProviderInstance{
		// 		{
		// 			Name: to.Ptr("myProviderInstance1"),
		// 			Type: to.Ptr("Microsoft.Workloads/monitors/providerInstances"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Workloads/monitors/mySapMonitor/providerInstances/myProviderInstance1"),
		// 			Properties: &armworkloads.ProviderInstanceProperties{
		// 				ProviderSettings: &armworkloads.HanaDbProviderInstanceProperties{
		// 					ProviderType: to.Ptr("SapHana"),
		// 					DbName: to.Ptr("db"),
		// 					DbPasswordURI: to.Ptr(""),
		// 					DbUsername: to.Ptr("user"),
		// 					Hostname: to.Ptr("name"),
		// 					InstanceNumber: to.Ptr("00"),
		// 					SapSid: to.Ptr("SID"),
		// 					SQLPort: to.Ptr("0000"),
		// 					SSLCertificateURI: to.Ptr("https://storageaccount.blob.core.windows.net/containername/filename"),
		// 					SSLHostNameInCertificate: to.Ptr("xyz.domain.com"),
		// 					SSLPreference: to.Ptr(armworkloads.SSLPreferenceServerCertificate),
		// 				},
		// 				ProvisioningState: to.Ptr(armworkloads.WorkloadMonitorProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myProviderInstance2"),
		// 			Type: to.Ptr("Microsoft.Workloads/monitors/providerInstances"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Workloads/monitors/mySapMonitor/providerInstances/myProviderInstance1"),
		// 			Properties: &armworkloads.ProviderInstanceProperties{
		// 				ProviderSettings: &armworkloads.SapNetWeaverProviderInstanceProperties{
		// 					ProviderType: to.Ptr("SapNetWeaver"),
		// 					SapClientID: to.Ptr("111"),
		// 					SapHostFileEntries: []*string{
		// 						to.Ptr("127.0.0.1 name fqdn")},
		// 						SapHostname: to.Ptr("name"),
		// 						SapInstanceNr: to.Ptr("00"),
		// 						SapPasswordURI: to.Ptr(""),
		// 						SapPortNumber: to.Ptr("1234"),
		// 						SapSid: to.Ptr("SID"),
		// 						SapUsername: to.Ptr("username"),
		// 						SSLCertificateURI: to.Ptr("https://storageaccount.blob.core.windows.net/containername/filename"),
		// 						SSLPreference: to.Ptr(armworkloads.SSLPreferenceServerCertificate),
		// 					},
		// 					ProvisioningState: to.Ptr(armworkloads.WorkloadMonitorProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myProviderInstance4"),
		// 				Type: to.Ptr("Microsoft.Workloads/monitors/providerInstances"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Workloads/monitors/mySapMonitor/providerInstances/myProviderInstance1"),
		// 				Properties: &armworkloads.ProviderInstanceProperties{
		// 					ProviderSettings: &armworkloads.MsSQLServerProviderInstanceProperties{
		// 						ProviderType: to.Ptr("MsSqlServer"),
		// 						DbPort: to.Ptr("5912"),
		// 						DbUsername: to.Ptr("user"),
		// 						Hostname: to.Ptr("hostname"),
		// 						SapSid: to.Ptr("sid"),
		// 						SSLCertificateURI: to.Ptr("https://storageaccount.blob.core.windows.net/containername/filename"),
		// 						SSLPreference: to.Ptr(armworkloads.SSLPreferenceServerCertificate),
		// 					},
		// 					ProvisioningState: to.Ptr(armworkloads.WorkloadMonitorProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myProviderInstance3"),
		// 				Type: to.Ptr("Microsoft.Workloads/monitors/providerInstances"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Workloads/monitors/mySapMonitor/providerInstances/myProviderInstance1"),
		// 				Properties: &armworkloads.ProviderInstanceProperties{
		// 					ProviderSettings: &armworkloads.PrometheusOSProviderInstanceProperties{
		// 						ProviderType: to.Ptr("PrometheusOS"),
		// 						PrometheusURL: to.Ptr("http://192.168.0.0:9090/metrics"),
		// 						SapSid: to.Ptr("SID"),
		// 						SSLCertificateURI: to.Ptr("https://storageaccount.blob.core.windows.net/containername/filename"),
		// 						SSLPreference: to.Ptr(armworkloads.SSLPreferenceServerCertificate),
		// 					},
		// 					ProvisioningState: to.Ptr(armworkloads.WorkloadMonitorProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myProviderInstance3"),
		// 				Type: to.Ptr("Microsoft.Workloads/monitors/providerInstances"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Workloads/monitors/mySapMonitor/providerInstances/myProviderInstance1"),
		// 				Properties: &armworkloads.ProviderInstanceProperties{
		// 					ProviderSettings: &armworkloads.PrometheusHaClusterProviderInstanceProperties{
		// 						ProviderType: to.Ptr("PrometheusHaCluster"),
		// 						ClusterName: to.Ptr("clusterName"),
		// 						Hostname: to.Ptr("hostname"),
		// 						PrometheusURL: to.Ptr("http://192.168.0.0:9090/metrics"),
		// 						Sid: to.Ptr("sid"),
		// 						SSLCertificateURI: to.Ptr("https://storageaccount.blob.core.windows.net/containername/filename"),
		// 						SSLPreference: to.Ptr(armworkloads.SSLPreferenceServerCertificate),
		// 					},
		// 					ProvisioningState: to.Ptr(armworkloads.WorkloadMonitorProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myProviderInstance3"),
		// 				Type: to.Ptr("Microsoft.Workloads/monitors/providerInstances"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Workloads/monitors/mySapMonitor/providerInstances/myProviderInstance1"),
		// 				Properties: &armworkloads.ProviderInstanceProperties{
		// 					ProviderSettings: &armworkloads.DB2ProviderInstanceProperties{
		// 						ProviderType: to.Ptr("Db2"),
		// 						DbName: to.Ptr("OPA"),
		// 						DbPasswordURI: to.Ptr(""),
		// 						DbPort: to.Ptr("5912"),
		// 						DbUsername: to.Ptr("Db2OPA"),
		// 						Hostname: to.Ptr("vmname.azure.com"),
		// 						SapSid: to.Ptr("SID"),
		// 						SSLCertificateURI: to.Ptr("https://storageaccount.blob.core.windows.net/containername/filename"),
		// 						SSLPreference: to.Ptr(armworkloads.SSLPreferenceServerCertificate),
		// 					},
		// 					ProvisioningState: to.Ptr(armworkloads.WorkloadMonitorProvisioningStateSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}
