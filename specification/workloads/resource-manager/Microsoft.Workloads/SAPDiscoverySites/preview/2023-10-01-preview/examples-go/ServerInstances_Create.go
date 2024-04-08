package armmigrationdiscoverysap_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationdiscovery/armmigrationdiscoverysap"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/475747ff6322e9bf180b8911d871561b264379c3/specification/workloads/resource-manager/Microsoft.Workloads/SAPDiscoverySites/preview/2023-10-01-preview/examples/ServerInstances_Create.json
func ExampleServerInstancesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationdiscoverysap.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerInstancesClient().BeginCreate(ctx, "test-rg", "SampleSite", "MPP_MPP", "APP_SapServer1", armmigrationdiscoverysap.ServerInstance{}, nil)
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
	// res.ServerInstance = armmigrationdiscoverysap.ServerInstance{
	// 	Name: to.Ptr("APP_SapServer1"),
	// 	Type: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances/serverInstances"),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapDiscoverySites/SampleSite/sapInstances/MPP_MPP/serverInstances/APP_SapServer1"),
	// 	SystemData: &armmigrationdiscoverysap.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-26T14:15:22.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("UserName"),
	// 		CreatedByType: to.Ptr(armmigrationdiscoverysap.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-26T14:15:22.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("UserName"),
	// 		LastModifiedByType: to.Ptr(armmigrationdiscoverysap.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmigrationdiscoverysap.ServerInstanceProperties{
	// 		ConfigurationData: &armmigrationdiscoverysap.ConfigurationData{
	// 			CPU: to.Ptr[int32](8),
	// 			CPUInMhz: to.Ptr[int32](2300),
	// 			CPUType: to.Ptr("AMD EPYC 7452 Processor"),
	// 			DatabaseType: to.Ptr(armmigrationdiscoverysap.DatabaseTypeAdabas),
	// 			HardwareManufacturer: to.Ptr("Microsoft Corporation"),
	// 			Model: to.Ptr("Virtual Machine"),
	// 			RAM: to.Ptr[int32](256),
	// 			Saps: to.Ptr[int32](7000),
	// 			TargetHanaRAMSizeGB: to.Ptr[int32](512),
	// 			TotalDiskIops: to.Ptr[int32](1000),
	// 			TotalDiskSizeGB: to.Ptr[int32](512),
	// 		},
	// 		InstanceSid: to.Ptr("MPP"),
	// 		OperatingSystem: to.Ptr(armmigrationdiscoverysap.OperatingSystem("IBM")),
	// 		PerformanceData: &armmigrationdiscoverysap.ExcelPerformanceData{
	// 			DataSource: to.Ptr(armmigrationdiscoverysap.DataSourceExcel),
	// 			MaxCPULoad: to.Ptr[int32](10),
	// 			TotalSourceDbSizeGB: to.Ptr[int32](750),
	// 		},
	// 		ProvisioningState: to.Ptr(armmigrationdiscoverysap.ProvisioningStateSucceeded),
	// 		SapInstanceType: to.Ptr(armmigrationdiscoverysap.SapInstanceTypeAPP),
	// 		SapProduct: to.Ptr("SAP ERP ENHANCE PACKAGE"),
	// 		SapProductVersion: to.Ptr("6.08"),
	// 		ServerName: to.Ptr("SapServer1"),
	// 	},
	// }
}
