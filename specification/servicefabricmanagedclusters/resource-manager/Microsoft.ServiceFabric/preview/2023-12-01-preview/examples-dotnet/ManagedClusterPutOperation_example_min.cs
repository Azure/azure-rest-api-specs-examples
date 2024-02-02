using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ServiceFabricManagedClusters;
using Azure.ResourceManager.ServiceFabricManagedClusters.Models;

// Generated from example definition: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2023-12-01-preview/examples/ManagedClusterPutOperation_example_min.json
// this example is just showing the usage of "ManagedClusters_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ServiceFabricManagedClusterResource
ServiceFabricManagedClusterCollection collection = resourceGroupResource.GetServiceFabricManagedClusters();

// invoke the operation
string clusterName = "myCluster";
ServiceFabricManagedClusterData data = new ServiceFabricManagedClusterData(new AzureLocation("eastus"), new ServiceFabricManagedClustersSku(ServiceFabricManagedClustersSkuName.Basic))
{
    DnsName = "myCluster",
    AdminUserName = "vmadmin",
    AdminPassword = "{vm-password}",
    FabricSettings =
    {
    new ClusterFabricSettingsSection("ManagedIdentityTokenService",new ClusterFabricSettingsParameterDescription[]
    {
    new ClusterFabricSettingsParameterDescription("IsEnabled","true")
    })
    },
    ClusterUpgradeMode = ManagedClusterUpgradeMode.Automatic,
    ClusterUpgradeCadence = ManagedClusterUpgradeCadence.Wave1,
};
ArmOperation<ServiceFabricManagedClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, clusterName, data);
ServiceFabricManagedClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceFabricManagedClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
