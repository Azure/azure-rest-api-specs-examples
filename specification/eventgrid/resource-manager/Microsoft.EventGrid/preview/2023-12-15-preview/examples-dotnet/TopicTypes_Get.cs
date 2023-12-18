using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventGrid;
using Azure.ResourceManager.EventGrid.Models;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/TopicTypes_Get.json
// this example is just showing the usage of "TopicTypes_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TopicTypeResource created on azure
// for more information of creating TopicTypeResource, please refer to the document of TopicTypeResource
string topicTypeName = "Microsoft.Storage.StorageAccounts";
ResourceIdentifier topicTypeResourceId = TopicTypeResource.CreateResourceIdentifier(topicTypeName);
TopicTypeResource topicType = client.GetTopicTypeResource(topicTypeResourceId);

// invoke the operation
TopicTypeResource result = await topicType.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TopicTypeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
