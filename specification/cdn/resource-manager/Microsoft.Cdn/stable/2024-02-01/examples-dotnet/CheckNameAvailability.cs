using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/CheckNameAvailability.json
// this example is just showing the usage of "CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
CdnNameAvailabilityContent content = new CdnNameAvailabilityContent("sampleName", CdnResourceType.Endpoints);
CdnNameAvailabilityResult result = await tenantResource.CheckCdnNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
