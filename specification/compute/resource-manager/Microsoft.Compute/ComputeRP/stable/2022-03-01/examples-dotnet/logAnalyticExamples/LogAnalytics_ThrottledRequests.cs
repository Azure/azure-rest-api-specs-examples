using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/logAnalyticExamples/LogAnalytics_ThrottledRequests.json
// this example is just showing the usage of "LogAnalytics_ExportThrottledRequests" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "{subscription-id}";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("westus");
ThrottledRequestsContent content = new ThrottledRequestsContent(new Uri("https://somesasuri"), DateTimeOffset.Parse("2018-01-21T01:54:06.862601Z"), DateTimeOffset.Parse("2018-01-23T01:54:06.862601Z"))
{
    GroupByOperationName = true,
    GroupByResourceName = false,
    GroupByClientApplicationId = false,
    GroupByUserAgent = false,
};
ArmOperation<LogAnalytics> lro = await subscriptionResource.ExportLogAnalyticsThrottledRequestsAsync(WaitUntil.Completed, location, content);
LogAnalytics result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
