using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DataShare;

// Generated from example definition: specification/datashare/resource-manager/Microsoft.DataShare/stable/2021-08-01/examples/ConsumerInvitations_Get.json
// this example is just showing the usage of "ConsumerInvitations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this DataShareConsumerInvitationResource
DataShareConsumerInvitationCollection collection = tenantResource.GetDataShareConsumerInvitations();

// invoke the operation
AzureLocation location = new AzureLocation("East US 2");
Guid invitationId = Guid.Parse("dfbbc788-19eb-4607-a5a1-c74181bfff03");
NullableResponse<DataShareConsumerInvitationResource> response = await collection.GetIfExistsAsync(location, invitationId);
DataShareConsumerInvitationResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DataShareConsumerInvitationData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
