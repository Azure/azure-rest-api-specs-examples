using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasAuthorizationRuleGet.json
// this example is just showing the usage of "DisasterRecoveryAuthorizationRules_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusDisasterRecoveryResource created on azure
// for more information of creating ServiceBusDisasterRecoveryResource, please refer to the document of ServiceBusDisasterRecoveryResource
string subscriptionId = "exampleSubscriptionId";
string resourceGroupName = "exampleResourceGroup";
string namespaceName = "sdk-Namespace-9080";
string @alias = "sdk-DisasterRecovery-4879";
ResourceIdentifier serviceBusDisasterRecoveryResourceId = ServiceBusDisasterRecoveryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, @alias);
ServiceBusDisasterRecoveryResource serviceBusDisasterRecovery = client.GetServiceBusDisasterRecoveryResource(serviceBusDisasterRecoveryResourceId);

// get the collection of this ServiceBusDisasterRecoveryAuthorizationRuleResource
ServiceBusDisasterRecoveryAuthorizationRuleCollection collection = serviceBusDisasterRecovery.GetServiceBusDisasterRecoveryAuthorizationRules();

// invoke the operation
string authorizationRuleName = "sdk-Authrules-4879";
bool result = await collection.ExistsAsync(authorizationRuleName);

Console.WriteLine($"Succeeded: {result}");
