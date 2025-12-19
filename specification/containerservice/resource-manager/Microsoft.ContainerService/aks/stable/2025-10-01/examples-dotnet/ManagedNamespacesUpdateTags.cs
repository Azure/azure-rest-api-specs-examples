using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerService.Models;
using Azure.ResourceManager.ContainerService;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/ManagedNamespacesUpdateTags.json
// this example is just showing the usage of "ManagedNamespaces_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedClusterNamespaceResource created on azure
// for more information of creating ManagedClusterNamespaceResource, please refer to the document of ManagedClusterNamespaceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string resourceName = "clustername1";
string managedNamespaceName = "namespace1";
ResourceIdentifier managedClusterNamespaceResourceId = ManagedClusterNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, managedNamespaceName);
ManagedClusterNamespaceResource managedClusterNamespace = client.GetManagedClusterNamespaceResource(managedClusterNamespaceResourceId);

// invoke the operation
ContainerServiceTagsObject containerServiceTagsObject = new ContainerServiceTagsObject
{
    Tags =
    {
    ["tagKey1"] = "tagValue1",
    ["tagKey2"] = "tagValue2"
    },
};
ManagedClusterNamespaceResource result = await managedClusterNamespace.UpdateAsync(containerServiceTagsObject);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedClusterNamespaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
