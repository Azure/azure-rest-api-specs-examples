using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataLakeAnalytics;
using Azure.ResourceManager.DataLakeAnalytics.Models;

// Generated from example definition: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/DataLakeStoreAccounts_Add.json
// this example is just showing the usage of "DataLakeStoreAccounts_Add" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataLakeStoreAccountInformationResource created on azure
// for more information of creating DataLakeStoreAccountInformationResource, please refer to the document of DataLakeStoreAccountInformationResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "contosorg";
string accountName = "contosoadla";
string dataLakeStoreAccountName = "test_adls_account";
ResourceIdentifier dataLakeStoreAccountInformationResourceId = DataLakeStoreAccountInformationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, dataLakeStoreAccountName);
DataLakeStoreAccountInformationResource dataLakeStoreAccountInformation = client.GetDataLakeStoreAccountInformationResource(dataLakeStoreAccountInformationResourceId);

// invoke the operation
DataLakeStoreAccountInformationCreateOrUpdateContent content = new DataLakeStoreAccountInformationCreateOrUpdateContent()
{
    Suffix = "test_suffix",
};
await dataLakeStoreAccountInformation.UpdateAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");
