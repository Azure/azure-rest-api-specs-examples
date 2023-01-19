using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateModule.json
// this example is just showing the usage of "Module_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationAccountModuleResource created on azure
// for more information of creating AutomationAccountModuleResource, please refer to the document of AutomationAccountModuleResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "MyAutomationAccount";
string moduleName = "MyModule";
ResourceIdentifier automationAccountModuleResourceId = AutomationAccountModuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, moduleName);
AutomationAccountModuleResource automationAccountModule = client.GetAutomationAccountModuleResource(automationAccountModuleResourceId);

// invoke the operation
AutomationAccountModulePatch patch = new AutomationAccountModulePatch()
{
    ContentLink = new AutomationContentLink()
    {
        Uri = new Uri("https://teststorage.blob.core.windows.net/mycontainer/MyModule.zip"),
        ContentHash = new AutomationContentHash("sha265", "07E108A962B81DD9C9BAA89BB47C0F6EE52B29E83758B07795E408D258B2B87A"),
        Version = "1.0.0.0",
    },
};
AutomationAccountModuleResource result = await automationAccountModule.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutomationModuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
