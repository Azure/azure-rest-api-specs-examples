using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Subscriptions/SBSubscriptionDelete.json
// this example is just showing the usage of "Subscriptions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusSubscriptionResource created on azure
// for more information of creating ServiceBusSubscriptionResource, please refer to the document of ServiceBusSubscriptionResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "ResourceGroup";
string namespaceName = "sdk-Namespace-5882";
string topicName = "sdk-Topics-1804";
string subscriptionName = "sdk-Subscriptions-3670";
ResourceIdentifier serviceBusSubscriptionResourceId = ServiceBusSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, topicName, subscriptionName);
ServiceBusSubscriptionResource serviceBusSubscription = client.GetServiceBusSubscriptionResource(serviceBusSubscriptionResourceId);

// invoke the operation
await serviceBusSubscription.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
