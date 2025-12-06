using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/CosmosDBGremlinGraphDelete.json
// this example is just showing the usage of "GremlinResources_DeleteGremlinGraph" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GremlinGraphResource created on azure
// for more information of creating GremlinGraphResource, please refer to the document of GremlinGraphResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string databaseName = "databaseName";
string graphName = "graphName";
ResourceIdentifier gremlinGraphResourceId = GremlinGraphResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, databaseName, graphName);
GremlinGraphResource gremlinGraph = client.GetGremlinGraphResource(gremlinGraphResourceId);

// invoke the operation
await gremlinGraph.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
