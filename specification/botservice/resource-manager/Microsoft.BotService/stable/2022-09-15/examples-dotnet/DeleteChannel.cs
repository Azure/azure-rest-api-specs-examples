using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.BotService;
using Azure.ResourceManager.BotService.Models;

// Generated from example definition: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/DeleteChannel.json
// this example is just showing the usage of "Channels_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BotChannelResource created on azure
// for more information of creating BotChannelResource, please refer to the document of BotChannelResource
string subscriptionId = "subscription-id";
string resourceGroupName = "OneResourceGroupName";
string resourceName = "samplebotname";
BotChannelName channelName = BotChannelName.EmailChannel;
ResourceIdentifier botChannelResourceId = BotChannelResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, channelName);
BotChannelResource botChannel = client.GetBotChannelResource(botChannelResourceId);

// invoke the operation
await botChannel.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
