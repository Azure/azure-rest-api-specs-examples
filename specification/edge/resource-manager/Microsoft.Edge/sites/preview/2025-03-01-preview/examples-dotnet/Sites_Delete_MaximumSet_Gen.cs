using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SiteManager.Models;
using Azure.ResourceManager.SiteManager;

// Generated from example definition: 2025-03-01-preview/Sites_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "Site_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeSiteResource created on azure
// for more information of creating EdgeSiteResource, please refer to the document of EdgeSiteResource
string subscriptionId = "0154f7fe-df09-4981-bf82-7ad5c1f596eb";
string resourceGroupName = "rgsites";
string siteName = "string";
ResourceIdentifier edgeSiteResourceId = EdgeSiteResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, siteName);
EdgeSiteResource edgeSite = client.GetEdgeSiteResource(edgeSiteResourceId);

// invoke the operation
await edgeSite.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
