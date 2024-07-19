using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ServiceBus.Models;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasCheckNameAvailability.json
// this example is just showing the usage of "DisasterRecoveryConfigs_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusNamespaceResource created on azure
// for more information of creating ServiceBusNamespaceResource, please refer to the document of ServiceBusNamespaceResource
string subscriptionId = "exampleSubscriptionId";
string resourceGroupName = "exampleResourceGroup";
string namespaceName = "sdk-Namespace-9080";
ResourceIdentifier serviceBusNamespaceResourceId = ServiceBusNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName);
ServiceBusNamespaceResource serviceBusNamespace = client.GetServiceBusNamespaceResource(serviceBusNamespaceResourceId);

// invoke the operation
ServiceBusNameAvailabilityContent content = new ServiceBusNameAvailabilityContent("sdk-DisasterRecovery-9474");
ServiceBusNameAvailabilityResult result = await serviceBusNamespace.CheckServiceBusDisasterRecoveryNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
