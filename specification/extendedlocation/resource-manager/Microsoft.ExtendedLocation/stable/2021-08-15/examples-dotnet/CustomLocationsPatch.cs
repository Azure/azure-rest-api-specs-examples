using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ExtendedLocations;
using Azure.ResourceManager.ExtendedLocations.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsPatch.json
// this example is just showing the usage of "CustomLocations_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CustomLocationResource created on azure
// for more information of creating CustomLocationResource, please refer to the document of CustomLocationResource
string subscriptionId = "11111111-2222-3333-4444-555555555555";
string resourceGroupName = "testresourcegroup";
string resourceName = "customLocation01";
ResourceIdentifier customLocationResourceId = CustomLocationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
CustomLocationResource customLocation = client.GetCustomLocationResource(customLocationResourceId);

// invoke the operation
CustomLocationPatch patch = new CustomLocationPatch()
{
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    Tags =
    {
    ["archv3"] = "",
    ["tier"] = "testing",
    },
    ClusterExtensionIds =
    {
    new ResourceIdentifier("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension"),new ResourceIdentifier("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/barExtension")
    },
};
CustomLocationResource result = await customLocation.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CustomLocationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
