using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;
using Azure.ResourceManager.CosmosDB.Models;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/CosmosDBLocationGet.json
// this example is just showing the usage of "Locations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBLocationResource created on azure
// for more information of creating CosmosDBLocationResource, please refer to the document of CosmosDBLocationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
AzureLocation location = new AzureLocation("westus");
ResourceIdentifier cosmosDBLocationResourceId = CosmosDBLocationResource.CreateResourceIdentifier(subscriptionId, location);
CosmosDBLocationResource cosmosDBLocation = client.GetCosmosDBLocationResource(cosmosDBLocationResourceId);

// invoke the operation
CosmosDBLocationResource result = await cosmosDBLocation.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CosmosDBLocationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
