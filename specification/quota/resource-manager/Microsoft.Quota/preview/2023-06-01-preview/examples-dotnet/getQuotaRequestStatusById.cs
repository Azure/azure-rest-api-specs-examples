using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Quota;

// Generated from example definition: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/getQuotaRequestStatusById.json
// this example is just showing the usage of "QuotaRequestStatus_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this QuotaRequestDetailResource created on azure
// for more information of creating QuotaRequestDetailResource, please refer to the document of QuotaRequestDetailResource
string scope = "subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Compute/locations/eastus";
string id = "2B5C8515-37D8-4B6A-879B-CD641A2CF605";
ResourceIdentifier quotaRequestDetailResourceId = QuotaRequestDetailResource.CreateResourceIdentifier(scope, id);
QuotaRequestDetailResource quotaRequestDetail = client.GetQuotaRequestDetailResource(quotaRequestDetailResourceId);

// invoke the operation
QuotaRequestDetailResource result = await quotaRequestDetail.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
QuotaRequestDetailData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
