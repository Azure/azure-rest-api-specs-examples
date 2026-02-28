using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Batch;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/Batch/stable/2025-06-01/examples/PrivateEndpointConnectionsList.json
// this example is just showing the usage of "PrivateEndpointConnection_ListByBatchAccount" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BatchAccountResource created on azure
// for more information of creating BatchAccountResource, please refer to the document of BatchAccountResource
string subscriptionId = "12345678-1234-1234-1234-123456789012";
string resourceGroupName = "default-azurebatch-japaneast";
string accountName = "sampleacct";
ResourceIdentifier batchAccountResourceId = BatchAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
BatchAccountResource batchAccount = client.GetBatchAccountResource(batchAccountResourceId);

// get the collection of this BatchPrivateEndpointConnectionResource
BatchPrivateEndpointConnectionCollection collection = batchAccount.GetBatchPrivateEndpointConnections();

// invoke the operation and iterate over the result
await foreach (BatchPrivateEndpointConnectionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    BatchPrivateEndpointConnectionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
