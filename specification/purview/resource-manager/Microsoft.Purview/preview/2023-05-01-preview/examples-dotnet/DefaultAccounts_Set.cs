using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Purview;
using Azure.ResourceManager.Purview.Models;

// Generated from example definition: specification/purview/resource-manager/Microsoft.Purview/preview/2023-05-01-preview/examples/DefaultAccounts_Set.json
// this example is just showing the usage of "DefaultAccounts_Set" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
DefaultPurviewAccountPayload defaultAccountPayload = new DefaultPurviewAccountPayload()
{
    AccountName = "myDefaultAccount",
    ResourceGroupName = "rg-1",
    Scope = "12345678-1234-1234-1234-12345678abcd",
    ScopeTenantId = Guid.Parse("12345678-1234-1234-1234-12345678abcd"),
    ScopeType = PurviewAccountScopeType.Tenant,
    SubscriptionId = "12345678-1234-1234-1234-12345678aaaa",
};
DefaultPurviewAccountPayload result = await tenantResource.SetDefaultAccountAsync(defaultAccountPayload);

Console.WriteLine($"Succeeded: {result}");
