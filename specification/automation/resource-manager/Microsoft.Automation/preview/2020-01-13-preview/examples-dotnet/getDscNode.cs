using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automation;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getDscNode.json
// this example is just showing the usage of "DscNode_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationAccountResource created on azure
// for more information of creating AutomationAccountResource, please refer to the document of AutomationAccountResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "myAutomationAccount33";
ResourceIdentifier automationAccountResourceId = AutomationAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName);
AutomationAccountResource automationAccount = client.GetAutomationAccountResource(automationAccountResourceId);

// get the collection of this DscNodeResource
DscNodeCollection collection = automationAccount.GetDscNodes();

// invoke the operation
string nodeId = "nodeId";
NullableResponse<DscNodeResource> response = await collection.GetIfExistsAsync(nodeId);
DscNodeResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DscNodeData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
