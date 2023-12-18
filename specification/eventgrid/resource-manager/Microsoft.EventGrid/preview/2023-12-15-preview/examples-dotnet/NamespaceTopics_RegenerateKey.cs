using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventGrid;
using Azure.ResourceManager.EventGrid.Models;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/NamespaceTopics_RegenerateKey.json
// this example is just showing the usage of "NamespaceTopics_RegenerateKey" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NamespaceTopicResource created on azure
// for more information of creating NamespaceTopicResource, please refer to the document of NamespaceTopicResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string namespaceName = "examplenamespace2";
string topicName = "examplenamespacetopic2";
ResourceIdentifier namespaceTopicResourceId = NamespaceTopicResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, topicName);
NamespaceTopicResource namespaceTopic = client.GetNamespaceTopicResource(namespaceTopicResourceId);

// invoke the operation
TopicRegenerateKeyContent content = new TopicRegenerateKeyContent("key1");
ArmOperation<TopicSharedAccessKeys> lro = await namespaceTopic.RegenerateKeyAsync(WaitUntil.Completed, content);
TopicSharedAccessKeys result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
