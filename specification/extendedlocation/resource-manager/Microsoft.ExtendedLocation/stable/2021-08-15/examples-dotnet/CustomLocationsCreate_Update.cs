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

// Generated from example definition: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsCreate_Update.json
// this example is just showing the usage of "CustomLocations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "11111111-2222-3333-4444-555555555555";
string resourceGroupName = "testresourcegroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this CustomLocationResource
CustomLocationCollection collection = resourceGroupResource.GetCustomLocations();

// invoke the operation
string resourceName = "customLocation01";
CustomLocationData data = new CustomLocationData(new AzureLocation("West US"))
{
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    Authentication = new CustomLocationAuthentication()
    {
        CustomLocationPropertiesAuthenticationType = "KubeConfig",
        Value = "<base64 KubeConfig>",
    },
    ClusterExtensionIds =
    {
    new ResourceIdentifier("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedCluster/someCluster/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension")
    },
    DisplayName = "customLocationLocation01",
    HostResourceId = new ResourceIdentifier("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01"),
    Namespace = "namespace01",
};
ArmOperation<CustomLocationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data);
CustomLocationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CustomLocationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
