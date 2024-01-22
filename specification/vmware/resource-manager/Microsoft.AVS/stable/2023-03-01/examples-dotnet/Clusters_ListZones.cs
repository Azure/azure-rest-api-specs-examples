using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Avs;
using Azure.ResourceManager.Avs.Models;

// Generated from example definition: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/Clusters_ListZones.json
// this example is just showing the usage of "Clusters_ListZones" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AvsPrivateCloudClusterResource created on azure
// for more information of creating AvsPrivateCloudClusterResource, please refer to the document of AvsPrivateCloudClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
string clusterName = "cluster1";
ResourceIdentifier avsPrivateCloudClusterResourceId = AvsPrivateCloudClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName, clusterName);
AvsPrivateCloudClusterResource avsPrivateCloudCluster = client.GetAvsPrivateCloudClusterResource(avsPrivateCloudClusterResourceId);

// invoke the operation and iterate over the result
await foreach (AvsClusterZone item in avsPrivateCloudCluster.GetZonesAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
