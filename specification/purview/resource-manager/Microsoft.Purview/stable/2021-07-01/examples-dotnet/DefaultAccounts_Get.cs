using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Purview;
using Azure.ResourceManager.Purview.Models;

// Generated from example definition: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/DefaultAccounts_Get.json
// this example is just showing the usage of "DefaultAccounts_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
Guid scopeTenantId = Guid.Parse("11733A4E-BA84-46FF-91D1-AFF1A3215A90");
PurviewAccountScopeType scopeType = PurviewAccountScopeType.Tenant;
string scope = "11733A4E-BA84-46FF-91D1-AFF1A3215A90";
DefaultPurviewAccountPayload result = await tenantResource.GetDefaultAccountAsync(scopeTenantId, scopeType, scope: scope);

Console.WriteLine($"Succeeded: {result}");
