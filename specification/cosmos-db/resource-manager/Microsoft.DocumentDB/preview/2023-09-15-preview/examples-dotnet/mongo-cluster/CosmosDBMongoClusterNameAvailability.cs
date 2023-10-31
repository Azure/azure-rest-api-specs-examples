using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;
using Azure.ResourceManager.CosmosDB.Models;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/mongo-cluster/CosmosDBMongoClusterNameAvailability.json
// this example is just showing the usage of "MongoClusters_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBLocationResource created on azure
// for more information of creating CosmosDBLocationResource, please refer to the document of CosmosDBLocationResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
AzureLocation location = new AzureLocation("westus2");
ResourceIdentifier cosmosDBLocationResourceId = CosmosDBLocationResource.CreateResourceIdentifier(subscriptionId, location);
CosmosDBLocationResource cosmosDBLocation = client.GetCosmosDBLocationResource(cosmosDBLocationResourceId);

// invoke the operation
CheckCosmosDBNameAvailabilityContent content = new CheckCosmosDBNameAvailabilityContent()
{
    Name = "newmongocluster",
    ResourceType = "Microsoft.DocumentDB/mongoClusters",
};
CheckCosmosDBNameAvailabilityResponse result = await cosmosDBLocation.CheckMongoClusterNameAailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
