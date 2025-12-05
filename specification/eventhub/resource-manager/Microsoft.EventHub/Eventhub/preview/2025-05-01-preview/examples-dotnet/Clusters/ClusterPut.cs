using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/Eventhub/preview/2025-05-01-preview/examples/Clusters/ClusterPut.json
// this example is just showing the usage of "Clusters_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this EventHubsClusterResource
EventHubsClusterCollection collection = resourceGroupResource.GetEventHubsClusters();

// invoke the operation
string clusterName = "testCluster";
EventHubsClusterData data = new EventHubsClusterData(new AzureLocation("South Central US"))
{
    Sku = new EventHubsClusterSku(EventHubsClusterSkuName.Dedicated)
    {
        Capacity = 1,
    },
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2"
    },
};
ArmOperation<EventHubsClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, clusterName, data);
EventHubsClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventHubsClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
