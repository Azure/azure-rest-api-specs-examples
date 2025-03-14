using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataLakeAnalytics.Models;
using Azure.ResourceManager.DataLakeAnalytics;

// Generated from example definition: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/examples/ComputePolicies_Delete.json
// this example is just showing the usage of "ComputePolicies_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataLakeAnalyticsComputePolicyResource created on azure
// for more information of creating DataLakeAnalyticsComputePolicyResource, please refer to the document of DataLakeAnalyticsComputePolicyResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "contosorg";
string accountName = "contosoadla";
string computePolicyName = "test_policy";
ResourceIdentifier dataLakeAnalyticsComputePolicyResourceId = DataLakeAnalyticsComputePolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, computePolicyName);
DataLakeAnalyticsComputePolicyResource dataLakeAnalyticsComputePolicy = client.GetDataLakeAnalyticsComputePolicyResource(dataLakeAnalyticsComputePolicyResourceId);

// invoke the operation
await dataLakeAnalyticsComputePolicy.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
