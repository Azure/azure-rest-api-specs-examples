using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-05-15-preview/examples/throughputPool/CosmosDBThroughputPoolAccountDelete.json
// this example is just showing the usage of "ThroughputPoolAccount_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBThroughputPoolAccountResource created on azure
// for more information of creating CosmosDBThroughputPoolAccountResource, please refer to the document of CosmosDBThroughputPoolAccountResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "rgName";
string throughputPoolName = "tp1";
string throughputPoolAccountName = "db1";
ResourceIdentifier cosmosDBThroughputPoolAccountResourceId = CosmosDBThroughputPoolAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, throughputPoolName, throughputPoolAccountName);
CosmosDBThroughputPoolAccountResource cosmosDBThroughputPoolAccount = client.GetCosmosDBThroughputPoolAccountResource(cosmosDBThroughputPoolAccountResourceId);

// invoke the operation
await cosmosDBThroughputPoolAccount.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
