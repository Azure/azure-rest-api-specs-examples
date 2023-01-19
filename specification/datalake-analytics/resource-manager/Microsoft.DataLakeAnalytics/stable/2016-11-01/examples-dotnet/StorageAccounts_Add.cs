using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataLakeAnalytics;
using Azure.ResourceManager.DataLakeAnalytics.Models;

// Generated from example definition: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/StorageAccounts_Add.json
// this example is just showing the usage of "StorageAccounts_Add" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataLakeAnalyticsAccountResource created on azure
// for more information of creating DataLakeAnalyticsAccountResource, please refer to the document of DataLakeAnalyticsAccountResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "contosorg";
string accountName = "contosoadla";
ResourceIdentifier dataLakeAnalyticsAccountResourceId = DataLakeAnalyticsAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
DataLakeAnalyticsAccountResource dataLakeAnalyticsAccount = client.GetDataLakeAnalyticsAccountResource(dataLakeAnalyticsAccountResourceId);

// get the collection of this DataLakeAnalyticsStorageAccountInformationResource
DataLakeAnalyticsStorageAccountInformationCollection collection = dataLakeAnalyticsAccount.GetAllDataLakeAnalyticsStorageAccountInformation();

// invoke the operation
string storageAccountName = "test_storage";
DataLakeAnalyticsStorageAccountInformationCreateOrUpdateContent content = new DataLakeAnalyticsStorageAccountInformationCreateOrUpdateContent("34adfa4f-cedf-4dc0-ba29-b6d1a69ab346")
{
    Suffix = "test_suffix",
};
await collection.CreateOrUpdateAsync(WaitUntil.Completed, storageAccountName, content);

Console.WriteLine($"Succeeded");
