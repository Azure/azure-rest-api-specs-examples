using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApplicationInsights.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ApplicationInsights;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/examples/ComponentsPurge.json
// this example is just showing the usage of "Components_Purge" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApplicationInsightsComponentResource created on azure
// for more information of creating ApplicationInsightsComponentResource, please refer to the document of ApplicationInsightsComponentResource
string subscriptionId = "00000000-0000-0000-0000-00000000000";
string resourceGroupName = "OIAutoRest5123";
string resourceName = "aztest5048";
ResourceIdentifier applicationInsightsComponentResourceId = ApplicationInsightsComponentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
ApplicationInsightsComponentResource applicationInsightsComponent = client.GetApplicationInsightsComponentResource(applicationInsightsComponentResourceId);

// invoke the operation
ComponentPurgeContent content = new ComponentPurgeContent("Heartbeat", new ComponentPurgeFilters[]
{
new ComponentPurgeFilters()
{
Column = "TimeGenerated",
Operator = ">",
Value = BinaryData.FromString("\"2017-09-01T00:00:00\""),
}
});
ComponentPurgeResult result = await applicationInsightsComponent.PurgeAsync(content);

Console.WriteLine($"Succeeded: {result}");
