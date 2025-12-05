using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/ServiceBus/preview/2025-05-01-preview/examples/Queues/SBQueueDelete.json
// this example is just showing the usage of "Queues_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusQueueResource created on azure
// for more information of creating ServiceBusQueueResource, please refer to the document of ServiceBusQueueResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ArunMonocle";
string namespaceName = "sdk-Namespace-183";
string queueName = "sdk-Queues-8708";
ResourceIdentifier serviceBusQueueResourceId = ServiceBusQueueResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, queueName);
ServiceBusQueueResource serviceBusQueue = client.GetServiceBusQueueResource(serviceBusQueueResourceId);

// invoke the operation
await serviceBusQueue.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
