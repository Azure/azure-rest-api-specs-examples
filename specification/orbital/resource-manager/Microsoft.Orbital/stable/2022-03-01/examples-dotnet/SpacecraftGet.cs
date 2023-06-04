using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Orbital;
using Azure.ResourceManager.Orbital.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/SpacecraftGet.json
// this example is just showing the usage of "Spacecrafts_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OrbitalSpacecraftResource created on azure
// for more information of creating OrbitalSpacecraftResource, please refer to the document of OrbitalSpacecraftResource
string subscriptionId = "c1be1141-a7c9-4aac-9608-3c2e2f1152c3";
string resourceGroupName = "contoso-Rgp";
string spacecraftName = "CONTOSO_SAT";
ResourceIdentifier orbitalSpacecraftResourceId = OrbitalSpacecraftResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, spacecraftName);
OrbitalSpacecraftResource orbitalSpacecraft = client.GetOrbitalSpacecraftResource(orbitalSpacecraftResourceId);

// invoke the operation
OrbitalSpacecraftResource result = await orbitalSpacecraft.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
OrbitalSpacecraftData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
