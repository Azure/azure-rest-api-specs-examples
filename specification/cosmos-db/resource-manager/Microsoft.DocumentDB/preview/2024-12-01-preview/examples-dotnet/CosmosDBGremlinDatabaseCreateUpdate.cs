using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBGremlinDatabaseCreateUpdate.json
// this example is just showing the usage of "GremlinResources_CreateUpdateGremlinDatabase" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GremlinDatabaseResource created on azure
// for more information of creating GremlinDatabaseResource, please refer to the document of GremlinDatabaseResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string databaseName = "databaseName";
ResourceIdentifier gremlinDatabaseResourceId = GremlinDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, databaseName);
GremlinDatabaseResource gremlinDatabase = client.GetGremlinDatabaseResource(gremlinDatabaseResourceId);

// invoke the operation
GremlinDatabaseCreateOrUpdateContent content = new GremlinDatabaseCreateOrUpdateContent(new AzureLocation("West US"), new GremlinDatabaseResourceInfo("databaseName"))
{
    Options = new CosmosDBCreateUpdateConfig(),
    Tags = { },
};
ArmOperation<GremlinDatabaseResource> lro = await gremlinDatabase.UpdateAsync(WaitUntil.Completed, content);
GremlinDatabaseResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GremlinDatabaseData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
