using Azure;
using Azure.ResourceManager;
using System;
using System.Text;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerService.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ContainerService;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/ManagedClustersCreate_DedicatedHostGroup.json
// this example is just showing the usage of "ManagedClusters_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ContainerServiceManagedClusterResource
ContainerServiceManagedClusterCollection collection = resourceGroupResource.GetContainerServiceManagedClusters();

// invoke the operation
string resourceName = "clustername1";
ContainerServiceManagedClusterData data = new ContainerServiceManagedClusterData(new AzureLocation("location1"))
{
    Sku = new ManagedClusterSku
    {
        Name = new ManagedClusterSkuName("Basic"),
        Tier = ManagedClusterSkuTier.Free,
    },
    KubernetesVersion = "",
    DnsPrefix = "dnsprefix1",
    AgentPoolProfiles = {new ManagedClusterAgentPoolProfile("nodepool1")
    {
    Count = 3,
    VmSize = "Standard_DS2_v2",
    OSType = ContainerServiceOSType.Linux,
    AgentPoolType = AgentPoolType.VirtualMachineScaleSets,
    EnableNodePublicIP = true,
    HostGroupId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/Microsoft.Compute/hostGroups/hostgroup1"),
    }},
    LinuxProfile = new ContainerServiceLinuxProfile("azureuser", new ContainerServiceSshConfiguration(new ContainerServiceSshPublicKey[]
{
new ContainerServiceSshPublicKey("keydata")
})),
    WindowsProfile = new ManagedClusterWindowsProfile("azureuser")
    {
        AdminPassword = "replacePassword1234$",
    },
    ServicePrincipalProfile = new ManagedClusterServicePrincipalProfile("clientid")
    {
        Secret = "secret",
    },
    AddonProfiles = { },
    EnableRbac = true,
    NetworkProfile = new ContainerServiceNetworkProfile
    {
        OutboundType = ContainerServiceOutboundType.LoadBalancer,
        LoadBalancerSku = ContainerServiceLoadBalancerSku.Standard,
        LoadBalancerProfile = new ManagedClusterLoadBalancerProfile
        {
            ManagedOutboundIPs = new ManagedClusterLoadBalancerProfileManagedOutboundIPs
            {
                Count = 2,
            },
        },
    },
    AutoScalerProfile = new ManagedClusterAutoScalerProfile
    {
        ScanIntervalInSeconds = "20s",
        ScaleDownDelayAfterAdd = "15m",
    },
    DiskEncryptionSetId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
    Tags =
    {
    ["archv2"] = "",
    ["tier"] = "production"
    },
};
ArmOperation<ContainerServiceManagedClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data);
ContainerServiceManagedClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerServiceManagedClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
