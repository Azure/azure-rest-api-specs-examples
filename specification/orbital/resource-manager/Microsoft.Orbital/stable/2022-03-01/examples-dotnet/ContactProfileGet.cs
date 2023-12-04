using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Orbital;
using Azure.ResourceManager.Orbital.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-03-01/examples/ContactProfileGet.json
// this example is just showing the usage of "ContactProfiles_Get" operation, for the dependent resources, they will have to be created separately.

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
OrbitalContactProfileResource result = await orbitalContactProfile.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
OrbitalContactProfileData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
