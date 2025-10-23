using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PolicyInsights.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.PolicyInsights;

// Generated from example definition: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/PolicyMetadata_List.json
// this example is just showing the usage of "PolicyMetadata_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this PolicyMetadataResource
PolicyMetadataCollection collection = tenantResource.GetAllPolicyMetadata();

// invoke the operation and iterate over the result
await foreach (SlimPolicyMetadata item in collection.GetAllAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
