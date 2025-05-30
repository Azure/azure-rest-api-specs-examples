using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridContainerService.Models;
using Azure.ResourceManager.HybridContainerService;

// Generated from example definition: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/DeleteAgentPool.json
// this example is just showing the usage of "agentPool_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridContainerServiceAgentPoolResource created on azure
// for more information of creating HybridContainerServiceAgentPoolResource, please refer to the document of HybridContainerServiceAgentPoolResource
string connectedClusterResourceUri = "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster";
string agentPoolName = "testnodepool";
ResourceIdentifier hybridContainerServiceAgentPoolResourceId = HybridContainerServiceAgentPoolResource.CreateResourceIdentifier(connectedClusterResourceUri, agentPoolName);
HybridContainerServiceAgentPoolResource hybridContainerServiceAgentPool = client.GetHybridContainerServiceAgentPoolResource(hybridContainerServiceAgentPoolResourceId);

// invoke the operation
await hybridContainerServiceAgentPool.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
