using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/Eventhub/preview/2025-05-01-preview/examples/Clusters/ClusterPatch.json
// this example is just showing the usage of "Clusters_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsClusterResource created on azure
// for more information of creating EventHubsClusterResource, please refer to the document of EventHubsClusterResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "myResourceGroup";
string clusterName = "testCluster";
ResourceIdentifier eventHubsClusterResourceId = EventHubsClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
EventHubsClusterResource eventHubsCluster = client.GetEventHubsClusterResource(eventHubsClusterResourceId);

// invoke the operation
EventHubsClusterData data = new EventHubsClusterData(new AzureLocation("South Central US"))
{
    Tags =
    {
    ["tag3"] = "value3",
    ["tag4"] = "value4"
    },
};
ArmOperation<EventHubsClusterResource> lro = await eventHubsCluster.UpdateAsync(WaitUntil.Completed, data);
EventHubsClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventHubsClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
