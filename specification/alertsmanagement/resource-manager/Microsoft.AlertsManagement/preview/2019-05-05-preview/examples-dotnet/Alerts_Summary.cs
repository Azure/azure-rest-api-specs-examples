using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AlertsManagement;
using Azure.ResourceManager.AlertsManagement.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Alerts_Summary.json
// this example is just showing the usage of "Alerts_GetSummary" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "1e3ff1c0-771a-4119-a03b-be82a51e232d";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
SubscriptionResourceGetServiceAlertSummaryOptions options = new SubscriptionResourceGetServiceAlertSummaryOptions(groupby: new AlertsSummaryGroupByField("severity,alertState")) { };
ServiceAlertSummary result = await subscriptionResource.GetServiceAlertSummaryAsync(options);

Console.WriteLine($"Succeeded: {result}");
