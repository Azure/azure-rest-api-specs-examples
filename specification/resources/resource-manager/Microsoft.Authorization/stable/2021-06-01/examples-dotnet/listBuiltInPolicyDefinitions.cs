using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/listBuiltInPolicyDefinitions.json
// this example is just showing the usage of "PolicyDefinitions_ListBuiltIn" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenant = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this TenantPolicyDefinitionResource
TenantPolicyDefinitionCollection collection = tenant.GetTenantPolicyDefinitions();

// invoke the operation and iterate over the result
await foreach (TenantPolicyDefinitionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PolicyDefinitionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
