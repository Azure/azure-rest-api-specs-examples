using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Subscription.Models;
using Azure.ResourceManager.Subscription;

// Generated from example definition: specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/changeTenantPolicy.json
// this example is just showing the usage of "SubscriptionPolicy_AddUpdatePolicyForTenant" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantPolicyResource created on azure
// for more information of creating TenantPolicyResource, please refer to the document of TenantPolicyResource
ResourceIdentifier tenantPolicyResourceId = TenantPolicyResource.CreateResourceIdentifier();
TenantPolicyResource tenantPolicy = client.GetTenantPolicyResource(tenantPolicyResourceId);

// invoke the operation
TenantPolicyCreateOrUpdateContent content = new TenantPolicyCreateOrUpdateContent
{
    BlockSubscriptionsLeavingTenant = true,
    BlockSubscriptionsIntoTenant = true,
    ExemptedPrincipals = { Guid.Parse("e879cf0f-2b4d-5431-109a-f72fc9868693"), Guid.Parse("9792da87-c97b-410d-a97d-27021ba09ce6") },
};
ArmOperation<TenantPolicyResource> lro = await tenantPolicy.CreateOrUpdateAsync(WaitUntil.Completed, content);
TenantPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TenantPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
