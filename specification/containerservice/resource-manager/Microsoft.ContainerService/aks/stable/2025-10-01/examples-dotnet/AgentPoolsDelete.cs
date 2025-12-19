using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerService.Models;
using Azure.ResourceManager.ContainerService;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/AgentPoolsDelete.json
// this example is just showing the usage of "AgentPools_Delete" operation, for the dependent resources, they will have to be created separately.

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
await containerServiceAgentPool.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
