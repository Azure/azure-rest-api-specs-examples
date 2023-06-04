using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PolicyInsights;
using Azure.ResourceManager.PolicyInsights.Models;

// Generated from example definition: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyMetadata_GetResource.json
// this example is just showing the usage of "PolicyMetadata_GetResource" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this PolicyMetadataResource
PolicyMetadataCollection collection = tenantResource.GetAllPolicyMetadata();

// invoke the operation
string resourceName = "NIST_SP_800-53_R4_AC-2";
bool result = await collection.ExistsAsync(resourceName);

Console.WriteLine($"Succeeded: {result}");
