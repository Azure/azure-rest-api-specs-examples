using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridContainerService;
using Azure.ResourceManager.HybridContainerService.Models;

// Generated from example definition: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/ProvisionedClusterInstanceListUserKubeconfig.json
// this example is just showing the usage of "provisionedClusterInstances_ListUserKubeconfig" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProvisionedClusterResource created on azure
// for more information of creating ProvisionedClusterResource, please refer to the document of ProvisionedClusterResource
string connectedClusterResourceUri = "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster";
ResourceIdentifier provisionedClusterResourceId = ProvisionedClusterResource.CreateResourceIdentifier(connectedClusterResourceUri);
ProvisionedClusterResource provisionedCluster = client.GetProvisionedClusterResource(provisionedClusterResourceId);

// invoke the operation
ArmOperation<HybridContainerServiceCredentialListResult> lro = await provisionedCluster.GetUserKubeconfigAsync(WaitUntil.Completed);
HybridContainerServiceCredentialListResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
