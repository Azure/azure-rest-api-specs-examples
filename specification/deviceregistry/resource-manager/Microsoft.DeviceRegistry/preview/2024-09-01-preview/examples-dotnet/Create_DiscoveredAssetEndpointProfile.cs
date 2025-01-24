using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DeviceRegistry.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DeviceRegistry;

// Generated from example definition: 2024-09-01-preview/Create_DiscoveredAssetEndpointProfile.json
// this example is just showing the usage of "DiscoveredAssetEndpointProfile_CreateOrReplace" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DeviceRegistryDiscoveredAssetEndpointProfileResource
DeviceRegistryDiscoveredAssetEndpointProfileCollection collection = resourceGroupResource.GetDeviceRegistryDiscoveredAssetEndpointProfiles();

// invoke the operation
string discoveredAssetEndpointProfileName = "my-discoveredassetendpointprofile";
DeviceRegistryDiscoveredAssetEndpointProfileData data = new DeviceRegistryDiscoveredAssetEndpointProfileData(new AzureLocation("West Europe"), new DeviceRegistryExtendedLocation("CustomLocation", "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"))
{
    Properties = new DiscoveredAssetEndpointProfileProperties(new Uri("https://www.example.com/myTargetAddress"), "myEndpointProfileType", "11111111-1111-1111-1111-111111111111", 73766L)
    {
        AdditionalConfiguration = "{\"foo\": \"bar\"}",
        SupportedAuthenticationMethods = { AuthenticationMethod.Anonymous, AuthenticationMethod.Certificate, AuthenticationMethod.UsernamePassword },
    },
    Tags =
    {
    ["site"] = "building-1"
    },
};
ArmOperation<DeviceRegistryDiscoveredAssetEndpointProfileResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, discoveredAssetEndpointProfileName, data);
DeviceRegistryDiscoveredAssetEndpointProfileResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DeviceRegistryDiscoveredAssetEndpointProfileData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
