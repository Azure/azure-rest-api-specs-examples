using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Batch;
using Azure.ResourceManager.Batch.Models;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/PoolStopResize.json
// this example is just showing the usage of "Pool_StopResize" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BatchAccountPoolResource created on azure
// for more information of creating BatchAccountPoolResource, please refer to the document of BatchAccountPoolResource
string subscriptionId = "subid";
string resourceGroupName = "default-azurebatch-japaneast";
string accountName = "sampleacct";
string poolName = "testpool";
ResourceIdentifier batchAccountPoolResourceId = BatchAccountPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName);
BatchAccountPoolResource batchAccountPool = client.GetBatchAccountPoolResource(batchAccountPoolResourceId);

// invoke the operation
BatchAccountPoolResource result = await batchAccountPool.StopResizeAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BatchAccountPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
