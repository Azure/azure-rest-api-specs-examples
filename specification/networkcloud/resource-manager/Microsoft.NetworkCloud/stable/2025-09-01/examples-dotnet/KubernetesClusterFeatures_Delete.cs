using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/KubernetesClusterFeatures_Delete.json
// this example is just showing the usage of "KubernetesClusterFeatures_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkCloudKubernetesClusterFeatureResource created on azure
// for more information of creating NetworkCloudKubernetesClusterFeatureResource, please refer to the document of NetworkCloudKubernetesClusterFeatureResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
string kubernetesClusterName = "kubernetesClusterName";
string featureName = "featureName";
ResourceIdentifier networkCloudKubernetesClusterFeatureResourceId = NetworkCloudKubernetesClusterFeatureResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, kubernetesClusterName, featureName);
NetworkCloudKubernetesClusterFeatureResource networkCloudKubernetesClusterFeature = client.GetNetworkCloudKubernetesClusterFeatureResource(networkCloudKubernetesClusterFeatureResourceId);

// invoke the operation
ArmOperation<NetworkCloudOperationStatusResult> lro = await networkCloudKubernetesClusterFeature.DeleteAsync(WaitUntil.Completed);
NetworkCloudOperationStatusResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
