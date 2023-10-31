using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;
using Azure.ResourceManager.CosmosDB.Models;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/CosmosDBMongoDBCollectionDelete.json
// this example is just showing the usage of "MongoDBResources_DeleteMongoDBCollection" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MongoDBCollectionResource created on azure
// for more information of creating MongoDBCollectionResource, please refer to the document of MongoDBCollectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string databaseName = "databaseName";
string collectionName = "collectionName";
ResourceIdentifier mongoDBCollectionResourceId = MongoDBCollectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, databaseName, collectionName);
MongoDBCollectionResource mongoDBCollection = client.GetMongoDBCollectionResource(mongoDBCollectionResourceId);

// invoke the operation
await mongoDBCollection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
