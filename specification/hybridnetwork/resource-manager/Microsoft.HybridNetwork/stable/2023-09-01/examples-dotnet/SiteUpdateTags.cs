using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/SiteUpdateTags.json
// this example is just showing the usage of "Sites_UpdateTags" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteResource created on azure
// for more information of creating SiteResource, please refer to the document of SiteResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string siteName = "testSite";
ResourceIdentifier siteResourceId = SiteResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, siteName);
SiteResource site = client.GetSiteResource(siteResourceId);

// invoke the operation
TagsObject tagsObject = new TagsObject()
{
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2",
    },
};
SiteResource result = await site.UpdateAsync(tagsObject);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
