using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutionsAnalytics/PostIoTSecuritySolutionsSecurityAggregatedAlertDismiss.json
// this example is just showing the usage of "IotSecuritySolutionsAnalyticsAggregatedAlert_Dismiss" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotSecurityAggregatedAlertResource created on azure
// for more information of creating IotSecurityAggregatedAlertResource, please refer to the document of IotSecurityAggregatedAlertResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string resourceGroupName = "IoTEdgeResources";
string solutionName = "default";
string aggregatedAlertName = "IoT_Bruteforce_Fail/2019-02-02/dismiss";
ResourceIdentifier iotSecurityAggregatedAlertResourceId = IotSecurityAggregatedAlertResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, solutionName, aggregatedAlertName);
IotSecurityAggregatedAlertResource iotSecurityAggregatedAlert = client.GetIotSecurityAggregatedAlertResource(iotSecurityAggregatedAlertResourceId);

// invoke the operation
await iotSecurityAggregatedAlert.DismissAsync();

Console.WriteLine($"Succeeded");
