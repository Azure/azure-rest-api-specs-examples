using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Marketplace;
using Azure.ResourceManager.Marketplace.Models;

// Generated from example definition: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2023-01-01/examples/QueryApprovedPlans.json
// this example is just showing the usage of "PrivateStore_QueryApprovedPlans" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PrivateStoreResource created on azure
// for more information of creating PrivateStoreResource, please refer to the document of PrivateStoreResource
Guid privateStoreId = Guid.Parse("a0e28e55-90c4-41d8-8e34-bb7ef7775406");
ResourceIdentifier privateStoreResourceId = PrivateStoreResource.CreateResourceIdentifier(privateStoreId);
PrivateStoreResource privateStore = client.GetPrivateStoreResource(privateStoreResourceId);

// invoke the operation
QueryApprovedPlansContent content = new QueryApprovedPlansContent()
{
    OfferId = "marketplacetestthirdparty.md-test-third-party-2",
    PlanIds =
    {
    "testPlanA","testPlanB","testPlanC"
    },
    SubscriptionIds =
    {
    "85e3e079-c718-4e4c-abbe-f72fceba8305","7752d461-4bf1-4185-8b56-8a3f11486ac6"
    },
};
QueryApprovedPlansResult result = await privateStore.QueryApprovedPlansAsync(content: content);

Console.WriteLine($"Succeeded: {result}");
