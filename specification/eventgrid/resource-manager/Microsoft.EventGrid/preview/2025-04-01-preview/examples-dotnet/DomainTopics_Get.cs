using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/DomainTopics_Get.json
// this example is just showing the usage of "DomainTopics_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DomainTopicResource created on azure
// for more information of creating DomainTopicResource, please refer to the document of DomainTopicResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string domainName = "exampledomain2";
string domainTopicName = "topic1";
ResourceIdentifier domainTopicResourceId = DomainTopicResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, domainName, domainTopicName);
DomainTopicResource domainTopic = client.GetDomainTopicResource(domainTopicResourceId);

// invoke the operation
DomainTopicResource result = await domainTopic.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DomainTopicData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
