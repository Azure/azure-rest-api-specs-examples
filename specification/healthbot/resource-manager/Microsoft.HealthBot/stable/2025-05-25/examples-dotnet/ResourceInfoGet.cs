using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HealthBot.Models;
using Azure.ResourceManager.HealthBot;

// Generated from example definition: specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2025-05-25/examples/ResourceInfoGet.json
// this example is just showing the usage of "Bots_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthBotResource created on azure
// for more information of creating HealthBotResource, please refer to the document of HealthBotResource
string subscriptionId = "subid";
string resourceGroupName = "healthbotClient";
string botName = "samplebotname";
ResourceIdentifier healthBotResourceId = HealthBotResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, botName);
HealthBotResource healthBot = client.GetHealthBotResource(healthBotResourceId);

// invoke the operation
HealthBotResource result = await healthBot.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HealthBotData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
