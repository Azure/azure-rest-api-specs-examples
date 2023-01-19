using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.StreamAnalytics;
using Azure.ResourceManager.StreamAnalytics.Models;

// Generated from example definition: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Subscription_CompileQuery.json
// this example is just showing the usage of "Subscriptions_CompileQuery" operation, for the dependent resources, they will have to be created separately.

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
StreamAnalyticsCompileQuery compileQuery = new StreamAnalyticsCompileQuery("SELECT\r\n    *\r\nINTO\r\n    [output1]\r\nFROM\r\n    [input1]", StreamingJobType.Cloud)
{
    Inputs =
    {
    new StreamAnalyticsQueryInput("input1","Stream")
    },
    Functions =
    {
    new StreamAnalyticsQueryFunction("function1","Scalar","Microsoft.StreamAnalytics/JavascriptUdf",new StreamingJobFunctionInput[]
    {
    new StreamingJobFunctionInput()
    {
    DataType = "any",
    IsConfigurationParameter = null,
    }
    },new StreamingJobFunctionOutput()
    {
    DataType = "bigint",
    })
    },
    CompatibilityLevel = StreamingJobCompatibilityLevel.Level1_2,
};
StreamAnalyticsQueryCompilationResult result = await subscriptionResource.CompileQuerySubscriptionAsync(location, compileQuery);

Console.WriteLine($"Succeeded: {result}");
