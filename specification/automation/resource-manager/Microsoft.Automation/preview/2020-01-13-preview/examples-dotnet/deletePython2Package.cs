using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deletePython2Package.json
// this example is just showing the usage of "Python2Package_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationAccountPython2PackageResource created on azure
// for more information of creating AutomationAccountPython2PackageResource, please refer to the document of AutomationAccountPython2PackageResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "myAutomationAccount33";
string packageName = "OmsCompositeResources";
ResourceIdentifier automationAccountPython2PackageResourceId = AutomationAccountPython2PackageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, packageName);
AutomationAccountPython2PackageResource automationAccountPython2Package = client.GetAutomationAccountPython2PackageResource(automationAccountPython2PackageResourceId);

// invoke the operation
await automationAccountPython2Package.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
