using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.BotService.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.BotService;

// Generated from example definition: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/ListQnAMakerEndpointKeys.json
// this example is just showing the usage of "QnAMakerEndpointKeys_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "subscription-id";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
GetBotServiceQnAMakerEndpointKeyContent content = new GetBotServiceQnAMakerEndpointKeyContent
{
    Hostname = "https://xxx.cognitiveservices.azure.com/",
    Authkey = "testAuthKey",
};
GetBotServiceQnAMakerEndpointKeyResult result = await subscriptionResource.GetBotServiceQnAMakerEndpointKeyAsync(content);

Console.WriteLine($"Succeeded: {result}");
