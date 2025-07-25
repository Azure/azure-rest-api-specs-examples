using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/KubernetesClusters_L2LoadBalancer_Create.json
// this example is just showing the usage of "KubernetesClusters_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this NetworkCloudKubernetesClusterResource
NetworkCloudKubernetesClusterCollection collection = resourceGroupResource.GetNetworkCloudKubernetesClusters();

// invoke the operation
string kubernetesClusterName = "kubernetesClusterName";
NetworkCloudKubernetesClusterData data = new NetworkCloudKubernetesClusterData(
    new AzureLocation("location"),
    new ExtendedLocation("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName", "CustomLocation"),
    new ControlPlaneNodeConfiguration(3L, "NC_G6_28_v1")
    {
        AdministratorConfiguration = new AdministratorConfiguration
        {
            AdminUsername = "azure",
            SshPublicKeys = { new NetworkCloudSshPublicKey("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm") },
        },
        AvailabilityZones = { "1", "2", "3" },
    },
    new InitialAgentPoolConfiguration[]
{
new InitialAgentPoolConfiguration(3L, NetworkCloudAgentPoolMode.System, "SystemPool-1", "NC_P46_224_v1")
{
AdministratorConfiguration = new AdministratorConfiguration
{
AdminUsername = "azure",
SshPublicKeys = {new NetworkCloudSshPublicKey("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm")},
},
AgentOptions = new NetworkCloudAgentConfiguration(96L)
{
HugepagesSize = HugepagesSize.OneG,
},
AttachedNetworkConfiguration = new AttachedNetworkConfiguration
{
L2Networks = {new L2NetworkAttachmentConfiguration(new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l2Networks/l2NetworkName"))
{
PluginType = KubernetesPluginType.Dpdk,
}},
L3Networks = {new L3NetworkAttachmentConfiguration(new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"))
{
IpamEnabled = L3NetworkConfigurationIpamEnabled.False,
PluginType = KubernetesPluginType.Sriov,
}},
TrunkedNetworks = {new TrunkedNetworkAttachmentConfiguration(new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/trunkedNetworks/trunkedNetworkName"))
{
PluginType = KubernetesPluginType.Macvlan,
}},
},
AvailabilityZones = {"1", "2", "3"},
Labels = {new KubernetesLabel("kubernetes.label", "true")},
Taints = {new KubernetesLabel("kubernetes.taint", "true:NoSchedule")},
UpgradeSettings = new AgentPoolUpgradeSettings
{
MaxSurge = "1",
},
}
},
    "1.XX.Y",
    new KubernetesClusterNetworkConfiguration(new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/cloudServicesNetworks/cloudServicesNetworkName"), new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"))
    {
        AttachedNetworkConfiguration = new AttachedNetworkConfiguration
        {
            L2Networks = {new L2NetworkAttachmentConfiguration(new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l2Networks/l2NetworkName"))
            {
            PluginType = KubernetesPluginType.Dpdk,
            }},
            L3Networks = {new L3NetworkAttachmentConfiguration(new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"))
            {
            IpamEnabled = L3NetworkConfigurationIpamEnabled.False,
            PluginType = KubernetesPluginType.Sriov,
            }},
            TrunkedNetworks = {new TrunkedNetworkAttachmentConfiguration(new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/trunkedNetworks/trunkedNetworkName"))
            {
            PluginType = KubernetesPluginType.Macvlan,
            }},
        },
        DnsServiceIP = IPAddress.Parse("198.51.101.2"),
        L2ServiceLoadBalancerIPAddressPools = {new IPAddressPool(new string[]{"198.51.102.2-198.51.102.254"}, "pool1")
        {
        AutoAssign = BfdEnabled.True,
        }},
        PodCidrs = { "198.51.100.0/24" },
        ServiceCidrs = { "198.51.101.0/24" },
    })
{
    AadAdminGroupObjectIds = { "ffffffff-ffff-ffff-ffff-ffffffffffff" },
    AdministratorConfiguration = new AdministratorConfiguration
    {
        AdminUsername = "azure",
        SshPublicKeys = { new NetworkCloudSshPublicKey("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm") },
    },
    ManagedResourceGroupConfiguration = new ManagedResourceGroupConfiguration
    {
        Location = new AzureLocation("East US"),
        Name = "my-managed-rg",
    },
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2"
    },
};
ArmOperation<NetworkCloudKubernetesClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, kubernetesClusterName, data);
NetworkCloudKubernetesClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkCloudKubernetesClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
