using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/disasterRecoveryConfigs/EHAliasGet.json
// this example is just showing the usage of "DisasterRecoveryConfigs_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsDisasterRecoveryResource created on azure
// for more information of creating EventHubsDisasterRecoveryResource, please refer to the document of EventHubsDisasterRecoveryResource
string subscriptionId = "exampleSubscriptionId";
string resourceGroupName = "exampleResourceGroup";
string namespaceName = "sdk-Namespace-8859";
string @alias = "sdk-DisasterRecovery-3814";
ResourceIdentifier eventHubsDisasterRecoveryResourceId = EventHubsDisasterRecoveryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, @alias);
EventHubsDisasterRecoveryResource eventHubsDisasterRecovery = client.GetEventHubsDisasterRecoveryResource(eventHubsDisasterRecoveryResourceId);

// invoke the operation
EventHubsDisasterRecoveryResource result = await eventHubsDisasterRecovery.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventHubsDisasterRecoveryData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
