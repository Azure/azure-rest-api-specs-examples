using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Orbital.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Orbital;

// Generated from example definition: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/AvailableContactsList.json
// this example is just showing the usage of "Spacecrafts_ListAvailableContacts" operation, for the dependent resources, they will have to be created separately.

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
OrbitalAvailableContactsContent content = new OrbitalAvailableContactsContent(new WritableSubResource
{
    Id = new ResourceIdentifier("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/contactProfiles/CONTOSO-CP"),
}, "EASTUS2_0", DateTimeOffset.Parse("2022-03-01T11:30:00Z"), DateTimeOffset.Parse("2022-03-02T11:30:00Z"));
ArmOperation<OrbitalAvailableContactsResult> lro = await orbitalSpacecraft.GetAllAvailableContactsAsync(WaitUntil.Completed, content);
OrbitalAvailableContactsResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
