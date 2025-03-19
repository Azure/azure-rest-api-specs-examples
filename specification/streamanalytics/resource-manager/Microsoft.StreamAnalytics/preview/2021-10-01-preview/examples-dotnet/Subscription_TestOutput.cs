using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.StreamAnalytics.Models;
using Azure.ResourceManager.StreamAnalytics;

// Generated from example definition: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Subscription_TestOutput.json
// this example is just showing the usage of "Subscriptions_TestOutput" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("West US");
StreamAnalyticsTestOutput testOutput = new StreamAnalyticsTestOutput(new StreamingJobOutputData
{
    Datasource = new BlobOutputDataSource
    {
        StorageAccounts = {new StreamAnalyticsStorageAccount
        {
        AccountName = "someAccountName",
        AccountKey = "accountKey==",
        }},
        Container = "state",
        PathPattern = "{date}/{time}",
        DateFormat = "yyyy/MM/dd",
        TimeFormat = "HH",
    },
    Serialization = new CsvFormatSerialization
    {
        FieldDelimiter = ",",
        Encoding = StreamAnalyticsDataSerializationEncoding.Utf8,
    },
});
ArmOperation<StreamAnalyticsTestDatasourceResult> lro = await subscriptionResource.TestOutputSubscriptionAsync(WaitUntil.Completed, location, testOutput);
StreamAnalyticsTestDatasourceResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
