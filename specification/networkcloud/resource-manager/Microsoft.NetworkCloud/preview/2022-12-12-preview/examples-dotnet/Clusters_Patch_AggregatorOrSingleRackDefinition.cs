using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/Clusters_Patch_AggregatorOrSingleRackDefinition.json
// this example is just showing the usage of "Clusters_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ClusterResource created on azure
// for more information of creating ClusterResource, please refer to the document of ClusterResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string clusterName = "clusterName";
ResourceIdentifier clusterResourceId = ClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
ClusterResource cluster = client.GetClusterResource(clusterResourceId);

// invoke the operation
ClusterPatch patch = new ClusterPatch()
{
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2",
    },
    AggregatorOrSingleRackDefinition = new RackDefinition("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkRacks/networkRackName", "newSerialNumber", "/subscriptions/subscriptionId/providers/Microsoft.NetworkCloud/rackSkus/rackSkuName")
    {
        BareMetalMachineConfigurationData =
        {
        new BareMetalMachineConfigurationData(new AdministrativeCredentials("username")
        {
        Password = "{password}",
        },"AA:BB:CC:DD:EE:FF","00:BB:CC:DD:EE:FF",1,"BM1219XXX")
        {
        MachineDetails = "extraDetails",
        MachineName = "bmmName1",
        },new BareMetalMachineConfigurationData(new AdministrativeCredentials("username")
        {
        Password = "{password}",
        },"AA:BB:CC:DD:EE:00","00:BB:CC:DD:EE:00",2,"BM1219YYY")
        {
        MachineDetails = "extraDetails",
        MachineName = "bmmName2",
        }
        },
        RackLocation = "Foo Datacenter, Floor 3, Aisle 9, Rack 2",
        StorageApplianceConfigurationData =
        {
        new StorageApplianceConfigurationData(new AdministrativeCredentials("username")
        {
        Password = "{password}",
        },1,"BM1219XXX")
        {
        StorageApplianceName = "vmName",
        }
        },
    },
    ComputeDeploymentThreshold = new ValidationThreshold(ValidationThresholdGrouping.PerCluster, ValidationThresholdType.PercentSuccess, 90),
};
ArmOperation<ClusterResource> lro = await cluster.UpdateAsync(WaitUntil.Completed, patch);
ClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
