using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataLakeAnalytics;
using Azure.ResourceManager.DataLakeAnalytics.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/Accounts_Create.json
// this example is just showing the usage of "Accounts_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "contosorg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DataLakeAnalyticsAccountResource
DataLakeAnalyticsAccountCollection collection = resourceGroupResource.GetDataLakeAnalyticsAccounts();

// invoke the operation
string accountName = "contosoadla";
DataLakeAnalyticsAccountCreateOrUpdateContent content = new DataLakeAnalyticsAccountCreateOrUpdateContent(new AzureLocation("eastus2"), "test_adls", new DataLakeStoreForDataLakeAnalyticsAccountCreateOrUpdateContent[]
{
new DataLakeStoreForDataLakeAnalyticsAccountCreateOrUpdateContent("test_adls")
{
Suffix = "test_suffix",
}
})
{
    Tags =
    {
    ["test_key"] = "test_value",
    },
    StorageAccounts =
    {
    new StorageAccountForDataLakeAnalyticsAccountCreateOrUpdateContent("test_storage","34adfa4f-cedf-4dc0-ba29-b6d1a69ab346")
    {
    Suffix = "test_suffix",
    }
    },
    ComputePolicies =
    {
    new ComputePolicyForDataLakeAnalyticsAccountCreateOrUpdateContent("test_policy",Guid.Parse("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),AadObjectIdentifierType.User)
    {
    MaxDegreeOfParallelismPerJob = 1,
    MinPriorityPerJob = 1,
    }
    },
    FirewallRules =
    {
    new FirewallRuleForDataLakeAnalyticsAccountCreateOrUpdateContent("test_rule",IPAddress.Parse("1.1.1.1"),IPAddress.Parse("2.2.2.2"))
    },
    FirewallState = DataLakeAnalyticsFirewallState.Enabled,
    FirewallAllowAzureIPs = DataLakeAnalyticsFirewallAllowAzureIPsState.Enabled,
    NewTier = DataLakeAnalyticsCommitmentTierType.Consumption,
    MaxJobCount = 3,
    MaxDegreeOfParallelism = 30,
    MaxDegreeOfParallelismPerJob = 1,
    MinPriorityPerJob = 1,
    QueryStoreRetention = 30,
};
ArmOperation<DataLakeAnalyticsAccountResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, accountName, content);
DataLakeAnalyticsAccountResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataLakeAnalyticsAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
