using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs.Models;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/ApplicationGroup/ApplicationGroupGet.json
// this example is just showing the usage of "ApplicationGroup_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsNamespaceResource created on azure
// for more information of creating EventHubsNamespaceResource, please refer to the document of EventHubsNamespaceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contosotest";
string namespaceName = "contoso-ua-test-eh-system-1";
ResourceIdentifier eventHubsNamespaceResourceId = EventHubsNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName);
EventHubsNamespaceResource eventHubsNamespace = client.GetEventHubsNamespaceResource(eventHubsNamespaceResourceId);

// get the collection of this EventHubsApplicationGroupResource
EventHubsApplicationGroupCollection collection = eventHubsNamespace.GetEventHubsApplicationGroups();

// invoke the operation
string applicationGroupName = "appGroup1";
NullableResponse<EventHubsApplicationGroupResource> response = await collection.GetIfExistsAsync(applicationGroupName);
EventHubsApplicationGroupResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EventHubsApplicationGroupData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
