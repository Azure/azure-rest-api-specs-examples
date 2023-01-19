using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventHubs;
using Azure.ResourceManager.EventHubs.Models;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-01-01-preview/examples/disasterRecoveryConfigs/EHAliasAuthorizationRuleListKey.json
// this example is just showing the usage of "DisasterRecoveryConfigs_ListKeys" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsDisasterRecoveryAuthorizationRuleResource created on azure
// for more information of creating EventHubsDisasterRecoveryAuthorizationRuleResource, please refer to the document of EventHubsDisasterRecoveryAuthorizationRuleResource
string subscriptionId = "exampleSubscriptionId";
string resourceGroupName = "exampleResourceGroup";
string namespaceName = "sdk-Namespace-2702";
string @alias = "sdk-DisasterRecovery-4047";
string authorizationRuleName = "sdk-Authrules-1746";
ResourceIdentifier eventHubsDisasterRecoveryAuthorizationRuleResourceId = EventHubsDisasterRecoveryAuthorizationRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, @alias, authorizationRuleName);
EventHubsDisasterRecoveryAuthorizationRuleResource eventHubsDisasterRecoveryAuthorizationRule = client.GetEventHubsDisasterRecoveryAuthorizationRuleResource(eventHubsDisasterRecoveryAuthorizationRuleResourceId);

// invoke the operation
EventHubsAccessKeys result = await eventHubsDisasterRecoveryAuthorizationRule.GetKeysAsync();

Console.WriteLine($"Succeeded: {result}");
