using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBDatabaseGetUsages.json
// this example is just showing the usage of "Database_ListUsages" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBAccountResource created on azure
// for more information of creating CosmosDBAccountResource, please refer to the document of CosmosDBAccountResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
ResourceIdentifier cosmosDBAccountResourceId = CosmosDBAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
CosmosDBAccountResource cosmosDBAccount = client.GetCosmosDBAccountResource(cosmosDBAccountResourceId);

// invoke the operation and iterate over the result
string databaseRid = "databaseRid";
string filter = "$filter=name.value eq 'Storage'";
await foreach (CosmosDBBaseUsage item in cosmosDBAccount.GetUsagesDatabasesAsync(databaseRid, filter: filter))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
