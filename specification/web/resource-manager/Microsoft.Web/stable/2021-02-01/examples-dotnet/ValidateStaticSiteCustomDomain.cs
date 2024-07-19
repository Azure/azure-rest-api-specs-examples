using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/ValidateStaticSiteCustomDomain.json
// this example is just showing the usage of "StaticSites_ValidateCustomDomainCanBeAddedToStaticSite" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StaticSiteCustomDomainOverviewResource created on azure
// for more information of creating StaticSiteCustomDomainOverviewResource, please refer to the document of StaticSiteCustomDomainOverviewResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string name = "testStaticSite0";
string domainName = "custom.domain.net";
ResourceIdentifier staticSiteCustomDomainOverviewResourceId = StaticSiteCustomDomainOverviewResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, domainName);
StaticSiteCustomDomainOverviewResource staticSiteCustomDomainOverview = client.GetStaticSiteCustomDomainOverviewResource(staticSiteCustomDomainOverviewResourceId);

// invoke the operation
StaticSiteCustomDomainContent content = new StaticSiteCustomDomainContent();
await staticSiteCustomDomainOverview.ValidateCustomDomainCanBeAddedToStaticSiteAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");
