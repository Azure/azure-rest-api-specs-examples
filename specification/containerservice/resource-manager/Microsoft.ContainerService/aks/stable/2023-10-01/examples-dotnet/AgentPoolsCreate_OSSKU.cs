using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerService.Models;
using Azure.ResourceManager.ContainerService;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-10-01/examples/AgentPoolsCreate_OSSKU.json
// this example is just showing the usage of "AgentPools_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerServiceManagedClusterResource created on azure
// for more information of creating ContainerServiceManagedClusterResource, please refer to the document of ContainerServiceManagedClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string resourceName = "clustername1";
ResourceIdentifier containerServiceManagedClusterResourceId = ContainerServiceManagedClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
ContainerServiceManagedClusterResource containerServiceManagedCluster = client.GetContainerServiceManagedClusterResource(containerServiceManagedClusterResourceId);

// get the collection of this ContainerServiceAgentPoolResource
ContainerServiceAgentPoolCollection collection = containerServiceManagedCluster.GetContainerServiceAgentPools();

// invoke the operation
string agentPoolName = "agentpool1";
ContainerServiceAgentPoolData data = new ContainerServiceAgentPoolData
{
    Count = 3,
    VmSize = "Standard_DS2_v2",
    OSType = ContainerServiceOSType.Linux,
    OSSku = ContainerServiceOSSku.AzureLinux,
    OrchestratorVersion = "",
    KubeletConfig = new KubeletConfig
    {
        CpuManagerPolicy = "static",
        IsCpuCfsQuotaEnabled = true,
        CpuCfsQuotaPeriod = "200ms",
        ImageGcHighThreshold = 90,
        ImageGcLowThreshold = 70,
        TopologyManagerPolicy = "best-effort",
        AllowedUnsafeSysctls = { "kernel.msg*", "net.core.somaxconn" },
        FailStartWithSwapOn = false,
    },
    LinuxOSConfig = new LinuxOSConfig
    {
        Sysctls = new SysctlConfig
        {
            NetCoreWmemDefault = 12345,
            NetIPv4TcpTwReuse = true,
            NetIPv4IPLocalPortRange = "20000 60000",
            KernelThreadsMax = 99999,
        },
        TransparentHugePageEnabled = "always",
        TransparentHugePageDefrag = "madvise",
        SwapFileSizeInMB = 1500,
    },
};
ArmOperation<ContainerServiceAgentPoolResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, agentPoolName, data);
ContainerServiceAgentPoolResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerServiceAgentPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
