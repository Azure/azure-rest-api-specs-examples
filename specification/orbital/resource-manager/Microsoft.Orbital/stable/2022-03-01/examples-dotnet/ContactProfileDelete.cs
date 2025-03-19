using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Orbital.Models;
using Azure.ResourceManager.Orbital;

// Generated from example definition: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactProfileDelete.json
// this example is just showing the usage of "ContactProfiles_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OrbitalContactProfileResource created on azure
// for more information of creating OrbitalContactProfileResource, please refer to the document of OrbitalContactProfileResource
string subscriptionId = "c1be1141-a7c9-4aac-9608-3c2e2f1152c3";
string resourceGroupName = "contoso-Rgp";
string contactProfileName = "CONTOSO-CP";
ResourceIdentifier orbitalContactProfileResourceId = OrbitalContactProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, contactProfileName);
OrbitalContactProfileResource orbitalContactProfile = client.GetOrbitalContactProfileResource(orbitalContactProfileResourceId);

// invoke the operation
await orbitalContactProfile.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
