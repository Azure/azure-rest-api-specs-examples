using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Marketplace;
using Azure.ResourceManager.Marketplace.Models;

// Generated from example definition: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2023-01-01/examples/QueryRequestApproval.json
// this example is just showing the usage of "PrivateStore_QueryRequestApproval" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MarketplaceApprovalRequestResource created on azure
// for more information of creating MarketplaceApprovalRequestResource, please refer to the document of MarketplaceApprovalRequestResource
Guid privateStoreId = Guid.Parse("a0e28e55-90c4-41d8-8e34-bb7ef7775406");
string requestApprovalId = "marketplacetestthirdparty.md-test-third-party-2";
ResourceIdentifier marketplaceApprovalRequestResourceId = MarketplaceApprovalRequestResource.CreateResourceIdentifier(privateStoreId, requestApprovalId);
MarketplaceApprovalRequestResource marketplaceApprovalRequest = client.GetMarketplaceApprovalRequestResource(marketplaceApprovalRequestResourceId);

// invoke the operation
QueryApprovalRequestContent content = new QueryApprovalRequestContent()
{
    PublisherId = "marketplacetestthirdparty",
    PlanIds =
    {
    "testPlanA","testPlanB","*"
    },
};
QueryApprovalRequestResult result = await marketplaceApprovalRequest.QueryApprovalRequestAsync(content: content);

Console.WriteLine($"Succeeded: {result}");
