using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs.Models;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/NameSpaces/EHNameSpaceAuthorizationRuleDelete.json
// this example is just showing the usage of "Namespaces_DeleteAuthorizationRule" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsNamespaceAuthorizationRuleResource created on azure
// for more information of creating EventHubsNamespaceAuthorizationRuleResource, please refer to the document of EventHubsNamespaceAuthorizationRuleResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ArunMonocle";
string namespaceName = "sdk-Namespace-8980";
string authorizationRuleName = "sdk-Authrules-8929";
ResourceIdentifier eventHubsNamespaceAuthorizationRuleResourceId = EventHubsNamespaceAuthorizationRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, authorizationRuleName);
EventHubsNamespaceAuthorizationRuleResource eventHubsNamespaceAuthorizationRule = client.GetEventHubsNamespaceAuthorizationRuleResource(eventHubsNamespaceAuthorizationRuleResourceId);

// invoke the operation
await eventHubsNamespaceAuthorizationRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
