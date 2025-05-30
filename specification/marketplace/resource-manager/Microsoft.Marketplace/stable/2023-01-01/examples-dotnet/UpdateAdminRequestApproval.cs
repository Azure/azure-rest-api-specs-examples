using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Marketplace.Models;
using Azure.ResourceManager.Marketplace;

// Generated from example definition: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2023-01-01/examples/UpdateAdminRequestApproval.json
// this example is just showing the usage of "PrivateStore_UpdateAdminRequestApproval" operation, for the dependent resources, they will have to be created separately.

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
MarketplaceAdminApprovalRequestData data = new MarketplaceAdminApprovalRequestData
{
    PublisherId = "marketplacetestthirdparty",
    AdminAction = MarketplaceAdminAction.Approved,
    ApprovedPlans = { "testPlan" },
    Comment = "I'm ok with that",
    CollectionIds = { Guid.Parse("f8ee227e-85d7-477d-abbf-854d6decaf70"), Guid.Parse("39246ad6-c521-4fed-8de7-77dede2e873f") },
};
ArmOperation<MarketplaceAdminApprovalRequestResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, adminRequestApprovalId, data);
MarketplaceAdminApprovalRequestResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MarketplaceAdminApprovalRequestData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
