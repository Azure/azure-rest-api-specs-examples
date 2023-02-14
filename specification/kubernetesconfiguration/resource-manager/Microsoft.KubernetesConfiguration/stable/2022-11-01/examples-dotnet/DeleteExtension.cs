using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.KubernetesConfiguration;
using Azure.ResourceManager.KubernetesConfiguration.Models;

// Generated from example definition: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/DeleteExtension.json
// this example is just showing the usage of "Extensions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KubernetesClusterExtensionResource created on azure
// for more information of creating KubernetesClusterExtensionResource, please refer to the document of KubernetesClusterExtensionResource
string subscriptionId = "subId1";
string resourceGroupName = "rg1";
string clusterRp = "Microsoft.Kubernetes";
string clusterResourceName = "connectedClusters";
string clusterName = "clusterName1";
string extensionName = "ClusterMonitor";
ResourceIdentifier kubernetesClusterExtensionResourceId = KubernetesClusterExtensionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterRp, clusterResourceName, clusterName, extensionName);
KubernetesClusterExtensionResource kubernetesClusterExtension = client.GetKubernetesClusterExtensionResource(kubernetesClusterExtensionResourceId);

// invoke the operation
await kubernetesClusterExtension.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
