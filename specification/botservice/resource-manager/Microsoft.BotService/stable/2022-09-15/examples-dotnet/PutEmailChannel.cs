using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.BotService.Models;
using Azure.ResourceManager.BotService;

// Generated from example definition: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/PutEmailChannel.json
// this example is just showing the usage of "Channels_Create" operation, for the dependent resources, they will have to be created separately.

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
BotChannelName channelName = BotChannelName.EmailChannel;
BotChannelData data = new BotChannelData(new AzureLocation("global"))
{
    Properties = new EmailChannel
    {
        Properties = new EmailChannelProperties("a@b.com", true)
        {
            AuthMethod = EmailChannelAuthMethod.Graph,
            MagicCode = "000000",
        },
    },
};
ArmOperation<BotChannelResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, channelName, data);
BotChannelResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BotChannelData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
