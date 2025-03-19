using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SpringAppDiscovery;

// Generated from example definition: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/Summaries_Get_MinimumSet_Gen.json
// this example is just showing the usage of "Summaries_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SpringBootSiteResource created on azure
// for more information of creating SpringBootSiteResource, please refer to the document of SpringBootSiteResource
string subscriptionId = "libzegdqkcxmhqhhhcxm";
string resourceGroupName = "rgspringbootdiscovery";
string siteName = "xxkzlvbihwxunadjcpjpjmghmhxrqyvghtpfps";
ResourceIdentifier springBootSiteResourceId = SpringBootSiteResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, siteName);
SpringBootSiteResource springBootSite = client.GetSpringBootSiteResource(springBootSiteResourceId);

// get the collection of this SpringBootSiteSummaryResource
SpringBootSiteSummaryCollection collection = springBootSite.GetSpringBootSiteSummaries();

// invoke the operation
string summaryName = "vjB";
NullableResponse<SpringBootSiteSummaryResource> response = await collection.GetIfExistsAsync(summaryName);
SpringBootSiteSummaryResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SpringBootSiteSummaryData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
