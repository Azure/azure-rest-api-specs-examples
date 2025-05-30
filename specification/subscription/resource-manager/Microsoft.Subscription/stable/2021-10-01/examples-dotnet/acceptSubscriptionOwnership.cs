using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Subscription.Models;
using Azure.ResourceManager.Subscription;

// Generated from example definition: specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/acceptSubscriptionOwnership.json
// this example is just showing the usage of "Subscription_AcceptOwnership" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
string subscriptionId = "291bba3f-e0a5-47bc-a099-3bdcb2a50a05";
AcceptOwnershipContent content = new AcceptOwnershipContent
{
    Properties = new AcceptOwnershipRequestProperties("Test Subscription")
    {
        ManagementGroupId = null,
        Tags =
        {
        ["tag1"] = "Messi",
        ["tag2"] = "Ronaldo",
        ["tag3"] = "Lebron"
        },
    },
};
await tenantResource.AcceptSubscriptionOwnershipAsync(WaitUntil.Completed, subscriptionId, content);

Console.WriteLine("Succeeded");
