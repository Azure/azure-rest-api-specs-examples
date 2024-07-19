using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-05-15-preview/examples/throughputPool/CosmosDBThroughputPoolDelete.json
// this example is just showing the usage of "ThroughputPool_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBThroughputPoolResource created on azure
// for more information of creating CosmosDBThroughputPoolResource, please refer to the document of CosmosDBThroughputPoolResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "rgName";
string throughputPoolName = "tp1";
ResourceIdentifier cosmosDBThroughputPoolResourceId = CosmosDBThroughputPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, throughputPoolName);
CosmosDBThroughputPoolResource cosmosDBThroughputPool = client.GetCosmosDBThroughputPoolResource(cosmosDBThroughputPoolResourceId);

// invoke the operation
await cosmosDBThroughputPool.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
