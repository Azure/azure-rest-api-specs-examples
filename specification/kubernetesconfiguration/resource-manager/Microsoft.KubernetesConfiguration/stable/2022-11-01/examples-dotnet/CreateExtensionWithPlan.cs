using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.KubernetesConfiguration.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.KubernetesConfiguration;

// Generated from example definition: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-11-01/examples/CreateExtensionWithPlan.json
// this example is just showing the usage of "Extensions_Create" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subId1";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this KubernetesClusterExtensionResource
string clusterRp = "Microsoft.Kubernetes";
string clusterResourceName = "connectedClusters";
string clusterName = "clusterName1";
KubernetesClusterExtensionCollection collection = resourceGroupResource.GetKubernetesClusterExtensions(clusterRp, clusterResourceName, clusterName);

// invoke the operation
string extensionName = "azureVote";
KubernetesClusterExtensionData data = new KubernetesClusterExtensionData()
{
    Plan = new ArmPlan("azure-vote-standard", "Microsoft", "azure-vote-standard-offer-id"),
    ExtensionType = "azure-vote",
    AutoUpgradeMinorVersion = true,
    ReleaseTrain = "Preview",
};
ArmOperation<KubernetesClusterExtensionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, extensionName, data);
KubernetesClusterExtensionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KubernetesClusterExtensionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
