using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Authorization;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2015-07-01/examples/ElevateAccess.json
// this example is just showing the usage of "GlobalAdministrator_ElevateAccess" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
await tenantResource.ElevateAccessGlobalAdministratorAsync();

Console.WriteLine("Succeeded");
