using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.StreamAnalytics;
using Azure.ResourceManager.StreamAnalytics.Models;

// Generated from example definition: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Subscription_TestQuery.json
// this example is just showing the usage of "Subscriptions_TestQuery" operation, for the dependent resources, they will have to be created separately.

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
StreamAnalyticsTestQuery testQuery = new StreamAnalyticsTestQuery(new StreamingJobData(new AzureLocation("West US"))
{
    SkuName = StreamAnalyticsSkuName.Standard,
    EventsOutOfOrderPolicy = EventsOutOfOrderPolicy.Drop,
    OutputErrorPolicy = StreamingJobOutputErrorPolicy.Drop,
    EventsOutOfOrderMaxDelayInSeconds = 0,
    EventsLateArrivalMaxDelayInSeconds = 5,
    DataLocalion = new AzureLocation("en-US"),
    CompatibilityLevel = StreamingJobCompatibilityLevel.Level1_0,
    Inputs =
    {
    new StreamingJobInputData()
    {
    Properties = new StreamInputProperties()
    {
    Datasource = new RawStreamInputDataSource()
    {
    PayloadUri = new Uri("http://myinput.com"),
    },
    Serialization = new JsonFormatSerialization()
    {
    Encoding = StreamAnalyticsDataSerializationEncoding.Utf8,
    },
    },
    Name = "inputtest",
    }
    },
    Transformation = new StreamingJobTransformationData()
    {
        StreamingUnits = 1,
        Query = "Select Id, Name from inputtest",
        Name = "transformationtest",
    },
    Outputs =
    {
    new StreamingJobOutputData()
    {
    Datasource = new RawOutputDatasource()
    {
    PayloadUri = new Uri("http://myoutput.com"),
    },
    Serialization = new JsonFormatSerialization(),
    Name = "outputtest",
    }
    },
    Functions =
    {
    },
    Tags =
    {
    ["key1"] = "value1",
    ["key3"] = "value3",
    ["randomKey"] = "randomValue",
    },
})
{
    WriteUri = new Uri("http://myoutput.com"),
    Path = "/pathto/subdirectory",
};
ArmOperation<StreamAnalyticsQueryTestingResult> lro = await subscriptionResource.TestQuerySubscriptionAsync(WaitUntil.Completed, location, testQuery);
StreamAnalyticsQueryTestingResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
