using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-06-01-preview/examples/ClusterMetricsConfigurations_Create.json
// this example is just showing the usage of "MetricsConfigurations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkCloudClusterResource created on azure
// for more information of creating NetworkCloudClusterResource, please refer to the document of NetworkCloudClusterResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
string clusterName = "clusterName";
ResourceIdentifier networkCloudClusterResourceId = NetworkCloudClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
NetworkCloudClusterResource networkCloudCluster = client.GetNetworkCloudClusterResource(networkCloudClusterResourceId);

// get the collection of this NetworkCloudClusterMetricsConfigurationResource
NetworkCloudClusterMetricsConfigurationCollection collection = networkCloudCluster.GetNetworkCloudClusterMetricsConfigurations();

// invoke the operation
string metricsConfigurationName = "default";
NetworkCloudClusterMetricsConfigurationData data = new NetworkCloudClusterMetricsConfigurationData(new AzureLocation("location"), new ExtendedLocation("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName", "CustomLocation"), 15L)
{
    EnabledMetrics = { "metric1", "metric2" },
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2"
    },
};
ArmOperation<NetworkCloudClusterMetricsConfigurationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, metricsConfigurationName, data);
NetworkCloudClusterMetricsConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkCloudClusterMetricsConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
