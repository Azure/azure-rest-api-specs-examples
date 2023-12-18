using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventGrid;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/Namespaces_CreateOrUpdate.json
// this example is just showing the usage of "Namespaces_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this EventGridNamespaceResource
EventGridNamespaceCollection collection = resourceGroupResource.GetEventGridNamespaces();

// invoke the operation
string namespaceName = "exampleNamespaceName1";
EventGridNamespaceData data = new EventGridNamespaceData(new AzureLocation("westus"))
{
    TopicSpacesConfiguration = new TopicSpacesConfiguration()
    {
        State = TopicSpacesConfigurationState.Enabled,
        RouteTopicResourceId = "/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampleTopic1",
    },
    Tags =
    {
    ["tag1"] = "value11",
    ["tag2"] = "value22",
    },
};
ArmOperation<EventGridNamespaceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, namespaceName, data);
EventGridNamespaceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventGridNamespaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
