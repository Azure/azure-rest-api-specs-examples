using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventHubs;
using Azure.ResourceManager.EventHubs.Models;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-01-01-preview/examples/EventHubs/EHEventHubAuthorizationRuleRegenerateKey.json
// this example is just showing the usage of "EventHubs_RegenerateKeys" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubAuthorizationRuleResource created on azure
// for more information of creating EventHubAuthorizationRuleResource, please refer to the document of EventHubAuthorizationRuleResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ArunMonocle";
string namespaceName = "sdk-namespace-960";
string eventHubName = "sdk-EventHub-532";
string authorizationRuleName = "sdk-Authrules-1534";
ResourceIdentifier eventHubAuthorizationRuleResourceId = EventHubAuthorizationRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, eventHubName, authorizationRuleName);
EventHubAuthorizationRuleResource eventHubAuthorizationRule = client.GetEventHubAuthorizationRuleResource(eventHubAuthorizationRuleResourceId);

// invoke the operation
EventHubsRegenerateAccessKeyContent content = new EventHubsRegenerateAccessKeyContent(EventHubsAccessKeyType.PrimaryKey);
EventHubsAccessKeys result = await eventHubAuthorizationRule.RegenerateKeysAsync(content);

Console.WriteLine($"Succeeded: {result}");
