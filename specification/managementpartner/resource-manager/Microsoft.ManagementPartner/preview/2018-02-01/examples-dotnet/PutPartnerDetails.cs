using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ManagementPartner;

// Generated from example definition: specification/managementpartner/resource-manager/Microsoft.ManagementPartner/preview/2018-02-01/examples/PutPartnerDetails.json
// this example is just showing the usage of "Partner_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this PartnerResponseResource
PartnerResponseCollection collection = tenantResource.GetPartnerResponses();

// invoke the operation
string partnerId = "123456";
ArmOperation<PartnerResponseResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, partnerId);
PartnerResponseResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PartnerResponseData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
