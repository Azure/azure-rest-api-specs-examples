using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SiteManager.Models;
using Azure.ResourceManager.SiteManager;

// Generated from example definition: 2025-06-01/Sites_Get_MaximumSet_Gen.json
// this example is just showing the usage of "Site_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupEdgeSiteResource created on azure
// for more information of creating ResourceGroupEdgeSiteResource, please refer to the document of ResourceGroupEdgeSiteResource
string subscriptionId = "0154f7fe-df09-4981-bf82-7ad5c1f596eb";
string resourceGroupName = "rgsites";
string siteName = "string";
ResourceIdentifier resourceGroupEdgeSiteResourceId = ResourceGroupEdgeSiteResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, siteName);
ResourceGroupEdgeSiteResource resourceGroupEdgeSite = client.GetResourceGroupEdgeSiteResource(resourceGroupEdgeSiteResourceId);

// invoke the operation
ResourceGroupEdgeSiteResource result = await resourceGroupEdgeSite.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EdgeSiteData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
