using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Cdn;
using Azure.ResourceManager.Cdn.Models;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/CustomDomains_Create.json
// this example is just showing the usage of "CdnCustomDomains_Create" operation, for the dependent resources, they will have to be created separately.

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
CdnCustomDomainCreateOrUpdateContent content = new CdnCustomDomainCreateOrUpdateContent()
{
    HostName = "www.someDomain.net",
};
ArmOperation<CdnCustomDomainResource> lro = await cdnCustomDomain.UpdateAsync(WaitUntil.Completed, content);
CdnCustomDomainResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CdnCustomDomainData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
