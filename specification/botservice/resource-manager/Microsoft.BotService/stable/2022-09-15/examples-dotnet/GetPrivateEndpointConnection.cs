using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.BotService;
using Azure.ResourceManager.BotService.Models;

// Generated from example definition: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/GetPrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BotResource created on azure
// for more information of creating BotResource, please refer to the document of BotResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res6977";
string resourceName = "sto2527";
ResourceIdentifier botResourceId = BotResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
BotResource bot = client.GetBotResource(botResourceId);

// get the collection of this BotServicePrivateEndpointConnectionResource
BotServicePrivateEndpointConnectionCollection collection = bot.GetBotServicePrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "{privateEndpointConnectionName}";
bool result = await collection.ExistsAsync(privateEndpointConnectionName);

Console.WriteLine($"Succeeded: {result}");
