using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-09-01/examples/getDataPolicyManifest.json
// this example is just showing the usage of "DataPolicyManifests_GetByPolicyMode" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenant = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this DataPolicyManifestResource
DataPolicyManifestCollection collection = tenant.GetDataPolicyManifests();

// invoke the operation
string policyMode = "Microsoft.KeyVault.Data";
NullableResponse<DataPolicyManifestResource> response = await collection.GetIfExistsAsync(policyMode);
DataPolicyManifestResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DataPolicyManifestData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
