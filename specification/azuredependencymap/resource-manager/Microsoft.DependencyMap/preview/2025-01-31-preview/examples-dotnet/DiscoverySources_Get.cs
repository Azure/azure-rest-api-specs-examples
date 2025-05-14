using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DependencyMap;

// Generated from example definition: 2025-01-31-preview/DiscoverySources_Get.json
// this example is just showing the usage of "DiscoverySourceResource_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DependencyMapResource created on azure
// for more information of creating DependencyMapResource, please refer to the document of DependencyMapResource
string subscriptionId = "D6E58BDB-45F1-41EC-A884-1FC945058848";
string resourceGroupName = "rgdependencyMap";
string mapName = "mapsTest1";
ResourceIdentifier dependencyMapResourceId = DependencyMapResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, mapName);
DependencyMapResource dependencyMap = client.GetDependencyMapResource(dependencyMapResourceId);

// get the collection of this DependencyMapDiscoverySourceResource
DependencyMapDiscoverySourceCollection collection = dependencyMap.GetDependencyMapDiscoverySources();

// invoke the operation
string sourceName = "sourceTest1";
NullableResponse<DependencyMapDiscoverySourceResource> response = await collection.GetIfExistsAsync(sourceName);
DependencyMapDiscoverySourceResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DependencyMapDiscoverySourceData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
