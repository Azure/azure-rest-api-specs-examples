using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.BotService;
using Azure.ResourceManager.BotService.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/botservice/resource-manager/Microsoft.BotService/preview/2022-06-15-preview/examples/ListPrivateLinkResources.json
// this example is just showing the usage of "PrivateLinkResources_ListByBotResource" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
await foreach (BotServicePrivateLinkResource item in bot.GetPrivateLinkResourcesByBotResourceAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
