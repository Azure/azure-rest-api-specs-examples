using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Purview.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Purview;

// Generated from example definition: specification/purview/resource-manager/Microsoft.Purview/preview/2023-05-01-preview/examples/DefaultAccounts_Remove.json
// this example is just showing the usage of "DefaultAccounts_Remove" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
Guid scopeTenantId = Guid.Parse("ee85a74c-405e-4adc-bb47-ffa8ca0c9f31");
PurviewAccountScopeType scopeType = PurviewAccountScopeType.Tenant;
string scope = "12345678-1234-1234-1234-12345678abcd";
await tenantResource.RemoveDefaultAccountAsync(scopeTenantId, scopeType, scope: scope);

Console.WriteLine("Succeeded");
