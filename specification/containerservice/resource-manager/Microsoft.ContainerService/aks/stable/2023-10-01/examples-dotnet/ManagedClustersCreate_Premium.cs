using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerService.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ContainerService;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-10-01/examples/ManagedClustersCreate_Premium.json
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
    Sku = new ManagedClusterSku()
    {
        Name = ManagedClusterSkuName.Base,
        Tier = ManagedClusterSkuTier.Premium,
    },
    KubernetesVersion = "",
    DnsPrefix = "dnsprefix1",
    AgentPoolProfiles =
    {
    new ManagedClusterAgentPoolProfile("nodepool1")
    {
    Count = 3,
    VmSize = "Standard_DS2_v2",
    OSType = ContainerServiceOSType.Linux,
    AgentPoolType = AgentPoolType.VirtualMachineScaleSets,
    Mode = AgentPoolMode.System,
    EnableNodePublicIP = true,
    EnableEncryptionAtHost = true,
    }
    },
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
    AddonProfiles =
    {
    },
    EnableRbac = true,
    SupportPlan = KubernetesSupportPlan.AKSLongTermSupport,
    EnablePodSecurityPolicy = true,
    NetworkProfile = new ContainerServiceNetworkProfile()
    {
        OutboundType = ContainerServiceOutboundType.LoadBalancer,
        LoadBalancerSku = ContainerServiceLoadBalancerSku.Standard,
        LoadBalancerProfile = new ManagedClusterLoadBalancerProfile()
        {
            ManagedOutboundIPs = new ManagedClusterLoadBalancerProfileManagedOutboundIPs()
            {
                Count = 2,
            },
        },
    },
    AutoScalerProfile = new ManagedClusterAutoScalerProfile()
    {
        ScanIntervalInSeconds = "20s",
        ScaleDownDelayAfterAdd = "15m",
    },
    ApiServerAccessProfile = new ManagedClusterApiServerAccessProfile()
    {
        DisableRunCommand = true,
    },
    Tags =
    {
    ["archv2"] = "",
    ["tier"] = "production",
    },
};
ArmOperation<ContainerServiceManagedClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data);
ContainerServiceManagedClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerServiceManagedClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
