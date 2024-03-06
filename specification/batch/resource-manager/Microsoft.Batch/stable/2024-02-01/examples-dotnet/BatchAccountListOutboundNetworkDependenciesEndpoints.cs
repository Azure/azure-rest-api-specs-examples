using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Batch;
using Azure.ResourceManager.Batch.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/BatchAccountListOutboundNetworkDependenciesEndpoints.json
// this example is just showing the usage of "BatchAccount_ListOutboundNetworkDependenciesEndpoints" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BatchAccountResource created on azure
// for more information of creating BatchAccountResource, please refer to the document of BatchAccountResource
string subscriptionId = "subid";
string resourceGroupName = "default-azurebatch-japaneast";
string accountName = "sampleacct";
ResourceIdentifier batchAccountResourceId = BatchAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
BatchAccountResource batchAccount = client.GetBatchAccountResource(batchAccountResourceId);

// invoke the operation and iterate over the result
await foreach (BatchAccountOutboundEnvironmentEndpoint item in batchAccount.GetOutboundNetworkDependenciesEndpointsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
