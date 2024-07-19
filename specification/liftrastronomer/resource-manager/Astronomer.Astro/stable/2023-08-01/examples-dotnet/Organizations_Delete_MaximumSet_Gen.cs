using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Astro.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Astro;

// Generated from example definition: specification/liftrastronomer/resource-manager/Astronomer.Astro/stable/2023-08-01/examples/Organizations_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "Organizations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AstroOrganizationResource created on azure
// for more information of creating AstroOrganizationResource, please refer to the document of AstroOrganizationResource
string subscriptionId = "43454B17-172A-40FE-80FA-549EA23D12B3";
string resourceGroupName = "rgastronomer";
string organizationName = "q:";
ResourceIdentifier astroOrganizationResourceId = AstroOrganizationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, organizationName);
AstroOrganizationResource astroOrganization = client.GetAstroOrganizationResource(astroOrganizationResourceId);

// invoke the operation
await astroOrganization.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
