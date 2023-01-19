using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/stopJob.json
// this example is just showing the usage of "Job_Stop" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationJobResource created on azure
// for more information of creating AutomationJobResource, please refer to the document of AutomationJobResource
string subscriptionId = "51766542-3ed7-4a72-a187-0c8ab644ddab";
string resourceGroupName = "mygroup";
string automationAccountName = "ContoseAutomationAccount";
string jobName = "foo";
ResourceIdentifier automationJobResourceId = AutomationJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, jobName);
AutomationJobResource automationJob = client.GetAutomationJobResource(automationJobResourceId);

// invoke the operation
await automationJob.StopAsync();

Console.WriteLine($"Succeeded");
