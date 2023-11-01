using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;
using Azure.ResourceManager.CosmosDB.Models;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/CosmosDBGremlinGraphGet.json
// this example is just showing the usage of "GremlinResources_GetGremlinGraph" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GremlinDatabaseResource created on azure
// for more information of creating GremlinDatabaseResource, please refer to the document of GremlinDatabaseResource
string subscriptionId = "subid";
string resourceGroupName = "rgName";
string accountName = "ddb1";
string databaseName = "databaseName";
ResourceIdentifier gremlinDatabaseResourceId = GremlinDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, databaseName);
GremlinDatabaseResource gremlinDatabase = client.GetGremlinDatabaseResource(gremlinDatabaseResourceId);

// get the collection of this GremlinGraphResource
GremlinGraphCollection collection = gremlinDatabase.GetGremlinGraphs();

// invoke the operation
string graphName = "graphName";
NullableResponse<GremlinGraphResource> response = await collection.GetIfExistsAsync(graphName);
GremlinGraphResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    GremlinGraphData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
