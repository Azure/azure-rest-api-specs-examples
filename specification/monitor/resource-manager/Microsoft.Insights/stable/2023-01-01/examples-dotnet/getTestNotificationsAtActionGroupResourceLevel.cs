using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2023-01-01/examples/getTestNotificationsAtActionGroupResourceLevel.json
// this example is just showing the usage of "ActionGroups_GetTestNotificationsAtActionGroupResourceLevel" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ActionGroupResource created on azure
// for more information of creating ActionGroupResource, please refer to the document of ActionGroupResource
string subscriptionId = "11111111-1111-1111-1111-111111111111";
string resourceGroupName = "TestRgName";
string actionGroupName = "TestAgName";
ResourceIdentifier actionGroupResourceId = ActionGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, actionGroupName);
ActionGroupResource actionGroup = client.GetActionGroupResource(actionGroupResourceId);

// invoke the operation
string notificationId = "11000222191287";
NotificationStatus result = await actionGroup.GetNotificationStatusAsync(notificationId);

Console.WriteLine($"Succeeded: {result}");
