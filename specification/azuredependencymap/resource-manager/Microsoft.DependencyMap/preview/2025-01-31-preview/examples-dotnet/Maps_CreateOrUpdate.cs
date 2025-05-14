using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DependencyMap;

// Generated from example definition: 2025-01-31-preview/Maps_CreateOrUpdate.json
// this example is just showing the usage of "MapsResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "D6E58BDB-45F1-41EC-A884-1FC945058848";
string resourceGroupName = "rgdependencyMap";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DependencyMapResource
DependencyMapCollection collection = resourceGroupResource.GetDependencyMaps();

// invoke the operation
string mapName = "mapsTest1";
DependencyMapData data = new DependencyMapData(new AzureLocation("wjtzelgfcmswfeflfltwxqveo"))
{
    Tags = { },
};
ArmOperation<DependencyMapResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, mapName, data);
DependencyMapResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DependencyMapData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
