using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2025-04-15/examples/Origins_Create.json
// this example is just showing the usage of "CdnOrigins_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CdnEndpointResource created on azure
// for more information of creating CdnEndpointResource, please refer to the document of CdnEndpointResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
string endpointName = "endpoint1";
ResourceIdentifier cdnEndpointResourceId = CdnEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, endpointName);
CdnEndpointResource cdnEndpoint = client.GetCdnEndpointResource(cdnEndpointResourceId);

// get the collection of this CdnOriginResource
CdnOriginCollection collection = cdnEndpoint.GetCdnOrigins();

// invoke the operation
string originName = "www-someDomain-net";
CdnOriginData data = new CdnOriginData
{
    HostName = "www.someDomain.net",
    HttpPort = 80,
    HttpsPort = 443,
    OriginHostHeader = "www.someDomain.net",
    Priority = 1,
    Weight = 50,
    Enabled = true,
    PrivateLinkResourceId = new ResourceIdentifier("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.Network/privateLinkServices/pls1"),
    PrivateLinkLocation = "eastus",
    PrivateLinkApprovalMessage = "Please approve the connection request for this Private Link",
};
ArmOperation<CdnOriginResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, originName, data);
CdnOriginResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CdnOriginData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
