using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Quota.Models;
using Azure.ResourceManager.Quota;

// Generated from example definition: specification/quota/resource-manager/Microsoft.Quota/stable/2025-03-01/examples/patchNetworkOneSkuQuotaRequest.json
// this example is just showing the usage of "Quota_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CurrentQuotaLimitBaseResource created on azure
// for more information of creating CurrentQuotaLimitBaseResource, please refer to the document of CurrentQuotaLimitBaseResource
string scope = "subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Network/locations/eastus";
string resourceName = "MinPublicIpInterNetworkPrefixLength";
ResourceIdentifier currentQuotaLimitBaseResourceId = CurrentQuotaLimitBaseResource.CreateResourceIdentifier(scope, resourceName);
CurrentQuotaLimitBaseResource currentQuotaLimitBase = client.GetCurrentQuotaLimitBaseResource(currentQuotaLimitBaseResourceId);

// invoke the operation
CurrentQuotaLimitBaseData data = new CurrentQuotaLimitBaseData
{
    Properties = new QuotaProperties
    {
        Limit = new QuotaLimitObject(10),
        Name = new QuotaRequestResourceName
        {
            Value = "MinPublicIpInterNetworkPrefixLength",
        },
        ResourceTypeName = "MinPublicIpInterNetworkPrefixLength",
    },
};
ArmOperation<CurrentQuotaLimitBaseResource> lro = await currentQuotaLimitBase.UpdateAsync(WaitUntil.Completed, data);
CurrentQuotaLimitBaseResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CurrentQuotaLimitBaseData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
