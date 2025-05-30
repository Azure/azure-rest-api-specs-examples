using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.BotService.Models;
using Azure.ResourceManager.BotService;

// Generated from example definition: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/UpdateConnection.json
// this example is just showing the usage of "BotConnection_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BotConnectionSettingResource created on azure
// for more information of creating BotConnectionSettingResource, please refer to the document of BotConnectionSettingResource
string subscriptionId = "subscription-id";
string resourceGroupName = "OneResourceGroupName";
string resourceName = "samplebotname";
string connectionName = "sampleConnection";
ResourceIdentifier botConnectionSettingResourceId = BotConnectionSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, connectionName);
BotConnectionSettingResource botConnectionSetting = client.GetBotConnectionSettingResource(botConnectionSettingResourceId);

// invoke the operation
BotConnectionSettingData data = new BotConnectionSettingData(new AzureLocation("global"))
{
    Properties = new BotConnectionSettingProperties
    {
        ClientId = "sampleclientid",
        ClientSecret = "samplesecret",
        Scopes = "samplescope",
        ServiceProviderId = "serviceproviderid",
        ServiceProviderDisplayName = "serviceProviderDisplayName",
        Parameters = {new BotConnectionSettingParameter
        {
        Key = "key1",
        Value = "value1",
        }, new BotConnectionSettingParameter
        {
        Key = "key2",
        Value = "value2",
        }},
    },
    ETag = new ETag("etag1"),
};
BotConnectionSettingResource result = await botConnectionSetting.UpdateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BotConnectionSettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
