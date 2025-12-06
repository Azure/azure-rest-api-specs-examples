using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/CosmosDBRestorableDatabaseAccountGet.json
// this example is just showing the usage of "RestorableDatabaseAccounts_GetByLocation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBLocationResource created on azure
// for more information of creating CosmosDBLocationResource, please refer to the document of CosmosDBLocationResource
string subscriptionId = "subid";
AzureLocation location = new AzureLocation("West US");
ResourceIdentifier cosmosDBLocationResourceId = CosmosDBLocationResource.CreateResourceIdentifier(subscriptionId, location);
CosmosDBLocationResource cosmosDBLocation = client.GetCosmosDBLocationResource(cosmosDBLocationResourceId);

// get the collection of this RestorableCosmosDBAccountResource
RestorableCosmosDBAccountCollection collection = cosmosDBLocation.GetRestorableCosmosDBAccounts();

// invoke the operation
Guid instanceId = Guid.Parse("d9b26648-2f53-4541-b3d8-3044f4f9810d");
NullableResponse<RestorableCosmosDBAccountResource> response = await collection.GetIfExistsAsync(instanceId);
RestorableCosmosDBAccountResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    RestorableCosmosDBAccountData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
