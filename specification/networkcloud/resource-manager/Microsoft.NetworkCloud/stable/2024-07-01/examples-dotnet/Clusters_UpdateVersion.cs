using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2024-07-01/examples/Clusters_UpdateVersion.json
// this example is just showing the usage of "Clusters_UpdateVersion" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkCloudClusterResource created on azure
// for more information of creating NetworkCloudClusterResource, please refer to the document of NetworkCloudClusterResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
string clusterName = "clusterName";
ResourceIdentifier networkCloudClusterResourceId = NetworkCloudClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
NetworkCloudClusterResource networkCloudCluster = client.GetNetworkCloudClusterResource(networkCloudClusterResourceId);

// invoke the operation
ClusterUpdateVersionContent content = new ClusterUpdateVersionContent("2.0");
ArmOperation<NetworkCloudOperationStatusResult> lro = await networkCloudCluster.UpdateVersionAsync(WaitUntil.Completed, content);
NetworkCloudOperationStatusResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
