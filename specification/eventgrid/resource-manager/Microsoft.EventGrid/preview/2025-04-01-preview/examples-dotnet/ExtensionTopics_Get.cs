using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/ExtensionTopics_Get.json
// this example is just showing the usage of "ExtensionTopics_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExtensionTopicResource created on azure
// for more information of creating ExtensionTopicResource, please refer to the document of ExtensionTopicResource
string scope = "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/microsoft.storage/storageaccounts/exampleResourceName/providers/Microsoft.eventgrid/extensionTopics/default";
ResourceIdentifier extensionTopicResourceId = ExtensionTopicResource.CreateResourceIdentifier(scope);
ExtensionTopicResource extensionTopic = client.GetExtensionTopicResource(extensionTopicResourceId);

// invoke the operation
ExtensionTopicResource result = await extensionTopic.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ExtensionTopicData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
