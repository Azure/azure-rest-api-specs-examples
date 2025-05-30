using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/preview/2024-10-01-preview/examples/enableReceiver.json
// this example is just showing the usage of "ActionGroups_EnableReceiver" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ActionGroupResource created on azure
// for more information of creating ActionGroupResource, please refer to the document of ActionGroupResource
string subscriptionId = "187f412d-1758-44d9-b052-169e2564721d";
string resourceGroupName = "Default-NotificationRules";
string actionGroupName = "SampleActionGroup";
ResourceIdentifier actionGroupResourceId = ActionGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, actionGroupName);
ActionGroupResource actionGroup = client.GetActionGroupResource(actionGroupResourceId);

// invoke the operation
ActionGroupEnableContent content = new ActionGroupEnableContent("John Doe's mobile");
await actionGroup.EnableReceiverAsync(content);

Console.WriteLine("Succeeded");
