using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBMongoDBCollectionRestore.json
// this example is just showing the usage of "MongoDBResources_CreateUpdateMongoDBCollection" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MongoDBDatabaseResource created on azure
// for more information of creating MongoDBDatabaseResource, please refer to the document of MongoDBDatabaseResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string databaseName = "databaseName";
ResourceIdentifier mongoDBDatabaseResourceId = MongoDBDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, databaseName);
MongoDBDatabaseResource mongoDBDatabase = client.GetMongoDBDatabaseResource(mongoDBDatabaseResourceId);

// get the collection of this MongoDBCollectionResource
MongoDBCollectionCollection collection = mongoDBDatabase.GetMongoDBCollections();

// invoke the operation
string collectionName = "collectionName";
MongoDBCollectionCreateOrUpdateContent content = new MongoDBCollectionCreateOrUpdateContent(new AzureLocation("West US"), new MongoDBCollectionResourceInfo("collectionName")
{
    RestoreParameters = new ResourceRestoreParameters
    {
        RestoreSource = "/subscriptions/subid/providers/Microsoft.DocumentDB/locations/WestUS/restorableDatabaseAccounts/restorableDatabaseAccountId",
        RestoreTimestampInUtc = DateTimeOffset.Parse("2022-07-20T18:28:00Z"),
        IsRestoreWithTtlDisabled = false,
    },
    CreateMode = CosmosDBAccountCreateMode.Restore,
})
{
    Options = new CosmosDBCreateUpdateConfig(),
    Tags = { },
};
ArmOperation<MongoDBCollectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, collectionName, content);
MongoDBCollectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MongoDBCollectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
