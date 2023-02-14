using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.KubernetesConfiguration;
using Azure.ResourceManager.KubernetesConfiguration.Models;

// Generated from example definition: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/GetExtensionWithPlan.json
// this example is just showing the usage of "Extensions_Get" operation, for the dependent resources, they will have to be created separately.

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
string extensionName = "azureVote";
ResourceIdentifier kubernetesClusterExtensionResourceId = KubernetesClusterExtensionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterRp, clusterResourceName, clusterName, extensionName);
KubernetesClusterExtensionResource kubernetesClusterExtension = client.GetKubernetesClusterExtensionResource(kubernetesClusterExtensionResourceId);

// invoke the operation
KubernetesClusterExtensionResource result = await kubernetesClusterExtension.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KubernetesClusterExtensionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
