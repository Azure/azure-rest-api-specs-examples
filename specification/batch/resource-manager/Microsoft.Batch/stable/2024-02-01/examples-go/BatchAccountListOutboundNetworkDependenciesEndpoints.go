package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/BatchAccountListOutboundNetworkDependenciesEndpoints.json
func ExampleAccountClient_NewListOutboundNetworkDependenciesEndpointsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountClient().NewListOutboundNetworkDependenciesEndpointsPager("default-azurebatch-japaneast", "sampleacct", nil)
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
		// page.OutboundEnvironmentEndpointCollection = armbatch.OutboundEnvironmentEndpointCollection{
		// 	Value: []*armbatch.OutboundEnvironmentEndpoint{
		// 		{
		// 			Category: to.Ptr("Azure Batch"),
		// 			Endpoints: []*armbatch.EndpointDependency{
		// 				{
		// 					Description: to.Ptr("Applicable to job manager tasks, tasks that use job scoped authentication, or any task that makes calls to Batch."),
		// 					DomainName: to.Ptr("sampleacct.japaneast.batch.azure.com"),
		// 					EndpointDetails: []*armbatch.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					Description: to.Ptr("Applicable to all Azure Batch pools."),
		// 					DomainName: to.Ptr("japaneast.service.batch.azure.com"),
		// 					EndpointDetails: []*armbatch.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("Azure Storage"),
		// 			Endpoints: []*armbatch.EndpointDependency{
		// 				{
		// 					Description: to.Ptr("AutoStorage endpoint for this Batch account. Applicable to all Azure Batch pools under this account."),
		// 					DomainName: to.Ptr("autostorageaccountname.blob.core.windows.net"),
		// 					EndpointDetails: []*armbatch.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					Description: to.Ptr("Applicable to all Azure Batch pools."),
		// 					DomainName: to.Ptr("*.blob.core.windows.net"),
		// 					EndpointDetails: []*armbatch.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					Description: to.Ptr("Applicable to all Azure Batch pools."),
		// 					DomainName: to.Ptr("*.table.core.windows.net"),
		// 					EndpointDetails: []*armbatch.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 				},
		// 				{
		// 					Description: to.Ptr("Applicable to all Azure Batch pools."),
		// 					DomainName: to.Ptr("*.queue.core.windows.net"),
		// 					EndpointDetails: []*armbatch.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("Microsoft Package Repository"),
		// 			Endpoints: []*armbatch.EndpointDependency{
		// 				{
		// 					Description: to.Ptr("Only applicable to pools containing a Mount Configuration. Learn about Mount Configurations in Batch at https://docs.microsoft.com/azure/batch/virtual-file-mount."),
		// 					DomainName: to.Ptr("packages.microsoft.com"),
		// 					EndpointDetails: []*armbatch.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 		},
		// 		{
		// 			Category: to.Ptr("Azure Key Vault"),
		// 			Endpoints: []*armbatch.EndpointDependency{
		// 				{
		// 					Description: to.Ptr("Only applicable to pools containing a Disk Encryption Configuration and whose VM size does not support encryption at host. Learn more about disk encryption in Azure Batch at https://docs.microsoft.com/azure/batch/disk-encryption. Learn more about encryption at host and supported VM sizes at https://docs.microsoft.com/azure/virtual-machines/disks-enable-host-based-encryption-portal."),
		// 					DomainName: to.Ptr("*.vault.azure.net"),
		// 					EndpointDetails: []*armbatch.EndpointDetail{
		// 						{
		// 							Port: to.Ptr[int32](443),
		// 					}},
		// 			}},
		// 	}},
		// }
	}
}
