using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/Eventhub/preview/2025-05-01-preview/examples/NameSpaces/EHNameSpaceUpdate.json
// this example is just showing the usage of "Namespaces_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsNamespaceResource created on azure
// for more information of creating EventHubsNamespaceResource, please refer to the document of EventHubsNamespaceResource
string subscriptionId = "SampleSubscription";
string resourceGroupName = "ResurceGroupSample";
string namespaceName = "NamespaceSample";
ResourceIdentifier eventHubsNamespaceResourceId = EventHubsNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName);
EventHubsNamespaceResource eventHubsNamespace = client.GetEventHubsNamespaceResource(eventHubsNamespaceResourceId);

// invoke the operation
EventHubsNamespaceData data = new EventHubsNamespaceData(new AzureLocation("East US"))
{
    Identity = new ManagedServiceIdentity("SystemAssigned, UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud2")] = new UserAssignedIdentity()
        },
    },
};
ArmOperation<EventHubsNamespaceResource> lro = await eventHubsNamespace.UpdateAsync(WaitUntil.Completed, data);
EventHubsNamespaceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventHubsNamespaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
