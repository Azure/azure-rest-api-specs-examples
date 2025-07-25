using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/ClientGroups_CreateOrUpdate.json
// this example is just showing the usage of "ClientGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventGridNamespaceClientGroupResource created on azure
// for more information of creating EventGridNamespaceClientGroupResource, please refer to the document of EventGridNamespaceClientGroupResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string namespaceName = "exampleNamespaceName1";
string clientGroupName = "exampleClientGroupName1";
ResourceIdentifier eventGridNamespaceClientGroupResourceId = EventGridNamespaceClientGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, clientGroupName);
EventGridNamespaceClientGroupResource eventGridNamespaceClientGroup = client.GetEventGridNamespaceClientGroupResource(eventGridNamespaceClientGroupResourceId);

// invoke the operation
EventGridNamespaceClientGroupData data = new EventGridNamespaceClientGroupData
{
    Description = "This is a test client group",
    Query = "attributes.b IN ['a', 'b', 'c']",
};
ArmOperation<EventGridNamespaceClientGroupResource> lro = await eventGridNamespaceClientGroup.UpdateAsync(WaitUntil.Completed, data);
EventGridNamespaceClientGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventGridNamespaceClientGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
