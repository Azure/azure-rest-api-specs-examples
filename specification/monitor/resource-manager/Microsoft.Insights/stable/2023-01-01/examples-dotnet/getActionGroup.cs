using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2023-01-01/examples/getActionGroup.json
// this example is just showing the usage of "ActionGroups_Get" operation, for the dependent resources, they will have to be created separately.

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
ActionGroupResource result = await actionGroup.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ActionGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
