using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;
using Azure.ResourceManager.CosmosDB.Models;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/CosmosDBSqlClientEncryptionKeyGet.json
// this example is just showing the usage of "SqlResources_GetClientEncryptionKey" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBSqlDatabaseResource created on azure
// for more information of creating CosmosDBSqlDatabaseResource, please refer to the document of CosmosDBSqlDatabaseResource
string subscriptionId = "subId";
string resourceGroupName = "rgName";
string accountName = "accountName";
string databaseName = "databaseName";
ResourceIdentifier cosmosDBSqlDatabaseResourceId = CosmosDBSqlDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, databaseName);
CosmosDBSqlDatabaseResource cosmosDBSqlDatabase = client.GetCosmosDBSqlDatabaseResource(cosmosDBSqlDatabaseResourceId);

// get the collection of this CosmosDBSqlClientEncryptionKeyResource
CosmosDBSqlClientEncryptionKeyCollection collection = cosmosDBSqlDatabase.GetCosmosDBSqlClientEncryptionKeys();

// invoke the operation
string clientEncryptionKeyName = "cekName";
NullableResponse<CosmosDBSqlClientEncryptionKeyResource> response = await collection.GetIfExistsAsync(clientEncryptionKeyName);
CosmosDBSqlClientEncryptionKeyResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CosmosDBSqlClientEncryptionKeyData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
