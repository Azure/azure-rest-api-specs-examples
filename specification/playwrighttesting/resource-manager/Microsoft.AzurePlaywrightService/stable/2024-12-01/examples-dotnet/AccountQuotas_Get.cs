using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PlaywrightTesting.Models;
using Azure.ResourceManager.PlaywrightTesting;

// Generated from example definition: 2024-12-01/AccountQuotas_Get.json
// this example is just showing the usage of "AccountQuota_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PlaywrightTestingAccountResource created on azure
// for more information of creating PlaywrightTestingAccountResource, please refer to the document of PlaywrightTestingAccountResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "dummyrg";
string accountName = "myPlaywrightAccount";
ResourceIdentifier playwrightTestingAccountResourceId = PlaywrightTestingAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
PlaywrightTestingAccountResource playwrightTestingAccount = client.GetPlaywrightTestingAccountResource(playwrightTestingAccountResourceId);

// get the collection of this PlaywrightTestingAccountQuotaResource
PlaywrightTestingAccountQuotaCollection collection = playwrightTestingAccount.GetAllPlaywrightTestingAccountQuota();

// invoke the operation
PlaywrightTestingQuotaName quotaName = PlaywrightTestingQuotaName.ScalableExecution;
NullableResponse<PlaywrightTestingAccountQuotaResource> response = await collection.GetIfExistsAsync(quotaName);
PlaywrightTestingAccountQuotaResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PlaywrightTestingAccountQuotaData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
