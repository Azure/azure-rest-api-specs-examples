using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-06-01-preview/examples/Clusters_Patch_CommandOutput.json
// this example is just showing the usage of "Clusters_Update" operation, for the dependent resources, they will have to be created separately.

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
NetworkCloudClusterPatch patch = new NetworkCloudClusterPatch
{
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity1")] = new UserAssignedIdentity(),
        [new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity2")] = new UserAssignedIdentity()
        },
    },
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2"
    },
    CommandOutputSettings = new CommandOutputSettings
    {
        AssociatedIdentity = new IdentitySelector
        {
            IdentityType = ManagedServiceIdentitySelectorType.UserAssignedIdentity,
            UserAssignedIdentityResourceId = new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/userIdentity2"),
        },
        ContainerUri = new Uri("https://myaccount.blob.core.windows.net/mycontainer?restype=container"),
    },
};
ArmOperation<NetworkCloudClusterResource> lro = await networkCloudCluster.UpdateAsync(WaitUntil.Completed, patch);
NetworkCloudClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkCloudClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
