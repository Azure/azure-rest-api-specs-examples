using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Marketplace.Models;
using Azure.ResourceManager.Marketplace;

// Generated from example definition: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2023-01-01/examples/GetAdminRequestApproval.json
// this example is just showing the usage of "PrivateStore_GetAdminRequestApproval" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PrivateStoreResource created on azure
// for more information of creating PrivateStoreResource, please refer to the document of PrivateStoreResource
Guid privateStoreId = Guid.Parse("a0e28e55-90c4-41d8-8e34-bb7ef7775406");
ResourceIdentifier privateStoreResourceId = PrivateStoreResource.CreateResourceIdentifier(privateStoreId);
PrivateStoreResource privateStore = client.GetPrivateStoreResource(privateStoreResourceId);

// get the collection of this MarketplaceAdminApprovalRequestResource
MarketplaceAdminApprovalRequestCollection collection = privateStore.GetMarketplaceAdminApprovalRequests();

// invoke the operation
string adminRequestApprovalId = "marketplacetestthirdparty.md-test-third-party-2";
string publisherId = "marketplacetestthirdparty";
NullableResponse<MarketplaceAdminApprovalRequestResource> response = await collection.GetIfExistsAsync(adminRequestApprovalId, publisherId);
MarketplaceAdminApprovalRequestResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MarketplaceAdminApprovalRequestData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
