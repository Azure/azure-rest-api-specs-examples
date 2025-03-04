using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Quota.Models;
using Azure.ResourceManager.Quota;

// Generated from example definition: specification/quota/resource-manager/Microsoft.Quota/stable/2025-03-01/examples/getNetworkOneSkuQuotaLimit.json
// this example is just showing the usage of "Quota_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this CurrentQuotaLimitBaseResource
string scope = "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Network/locations/eastus";
CurrentQuotaLimitBaseCollection collection = client.GetCurrentQuotaLimitBases(new ResourceIdentifier(scope));

// invoke the operation
string resourceName = "MinPublicIpInterNetworkPrefixLength";
NullableResponse<CurrentQuotaLimitBaseResource> response = await collection.GetIfExistsAsync(resourceName);
CurrentQuotaLimitBaseResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CurrentQuotaLimitBaseData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
