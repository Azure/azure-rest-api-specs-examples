using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/CosmosDBRestorableMongodbDatabaseList.json
// this example is just showing the usage of "RestorableMongodbDatabases_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RestorableCosmosDBAccountResource created on azure
// for more information of creating RestorableCosmosDBAccountResource, please refer to the document of RestorableCosmosDBAccountResource
string subscriptionId = "2296c272-5d55-40d9-bc05-4d56dc2d7588";
AzureLocation location = new AzureLocation("WestUS");
Guid instanceId = Guid.Parse("d9b26648-2f53-4541-b3d8-3044f4f9810d");
ResourceIdentifier restorableCosmosDBAccountResourceId = RestorableCosmosDBAccountResource.CreateResourceIdentifier(subscriptionId, location, instanceId);
RestorableCosmosDBAccountResource restorableCosmosDBAccount = client.GetRestorableCosmosDBAccountResource(restorableCosmosDBAccountResourceId);

// invoke the operation and iterate over the result
await foreach (RestorableMongoDBDatabase item in restorableCosmosDBAccount.GetRestorableMongoDBDatabasesAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
