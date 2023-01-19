using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfigurationRun/getSoftwareUpdateConfigurationRunById.json
// this example is just showing the usage of "SoftwareUpdateConfigurationRuns_GetById" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationAccountResource created on azure
// for more information of creating AutomationAccountResource, please refer to the document of AutomationAccountResource
string subscriptionId = "51766542-3ed7-4a72-a187-0c8ab644ddab";
string resourceGroupName = "mygroup";
string automationAccountName = "myaccount";
ResourceIdentifier automationAccountResourceId = AutomationAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName);
AutomationAccountResource automationAccount = client.GetAutomationAccountResource(automationAccountResourceId);

// invoke the operation
Guid softwareUpdateConfigurationRunId = Guid.Parse("2bd77cfa-2e9c-41b4-a45b-684a77cfeca9");
SoftwareUpdateConfigurationRun result = await automationAccount.GetSoftwareUpdateConfigurationRunAsync(softwareUpdateConfigurationRunId);

Console.WriteLine($"Succeeded: {result}");
