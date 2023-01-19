using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.BotService;
using Azure.ResourceManager.BotService.Models;

// Generated from example definition: specification/botservice/resource-manager/Microsoft.BotService/preview/2022-06-15-preview/examples/GetAlexaChannel.json
// this example is just showing the usage of "Channels_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BotResource created on azure
// for more information of creating BotResource, please refer to the document of BotResource
string subscriptionId = "subscription-id";
string resourceGroupName = "OneResourceGroupName";
string resourceName = "samplebotname";
ResourceIdentifier botResourceId = BotResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
BotResource bot = client.GetBotResource(botResourceId);

// get the collection of this BotChannelResource
BotChannelCollection collection = bot.GetBotChannels();

// invoke the operation
ChannelName channelName = ChannelName.AlexaChannel;
bool result = await collection.ExistsAsync(channelName);

Console.WriteLine($"Succeeded: {result}");
