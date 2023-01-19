using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataLakeAnalytics;
using Azure.ResourceManager.DataLakeAnalytics.Models;

// Generated from example definition: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/StorageAccounts_ListSasTokens.json
// this example is just showing the usage of "StorageAccounts_ListSasTokens" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataLakeAnalyticsStorageContainerResource created on azure
// for more information of creating DataLakeAnalyticsStorageContainerResource, please refer to the document of DataLakeAnalyticsStorageContainerResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "contosorg";
string accountName = "contosoadla";
string storageAccountName = "test_storage";
string containerName = "test_container";
ResourceIdentifier dataLakeAnalyticsStorageContainerResourceId = DataLakeAnalyticsStorageContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, storageAccountName, containerName);
DataLakeAnalyticsStorageContainerResource dataLakeAnalyticsStorageContainer = client.GetDataLakeAnalyticsStorageContainerResource(dataLakeAnalyticsStorageContainerResourceId);

// invoke the operation and iterate over the result
await foreach (DataLakeAnalyticsSasTokenInformation item in dataLakeAnalyticsStorageContainer.GetSasTokensAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
