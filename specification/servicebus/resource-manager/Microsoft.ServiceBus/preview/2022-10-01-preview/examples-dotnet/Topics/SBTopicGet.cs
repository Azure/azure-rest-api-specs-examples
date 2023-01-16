using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Topics/SBTopicGet.json
// this example is just showing the usage of "Topics_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusTopicResource created on azure
// for more information of creating ServiceBusTopicResource, please refer to the document of ServiceBusTopicResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ArunMonocle";
string namespaceName = "sdk-Namespace-1617";
string topicName = "sdk-Topics-5488";
ResourceIdentifier serviceBusTopicResourceId = ServiceBusTopicResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, topicName);
ServiceBusTopicResource serviceBusTopic = client.GetServiceBusTopicResource(serviceBusTopicResourceId);

// invoke the operation
ServiceBusTopicResource result = await serviceBusTopic.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceBusTopicData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
