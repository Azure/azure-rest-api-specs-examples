using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Batch;
using Azure.ResourceManager.Batch.Models;
using Azure.ResourceManager.Models;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolListWithFilter.json
// this example is just showing the usage of "Pool_ListByBatchAccount" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this BatchAccountPoolResource
BatchAccountPoolCollection collection = batchAccount.GetBatchAccountPools();

// invoke the operation and iterate over the result
string select = "properties/allocationState,properties/provisioningStateTransitionTime,properties/currentDedicatedNodes,properties/currentLowPriorityNodes";
string filter = "startswith(name, 'po') or (properties/allocationState eq 'Steady' and properties/provisioningStateTransitionTime lt datetime'2017-02-02')";
await foreach (BatchAccountPoolResource item in collection.GetAllAsync(select: select, filter: filter))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    BatchAccountPoolData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
