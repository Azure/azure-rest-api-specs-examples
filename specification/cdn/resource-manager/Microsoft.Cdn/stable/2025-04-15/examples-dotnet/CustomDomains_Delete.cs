using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2025-04-15/examples/CustomDomains_Delete.json
// this example is just showing the usage of "CdnCustomDomains_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CdnCustomDomainResource created on azure
// for more information of creating CdnCustomDomainResource, please refer to the document of CdnCustomDomainResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
string endpointName = "endpoint1";
string customDomainName = "www-someDomain-net";
ResourceIdentifier cdnCustomDomainResourceId = CdnCustomDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, endpointName, customDomainName);
CdnCustomDomainResource cdnCustomDomain = client.GetCdnCustomDomainResource(cdnCustomDomainResourceId);

// invoke the operation
await cdnCustomDomain.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
