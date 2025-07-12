using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.PlanetaryComputer.Models;
using Azure.ResourceManager.PlanetaryComputer;

// Generated from example definition: 2025-02-11-preview/GeoCatalogs_Update.json
// this example is just showing the usage of "GeoCatalog_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PlanetaryComputerGeoCatalogResource created on azure
// for more information of creating PlanetaryComputerGeoCatalogResource, please refer to the document of PlanetaryComputerGeoCatalogResource
string subscriptionId = "cd9b6cdf-dcf0-4dca-ab19-82be07b74704";
string resourceGroupName = "MyResourceGroup";
string catalogName = "MyCatalog";
ResourceIdentifier planetaryComputerGeoCatalogResourceId = PlanetaryComputerGeoCatalogResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, catalogName);
PlanetaryComputerGeoCatalogResource planetaryComputerGeoCatalog = client.GetPlanetaryComputerGeoCatalogResource(planetaryComputerGeoCatalogResourceId);

// invoke the operation
PlanetaryComputerGeoCatalogPatch patch = new PlanetaryComputerGeoCatalogPatch
{
    Tags =
    {
    ["MyTag"] = "MyValue"
    },
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/cd9b6cdf-dcf0-4dca-ab19-82be07b74704/resourceGroups/MyResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/MyManagedIdentity")] = new UserAssignedIdentity()
        },
    },
};
ArmOperation<PlanetaryComputerGeoCatalogResource> lro = await planetaryComputerGeoCatalog.UpdateAsync(WaitUntil.Completed, patch);
PlanetaryComputerGeoCatalogResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PlanetaryComputerGeoCatalogData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
