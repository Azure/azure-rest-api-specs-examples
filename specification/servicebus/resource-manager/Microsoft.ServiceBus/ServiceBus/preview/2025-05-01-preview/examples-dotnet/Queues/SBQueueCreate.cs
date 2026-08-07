using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/ServiceBus/preview/2025-05-01-preview/examples/Queues/SBQueueCreate.json
// this example is just showing the usage of "Queues_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusNamespaceResource created on azure
// for more information of creating ServiceBusNamespaceResource, please refer to the document of ServiceBusNamespaceResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ArunMonocle";
string namespaceName = "sdk-Namespace-3174";
ResourceIdentifier serviceBusNamespaceResourceId = ServiceBusNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName);
ServiceBusNamespaceResource serviceBusNamespace = client.GetServiceBusNamespaceResource(serviceBusNamespaceResourceId);

// get the collection of this ServiceBusQueueResource
ServiceBusQueueCollection collection = serviceBusNamespace.GetServiceBusQueues();

// invoke the operation
string queueName = "sdk-Queues-5647";
ServiceBusQueueData data = new ServiceBusQueueData
{
    EnablePartitioning = true,
};
ArmOperation<ServiceBusQueueResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, queueName, data);
ServiceBusQueueResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceBusQueueData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
