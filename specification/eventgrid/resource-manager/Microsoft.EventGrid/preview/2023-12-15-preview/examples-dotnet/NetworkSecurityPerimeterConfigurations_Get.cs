using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/NetworkSecurityPerimeterConfigurations_Get.json
// this example is just showing the usage of "NetworkSecurityPerimeterConfigurations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventGridTopicResource created on azure
// for more information of creating EventGridTopicResource, please refer to the document of EventGridTopicResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string resourceName = "exampleResourceName";
ResourceIdentifier eventGridTopicResourceId = EventGridTopicResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
EventGridTopicResource eventGridTopic = client.GetEventGridTopicResource(eventGridTopicResourceId);

// get the collection of this TopicNetworkSecurityPerimeterConfigurationResource
TopicNetworkSecurityPerimeterConfigurationCollection collection = eventGridTopic.GetTopicNetworkSecurityPerimeterConfigurations();

// invoke the operation
string perimeterGuid = "8f6b6269-84f2-4d09-9e31-1127efcd1e40perimeter";
string associationName = "someAssociation";
NullableResponse<TopicNetworkSecurityPerimeterConfigurationResource> response = await collection.GetIfExistsAsync(perimeterGuid, associationName);
TopicNetworkSecurityPerimeterConfigurationResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetworkSecurityPerimeterConfigurationData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
