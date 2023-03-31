package armloadtesting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/LoadTests_ListOutboundNetworkDependenciesEndpoints.json
func ExampleLoadTestsClient_NewListOutboundNetworkDependenciesEndpointsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armloadtesting.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLoadTestsClient().NewListOutboundNetworkDependenciesEndpointsPager("default-azureloadtest-japaneast", "sampleloadtest", nil)
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
		// page.OutboundEnvironmentEndpointCollection = armloadtesting.OutboundEnvironmentEndpointCollection{
		// 	Value: []*armloadtesting.OutboundEnvironmentEndpoint{
		// 		{
		// 			Category: to.Ptr("Azure Batch"),
		// 			Endpoints: []*armloadtesting.EndpointDependency{
		// 				{
		// 					Description: to.Ptr("Applicable to job manager tasks, tasks that use job scoped authentication, or any task that makes calls to Batch."),
		// 					DomainName: to.Ptr("sampleacct.japaneast.batch.azure.com"),
		// 					EndpointDetails: []*armloadtesting.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					Description: to.Ptr("Applicable to all Azure Batch pools."),
		// 					DomainName: to.Ptr("japaneast.service.batch.azure.com"),
		// 					EndpointDetails: []*armloadtesting.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("Azure Storage"),
		// 			Endpoints: []*armloadtesting.EndpointDependency{
		// 				{
		// 					Description: to.Ptr("AutoStorage endpoint for this Batch account. Applicable to all Azure Batch pools under this account."),
		// 					DomainName: to.Ptr("autostorageaccountname.blob.core.windows.net"),
		// 					EndpointDetails: []*armloadtesting.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					Description: to.Ptr("Applicable to all Azure Batch pools."),
		// 					DomainName: to.Ptr("*.blob.core.windows.net"),
		// 					EndpointDetails: []*armloadtesting.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					Description: to.Ptr("Applicable to all Azure Batch pools."),
		// 					DomainName: to.Ptr("*.table.core.windows.net"),
		// 					EndpointDetails: []*armloadtesting.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					Description: to.Ptr("Applicable to all Azure Batch pools."),
		// 					DomainName: to.Ptr("*.queue.core.windows.net"),
		// 					EndpointDetails: []*armloadtesting.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("Microsoft Package Repository"),
		// 			Endpoints: []*armloadtesting.EndpointDependency{
		// 				{
		// 					Description: to.Ptr("Only applicable to pools containing a Mount Configuration. Learn about Mount Configurations in Batch at https://docs.microsoft.com/azure/batch/virtual-file-mount."),
		// 					DomainName: to.Ptr("packages.microsoft.com"),
		// 					EndpointDetails: []*armloadtesting.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("Azure Key Vault"),
		// 			Endpoints: []*armloadtesting.EndpointDependency{
		// 				{
		// 					Description: to.Ptr("Only applicable to pools containing a Disk Encryption Configuration and whose VM size does not support encryption at host. Learn more about disk encryption in Azure Batch at https://docs.microsoft.com/azure/batch/disk-encryption. Learn more about encryption at host and supported VM sizes at https://docs.microsoft.com/azure/virtual-machines/disks-enable-host-based-encryption-portal."),
		// 					DomainName: to.Ptr("*.vault.azure.net"),
		// 					EndpointDetails: []*armloadtesting.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 	}},
		// }
	}
}
