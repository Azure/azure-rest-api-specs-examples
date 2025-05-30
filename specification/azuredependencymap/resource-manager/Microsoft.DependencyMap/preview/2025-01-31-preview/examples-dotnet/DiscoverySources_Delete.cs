using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DependencyMap.Models;
using Azure.ResourceManager.DependencyMap;

// Generated from example definition: 2025-01-31-preview/DiscoverySources_Delete.json
// this example is just showing the usage of "DiscoverySourceResource_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DependencyMapDiscoverySourceResource created on azure
// for more information of creating DependencyMapDiscoverySourceResource, please refer to the document of DependencyMapDiscoverySourceResource
string subscriptionId = "D6E58BDB-45F1-41EC-A884-1FC945058848";
string resourceGroupName = "rgdependencyMap";
string mapName = "mapsTest1";
string sourceName = "sourceTest1";
ResourceIdentifier dependencyMapDiscoverySourceResourceId = DependencyMapDiscoverySourceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, mapName, sourceName);
DependencyMapDiscoverySourceResource dependencyMapDiscoverySource = client.GetDependencyMapDiscoverySourceResource(dependencyMapDiscoverySourceResourceId);

// invoke the operation
await dependencyMapDiscoverySource.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
