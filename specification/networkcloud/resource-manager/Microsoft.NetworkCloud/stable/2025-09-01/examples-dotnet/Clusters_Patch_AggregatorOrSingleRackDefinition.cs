using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/Clusters_Patch_AggregatorOrSingleRackDefinition.json
// this example is just showing the usage of "Clusters_Update" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
NetworkCloudClusterPatch patch = new NetworkCloudClusterPatch
{
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2"
    },
    AggregatorOrSingleRackDefinition = new NetworkCloudRackDefinition(new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkRacks/networkRackName"), "newSerialNumber", new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/providers/Microsoft.NetworkCloud/rackSkus/rackSkuName"))
    {
        BareMetalMachineConfigurationData = {new BareMetalMachineConfiguration(new AdministrativeCredentials("username")
        {
        Password = "{password}",
        }, "AA:BB:CC:DD:EE:FF", "00:BB:CC:DD:EE:FF", 1L, "BM1219XXX")
        {
        MachineDetails = "extraDetails",
        MachineName = "bmmName1",
        }, new BareMetalMachineConfiguration(new AdministrativeCredentials("username")
        {
        Password = "{password}",
        }, "AA:BB:CC:DD:EE:00", "00:BB:CC:DD:EE:00", 2L, "BM1219YYY")
        {
        MachineDetails = "extraDetails",
        MachineName = "bmmName2",
        }},
        RackLocation = "Foo Datacenter, Floor 3, Aisle 9, Rack 2",
        StorageApplianceConfigurationData = {new StorageApplianceConfiguration(new AdministrativeCredentials("username")
        {
        Password = "{password}",
        }, 1L, "BM1219XXX")
        {
        StorageApplianceName = "vmName",
        }},
    },
    ComputeDeploymentThreshold = new ValidationThreshold(ValidationThresholdGrouping.PerCluster, ValidationThresholdType.PercentSuccess, 90L),
};
ArmOperation<NetworkCloudClusterResource> lro = await networkCloudCluster.UpdateAsync(WaitUntil.Completed, patch);
NetworkCloudClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkCloudClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
