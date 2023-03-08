using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.BotService;
using Azure.ResourceManager.BotService.Models;

// Generated from example definition: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/PutPrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BotResource created on azure
// for more information of creating BotResource, please refer to the document of BotResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res7687";
string resourceName = "sto9699";
ResourceIdentifier botResourceId = BotResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
BotResource bot = client.GetBotResource(botResourceId);

// get the collection of this BotServicePrivateEndpointConnectionResource
BotServicePrivateEndpointConnectionCollection collection = bot.GetBotServicePrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "{privateEndpointConnectionName}";
BotServicePrivateEndpointConnectionData data = new BotServicePrivateEndpointConnectionData()
{
    ConnectionState = new BotServicePrivateLinkServiceConnectionState()
    {
        Status = BotServicePrivateEndpointServiceConnectionStatus.Approved,
        Description = "Auto-Approved",
    },
};
ArmOperation<BotServicePrivateEndpointConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, privateEndpointConnectionName, data);
BotServicePrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BotServicePrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
