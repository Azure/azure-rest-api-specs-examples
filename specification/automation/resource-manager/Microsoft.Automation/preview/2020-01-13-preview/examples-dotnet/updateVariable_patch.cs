using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automation.Models;
using Azure.ResourceManager.Automation;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateVariable_patch.json
// this example is just showing the usage of "Variable_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationVariableResource created on azure
// for more information of creating AutomationVariableResource, please refer to the document of AutomationVariableResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "sampleAccount9";
string variableName = "sampleVariable";
ResourceIdentifier automationVariableResourceId = AutomationVariableResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, variableName);
AutomationVariableResource automationVariable = client.GetAutomationVariableResource(automationVariableResourceId);

// invoke the operation
AutomationVariablePatch patch = new AutomationVariablePatch
{
    Name = "sampleVariable",
    Value = "\"ComputerName3.domain.com\"",
};
AutomationVariableResource result = await automationVariable.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutomationVariableData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
