using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Orbital.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Orbital;

// Generated from example definition: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/SpacecraftCreate.json
// this example is just showing the usage of "Spacecrafts_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "c1be1141-a7c9-4aac-9608-3c2e2f1152c3";
string resourceGroupName = "contoso-Rgp";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this OrbitalSpacecraftResource
OrbitalSpacecraftCollection collection = resourceGroupResource.GetOrbitalSpacecrafts();

// invoke the operation
string spacecraftName = "CONTOSO_SAT";
OrbitalSpacecraftData data = new OrbitalSpacecraftData(new AzureLocation("eastus2"))
{
    NoradId = "36411",
    TitleLine = "CONTOSO_SAT",
    TleLine1 = "1 27424U 02022A   22167.05119303  .00000638  00000+0  15103-3 0  9994",
    TleLine2 = "2 27424  98.2477 108.9546 0000928  92.9194 327.0802 14.57300770 69982",
    Links = { new OrbitalSpacecraftLink("uplink_lhcp1", 2250, 2, OrbitalLinkDirection.Uplink, OrbitalLinkPolarization.Lhcp), new OrbitalSpacecraftLink("downlink_rhcp1", 8160, 15, OrbitalLinkDirection.Downlink, OrbitalLinkPolarization.Rhcp) },
};
ArmOperation<OrbitalSpacecraftResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, spacecraftName, data);
OrbitalSpacecraftResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
OrbitalSpacecraftData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
