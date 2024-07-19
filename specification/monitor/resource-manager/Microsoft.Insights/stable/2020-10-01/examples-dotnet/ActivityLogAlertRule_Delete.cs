using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_Delete.json
// this example is just showing the usage of "ActivityLogAlerts_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ActivityLogAlertResource created on azure
// for more information of creating ActivityLogAlertResource, please refer to the document of ActivityLogAlertResource
string subscriptionId = "187f412d-1758-44d9-b052-169e2564721d";
string resourceGroupName = "MyResourceGroup";
string activityLogAlertName = "SampleActivityLogAlertRule";
ResourceIdentifier activityLogAlertResourceId = ActivityLogAlertResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, activityLogAlertName);
ActivityLogAlertResource activityLogAlert = client.GetActivityLogAlertResource(activityLogAlertResourceId);

// invoke the operation
await activityLogAlert.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
