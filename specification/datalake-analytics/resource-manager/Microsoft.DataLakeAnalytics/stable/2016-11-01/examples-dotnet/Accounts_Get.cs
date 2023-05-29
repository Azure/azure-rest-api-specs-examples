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

// Generated from example definition: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/Accounts_Get.json
// this example is just showing the usage of "Accounts_Get" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
DataLakeAnalyticsAccountResource result = await dataLakeAnalyticsAccount.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataLakeAnalyticsAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
