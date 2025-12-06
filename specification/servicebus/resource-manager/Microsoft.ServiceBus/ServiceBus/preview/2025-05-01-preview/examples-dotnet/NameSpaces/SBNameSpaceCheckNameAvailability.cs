using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ServiceBus.Models;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/ServiceBus/preview/2025-05-01-preview/examples/NameSpaces/SBNameSpaceCheckNameAvailability.json
// this example is just showing the usage of "Namespaces_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
ServiceBusNameAvailabilityContent content = new ServiceBusNameAvailabilityContent("sdk-Namespace-2924");
ServiceBusNameAvailabilityResult result = await subscriptionResource.CheckServiceBusNamespaceNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
