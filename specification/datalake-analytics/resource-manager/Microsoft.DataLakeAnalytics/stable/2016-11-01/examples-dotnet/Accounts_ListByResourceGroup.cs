using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataLakeAnalytics.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DataLakeAnalytics;

// Generated from example definition: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/Accounts_ListByResourceGroup.json
// this example is just showing the usage of "Accounts_ListByResourceGroup" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
DataLakeAnalyticsAccountCollectionGetAllOptions options = new DataLakeAnalyticsAccountCollectionGetAllOptions { Filter = "test_filter", Top = 1, Skip = 1, Select = "test_select", Orderby = "test_orderby", Count = false };
await foreach (DataLakeAnalyticsAccountBasic item in collection.GetAllAsync(options))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
