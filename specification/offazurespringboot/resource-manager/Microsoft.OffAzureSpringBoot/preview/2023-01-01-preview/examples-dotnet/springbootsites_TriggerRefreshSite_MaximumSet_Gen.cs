using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SpringAppDiscovery;
using Azure.ResourceManager.SpringAppDiscovery.Models;

// Generated from example definition: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootsites_TriggerRefreshSite_MaximumSet_Gen.json
// this example is just showing the usage of "springbootsites_TriggerRefreshSite" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SpringBootSiteResource created on azure
// for more information of creating SpringBootSiteResource, please refer to the document of SpringBootSiteResource
string subscriptionId = "z";
string resourceGroupName = "rgspringbootsites";
string springbootsitesName = "czarpuxwoafaqsuptutcwyu";
ResourceIdentifier springBootSiteResourceId = SpringBootSiteResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, springbootsitesName);
SpringBootSiteResource springBootSite = client.GetSpringBootSiteResource(springBootSiteResourceId);

// invoke the operation
await springBootSite.TriggerRefreshSiteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
