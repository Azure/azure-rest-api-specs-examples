using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-06-01-preview/examples/KubernetesClusters_RestartNode.json
// this example is just showing the usage of "KubernetesClusters_RestartNode" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkCloudKubernetesClusterResource created on azure
// for more information of creating NetworkCloudKubernetesClusterResource, please refer to the document of NetworkCloudKubernetesClusterResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
string kubernetesClusterName = "kubernetesClusterName";
ResourceIdentifier networkCloudKubernetesClusterResourceId = NetworkCloudKubernetesClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, kubernetesClusterName);
NetworkCloudKubernetesClusterResource networkCloudKubernetesCluster = client.GetNetworkCloudKubernetesClusterResource(networkCloudKubernetesClusterResourceId);

// invoke the operation
KubernetesClusterRestartNodeContent content = new KubernetesClusterRestartNodeContent("nodeName");
ArmOperation<NetworkCloudOperationStatusResult> lro = await networkCloudKubernetesCluster.RestartNodeAsync(WaitUntil.Completed, content);
NetworkCloudOperationStatusResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
