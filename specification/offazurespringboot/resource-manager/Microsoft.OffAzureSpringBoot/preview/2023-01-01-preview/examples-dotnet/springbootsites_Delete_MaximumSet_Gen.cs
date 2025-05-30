using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SpringAppDiscovery.Models;
using Azure.ResourceManager.SpringAppDiscovery;

// Generated from example definition: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootsites_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "springbootsites_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SpringBootSiteResource created on azure
// for more information of creating SpringBootSiteResource, please refer to the document of SpringBootSiteResource
string subscriptionId = "chshxczdscjpcyvyethat";
string resourceGroupName = "rgspringbootsites";
string springbootsitesName = "xrmzlavpewxtfeitghdrj";
ResourceIdentifier springBootSiteResourceId = SpringBootSiteResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, springbootsitesName);
SpringBootSiteResource springBootSite = client.GetSpringBootSiteResource(springBootSiteResourceId);

// invoke the operation
await springBootSite.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
