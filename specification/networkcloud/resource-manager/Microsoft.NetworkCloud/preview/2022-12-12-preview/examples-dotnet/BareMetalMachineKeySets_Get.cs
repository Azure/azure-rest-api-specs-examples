using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BareMetalMachineKeySets_Get.json
// this example is just showing the usage of "BareMetalMachineKeySets_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ClusterResource created on azure
// for more information of creating ClusterResource, please refer to the document of ClusterResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string clusterName = "clusterName";
ResourceIdentifier clusterResourceId = ClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
ClusterResource cluster = client.GetClusterResource(clusterResourceId);

// get the collection of this BareMetalMachineKeySetResource
BareMetalMachineKeySetCollection collection = cluster.GetBareMetalMachineKeySets();

// invoke the operation
string bareMetalMachineKeySetName = "bareMetalMachineKeySetName";
bool result = await collection.ExistsAsync(bareMetalMachineKeySetName);

Console.WriteLine($"Succeeded: {result}");
