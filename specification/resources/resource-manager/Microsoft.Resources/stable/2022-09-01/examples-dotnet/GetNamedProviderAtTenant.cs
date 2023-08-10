using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Resources/stable/2022-09-01/examples/GetNamedProviderAtTenant.json
// this example is just showing the usage of "Providers_GetAtTenantScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenant = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
string resourceProviderNamespace = "Microsoft.Storage";
string expand = "resourceTypes/aliases";
TenantResourceProvider result = await tenant.GetTenantResourceProviderAsync(resourceProviderNamespace, expand: expand);

Console.WriteLine($"Succeeded: {result}");
