using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-05-15-preview/examples/CosmosDBMongoDBDatabaseCreateUpdate.json
// this example is just showing the usage of "MongoDBResources_CreateUpdateMongoDBDatabase" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this MongoDBDatabaseResource
MongoDBDatabaseCollection collection = cosmosDBAccount.GetMongoDBDatabases();

// invoke the operation
string databaseName = "databaseName";
MongoDBDatabaseCreateOrUpdateContent content = new MongoDBDatabaseCreateOrUpdateContent(new AzureLocation("West US"), new MongoDBDatabaseResourceInfo("databaseName"))
{
    Options = new CosmosDBCreateUpdateConfig(),
    Tags =
    {
    },
};
ArmOperation<MongoDBDatabaseResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, databaseName, content);
MongoDBDatabaseResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MongoDBDatabaseData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
