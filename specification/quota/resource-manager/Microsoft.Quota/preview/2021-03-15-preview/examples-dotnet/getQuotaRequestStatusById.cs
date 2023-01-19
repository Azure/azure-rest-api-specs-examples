using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Quota;

// Generated from example definition: specification/quota/resource-manager/Microsoft.Quota/preview/2021-03-15-preview/examples/getQuotaRequestStatusById.json
// this example is just showing the usage of "QuotaRequestStatus_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this QuotaRequestDetailResource
string scope = "subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Compute/locations/eastus";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
QuotaRequestDetailCollection collection = client.GetQuotaRequestDetails(scopeId);

// invoke the operation
string id = "2B5C8515-37D8-4B6A-879B-CD641A2CF605";
bool result = await collection.ExistsAsync(id);

Console.WriteLine($"Succeeded: {result}");
