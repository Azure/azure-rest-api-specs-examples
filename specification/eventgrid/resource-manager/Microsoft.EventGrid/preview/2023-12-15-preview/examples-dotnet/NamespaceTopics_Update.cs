using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventGrid;
using Azure.ResourceManager.EventGrid.Models;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/NamespaceTopics_Update.json
// this example is just showing the usage of "NamespaceTopics_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NamespaceTopicResource created on azure
// for more information of creating NamespaceTopicResource, please refer to the document of NamespaceTopicResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string namespaceName = "exampleNamespaceName1";
string topicName = "exampleNamespaceTopicName1";
ResourceIdentifier namespaceTopicResourceId = NamespaceTopicResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, topicName);
NamespaceTopicResource namespaceTopic = client.GetNamespaceTopicResource(namespaceTopicResourceId);

// invoke the operation
NamespaceTopicPatch patch = new NamespaceTopicPatch()
{
    EventRetentionInDays = 1,
};
ArmOperation<NamespaceTopicResource> lro = await namespaceTopic.UpdateAsync(WaitUntil.Completed, patch);
NamespaceTopicResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NamespaceTopicData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
