using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/PatchStaticSite.json
// this example is just showing the usage of "StaticSites_UpdateStaticSite" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StaticSiteResource created on azure
// for more information of creating StaticSiteResource, please refer to the document of StaticSiteResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string name = "testStaticSite0";
ResourceIdentifier staticSiteResourceId = StaticSiteResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
StaticSiteResource staticSite = client.GetStaticSiteResource(staticSiteResourceId);

// invoke the operation
StaticSitePatch patch = new StaticSitePatch();
StaticSiteResource result = await staticSite.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StaticSiteData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
