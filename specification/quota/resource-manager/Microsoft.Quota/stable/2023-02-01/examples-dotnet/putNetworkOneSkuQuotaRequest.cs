using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Quota;
using Azure.ResourceManager.Quota.Models;

// Generated from example definition: specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/putNetworkOneSkuQuotaRequest.json
// this example is just showing the usage of "Quota_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this CurrentQuotaLimitBaseResource
string scope = "subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Network/locations/eastus";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
CurrentQuotaLimitBaseCollection collection = client.GetCurrentQuotaLimitBases(scopeId);

// invoke the operation
string resourceName = "MinPublicIpInterNetworkPrefixLength";
CurrentQuotaLimitBaseData data = new CurrentQuotaLimitBaseData()
{
    Properties = new QuotaProperties()
    {
        Limit = new LimitObject(10),
        Name = new ResourceName()
        {
            Value = "MinPublicIpInterNetworkPrefixLength",
        },
        ResourceType = "MinPublicIpInterNetworkPrefixLength",
    },
};
ArmOperation<CurrentQuotaLimitBaseResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data);
CurrentQuotaLimitBaseResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CurrentQuotaLimitBaseData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
