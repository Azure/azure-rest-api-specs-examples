using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Origins_Update.json
// this example is just showing the usage of "CdnOrigins_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CdnOriginResource created on azure
// for more information of creating CdnOriginResource, please refer to the document of CdnOriginResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
string endpointName = "endpoint1";
string originName = "www-someDomain-net";
ResourceIdentifier cdnOriginResourceId = CdnOriginResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, endpointName, originName);
CdnOriginResource cdnOrigin = client.GetCdnOriginResource(cdnOriginResourceId);

// invoke the operation
CdnOriginPatch patch = new CdnOriginPatch
{
    HttpPort = 42,
    HttpsPort = 43,
    OriginHostHeader = "www.someDomain2.net",
    Priority = 1,
    Weight = 50,
    Enabled = true,
    PrivateLinkAlias = "APPSERVER.d84e61f0-0870-4d24-9746-7438fa0019d1.westus2.azure.privatelinkservice",
};
ArmOperation<CdnOriginResource> lro = await cdnOrigin.UpdateAsync(WaitUntil.Completed, patch);
CdnOriginResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CdnOriginData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
