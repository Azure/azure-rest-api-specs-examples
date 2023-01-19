using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deleteJobSchedule.json
// this example is just showing the usage of "JobSchedule_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationJobScheduleResource created on azure
// for more information of creating AutomationJobScheduleResource, please refer to the document of AutomationJobScheduleResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "ContoseAutomationAccount";
Guid jobScheduleId = Guid.Parse("0fa462ba-3aa2-4138-83ca-9ebc3bc55cdc");
ResourceIdentifier automationJobScheduleResourceId = AutomationJobScheduleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, jobScheduleId);
AutomationJobScheduleResource automationJobSchedule = client.GetAutomationJobScheduleResource(automationJobScheduleResourceId);

// invoke the operation
await automationJobSchedule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
