using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerService.Models;
using Azure.ResourceManager.ContainerService;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/AgentPoolsCreate_CRG.json
// this example is just showing the usage of "AgentPools_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerServiceAgentPoolResource created on azure
// for more information of creating ContainerServiceAgentPoolResource, please refer to the document of ContainerServiceAgentPoolResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string resourceName = "clustername1";
string agentPoolName = "agentpool1";
ResourceIdentifier containerServiceAgentPoolResourceId = ContainerServiceAgentPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, agentPoolName);
ContainerServiceAgentPoolResource containerServiceAgentPool = client.GetContainerServiceAgentPoolResource(containerServiceAgentPoolResourceId);

// invoke the operation
ContainerServiceAgentPoolData data = new ContainerServiceAgentPoolData
{
    Count = 3,
    VmSize = "Standard_DS2_v2",
    OSType = ContainerServiceOSType.Linux,
    OrchestratorVersion = "",
    CapacityReservationGroupId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.Compute/CapacityReservationGroups/crg1"),
};
ArmOperation<ContainerServiceAgentPoolResource> lro = await containerServiceAgentPool.UpdateAsync(WaitUntil.Completed, data);
ContainerServiceAgentPoolResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerServiceAgentPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
