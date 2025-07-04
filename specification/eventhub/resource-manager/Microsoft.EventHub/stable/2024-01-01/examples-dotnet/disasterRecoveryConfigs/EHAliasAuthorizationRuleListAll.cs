using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/disasterRecoveryConfigs/EHAliasAuthorizationRuleListAll.json
// this example is just showing the usage of "DisasterRecoveryConfigs_ListAuthorizationRules" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsDisasterRecoveryResource created on azure
// for more information of creating EventHubsDisasterRecoveryResource, please refer to the document of EventHubsDisasterRecoveryResource
string subscriptionId = "exampleSubscriptionId";
string resourceGroupName = "exampleResourceGroup";
string namespaceName = "sdk-Namespace-9080";
string @alias = "sdk-DisasterRecovery-4047";
ResourceIdentifier eventHubsDisasterRecoveryResourceId = EventHubsDisasterRecoveryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, @alias);
EventHubsDisasterRecoveryResource eventHubsDisasterRecovery = client.GetEventHubsDisasterRecoveryResource(eventHubsDisasterRecoveryResourceId);

// get the collection of this EventHubsDisasterRecoveryAuthorizationRuleResource
EventHubsDisasterRecoveryAuthorizationRuleCollection collection = eventHubsDisasterRecovery.GetEventHubsDisasterRecoveryAuthorizationRules();

// invoke the operation and iterate over the result
await foreach (EventHubsDisasterRecoveryAuthorizationRuleResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EventHubsAuthorizationRuleData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
