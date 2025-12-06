using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/CosmosDBRestorableGremlinGraphList.json
// this example is just showing the usage of "RestorableGremlinGraphs_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RestorableCosmosDBAccountResource created on azure
// for more information of creating RestorableCosmosDBAccountResource, please refer to the document of RestorableCosmosDBAccountResource
string subscriptionId = "subid";
AzureLocation location = new AzureLocation("WestUS");
Guid instanceId = Guid.Parse("98a570f2-63db-4117-91f0-366327b7b353");
ResourceIdentifier restorableCosmosDBAccountResourceId = RestorableCosmosDBAccountResource.CreateResourceIdentifier(subscriptionId, location, instanceId);
RestorableCosmosDBAccountResource restorableCosmosDBAccount = client.GetRestorableCosmosDBAccountResource(restorableCosmosDBAccountResourceId);

// invoke the operation and iterate over the result
string restorableGremlinDatabaseRid = "PD5DALigDgw=";
await foreach (RestorableGremlinGraph item in restorableCosmosDBAccount.GetRestorableGremlinGraphsAsync(restorableGremlinDatabaseRid: restorableGremlinDatabaseRid))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
