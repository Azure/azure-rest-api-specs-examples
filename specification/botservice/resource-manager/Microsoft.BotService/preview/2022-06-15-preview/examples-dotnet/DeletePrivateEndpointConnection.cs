using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.BotService;
using Azure.ResourceManager.BotService.Models;

// Generated from example definition: specification/botservice/resource-manager/Microsoft.BotService/preview/2022-06-15-preview/examples/DeletePrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BotServicePrivateEndpointConnectionResource created on azure
// for more information of creating BotServicePrivateEndpointConnectionResource, please refer to the document of BotServicePrivateEndpointConnectionResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res6977";
string resourceName = "sto2527";
string privateEndpointConnectionName = "{privateEndpointConnectionName}";
ResourceIdentifier botServicePrivateEndpointConnectionResourceId = BotServicePrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, privateEndpointConnectionName);
BotServicePrivateEndpointConnectionResource botServicePrivateEndpointConnection = client.GetBotServicePrivateEndpointConnectionResource(botServicePrivateEndpointConnectionResourceId);

// invoke the operation
await botServicePrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
