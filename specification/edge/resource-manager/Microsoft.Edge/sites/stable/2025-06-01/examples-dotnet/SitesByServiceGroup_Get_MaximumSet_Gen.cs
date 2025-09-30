using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SiteManager.Models;
using Azure.ResourceManager.SiteManager;

// Generated from example definition: 2025-06-01/SitesByServiceGroup_Get_MaximumSet_Gen.json
// this example is just showing the usage of "Site_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this ServiceGroupEdgeSiteResource
string servicegroupName = "string";
string scope = $"/providers/Microsoft.Management/serviceGroups/{servicegroupName}";
ServiceGroupEdgeSiteCollection collection = client.GetServiceGroupEdgeSites(new ResourceIdentifier(scope));

// invoke the operation
string siteName = "string";
NullableResponse<ServiceGroupEdgeSiteResource> response = await collection.GetIfExistsAsync(siteName);
ServiceGroupEdgeSiteResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EdgeSiteData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
