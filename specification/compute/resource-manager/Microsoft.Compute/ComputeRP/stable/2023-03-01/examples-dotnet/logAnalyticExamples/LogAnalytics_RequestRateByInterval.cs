using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/logAnalyticExamples/LogAnalytics_RequestRateByInterval.json
// this example is just showing the usage of "LogAnalytics_ExportRequestRateByInterval" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "{subscription-id}";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("westus");
RequestRateByIntervalContent content = new RequestRateByIntervalContent(new Uri("https://somesasuri"), DateTimeOffset.Parse("2018-01-21T01:54:06.862601Z"), DateTimeOffset.Parse("2018-01-23T01:54:06.862601Z"), IntervalInMins.FiveMins)
{
    GroupByResourceName = true,
};
ArmOperation<LogAnalytics> lro = await subscriptionResource.ExportLogAnalyticsRequestRateByIntervalAsync(WaitUntil.Completed, location, content);
LogAnalytics result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
