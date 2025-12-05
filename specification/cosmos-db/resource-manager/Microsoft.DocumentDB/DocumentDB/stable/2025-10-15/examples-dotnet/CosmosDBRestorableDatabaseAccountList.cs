using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/CosmosDBRestorableDatabaseAccountList.json
// this example is just showing the usage of "RestorableDatabaseAccounts_ListByLocation" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
await foreach (RestorableCosmosDBAccountResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    RestorableCosmosDBAccountData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
