using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/Eventhub/preview/2025-05-01-preview/examples/disasterRecoveryConfigs/EHAliasCheckNameAvailability.json
// this example is just showing the usage of "DisasterRecoveryConfigs_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsNamespaceResource created on azure
// for more information of creating EventHubsNamespaceResource, please refer to the document of EventHubsNamespaceResource
string subscriptionId = "exampleSubscriptionId";
string resourceGroupName = "exampleResourceGroup";
string namespaceName = "sdk-Namespace-9080";
ResourceIdentifier eventHubsNamespaceResourceId = EventHubsNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName);
EventHubsNamespaceResource eventHubsNamespace = client.GetEventHubsNamespaceResource(eventHubsNamespaceResourceId);

// invoke the operation
EventHubsNameAvailabilityContent content = new EventHubsNameAvailabilityContent("sdk-DisasterRecovery-9474");
EventHubsNameAvailabilityResult result = await eventHubsNamespace.CheckEventHubsDisasterRecoveryNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
