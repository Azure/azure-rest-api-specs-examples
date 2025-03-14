using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automation.Models;
using Azure.ResourceManager.Automation;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControlSyncJob/getSourceControlSyncJob.json
// this example is just showing the usage of "SourceControlSyncJob_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationSourceControlResource created on azure
// for more information of creating AutomationSourceControlResource, please refer to the document of AutomationSourceControlResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "myAutomationAccount33";
string sourceControlName = "MySourceControl";
ResourceIdentifier automationSourceControlResourceId = AutomationSourceControlResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, sourceControlName);
AutomationSourceControlResource automationSourceControl = client.GetAutomationSourceControlResource(automationSourceControlResourceId);

// invoke the operation
Guid sourceControlSyncJobId = Guid.Parse("ce6fe3e3-9db3-4096-a6b4-82bfb4c10a9a");
SourceControlSyncJobResult result = await automationSourceControl.GetSourceControlSyncJobAsync(sourceControlSyncJobId);

Console.WriteLine($"Succeeded: {result}");
