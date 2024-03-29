using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Quota;
using Azure.ResourceManager.Quota.Models;

// Generated from example definition: specification/quota/resource-manager/Microsoft.Quota/preview/2021-03-15-preview/examples/getComputeOneSkuQuotaLimit.json
// this example is just showing the usage of "Quota_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this CurrentQuotaLimitBaseResource
string scope = "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
CurrentQuotaLimitBaseCollection collection = client.GetCurrentQuotaLimitBases(scopeId);

// invoke the operation
string resourceName = "standardNDSFamily";
bool result = await collection.ExistsAsync(resourceName);

Console.WriteLine($"Succeeded: {result}");
