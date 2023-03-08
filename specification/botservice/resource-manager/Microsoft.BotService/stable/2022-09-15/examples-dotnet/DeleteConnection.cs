using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.BotService;
using Azure.ResourceManager.BotService.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/DeleteConnection.json
// this example is just showing the usage of "BotConnection_Delete" operation, for the dependent resources, they will have to be created separately.

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
await botConnectionSetting.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
