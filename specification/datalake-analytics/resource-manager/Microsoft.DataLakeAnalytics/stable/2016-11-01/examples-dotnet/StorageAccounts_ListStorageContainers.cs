using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataLakeAnalytics;

// Generated from example definition: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/StorageAccounts_ListStorageContainers.json
// this example is just showing the usage of "StorageAccounts_ListStorageContainers" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataLakeAnalyticsStorageAccountInformationResource created on azure
// for more information of creating DataLakeAnalyticsStorageAccountInformationResource, please refer to the document of DataLakeAnalyticsStorageAccountInformationResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "contosorg";
string accountName = "contosoadla";
string storageAccountName = "test_storage";
ResourceIdentifier dataLakeAnalyticsStorageAccountInformationResourceId = DataLakeAnalyticsStorageAccountInformationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, storageAccountName);
DataLakeAnalyticsStorageAccountInformationResource dataLakeAnalyticsStorageAccountInformation = client.GetDataLakeAnalyticsStorageAccountInformationResource(dataLakeAnalyticsStorageAccountInformationResourceId);

// get the collection of this DataLakeAnalyticsStorageContainerResource
DataLakeAnalyticsStorageContainerCollection collection = dataLakeAnalyticsStorageAccountInformation.GetDataLakeAnalyticsStorageContainers();

// invoke the operation and iterate over the result
await foreach (DataLakeAnalyticsStorageContainerResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DataLakeAnalyticsStorageContainerData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
