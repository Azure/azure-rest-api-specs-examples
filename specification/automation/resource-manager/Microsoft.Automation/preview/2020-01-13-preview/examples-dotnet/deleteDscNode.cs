using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deleteDscNode.json
// this example is just showing the usage of "DscNode_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DscNodeResource created on azure
// for more information of creating DscNodeResource, please refer to the document of DscNodeResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "myAutomationAccount9";
string nodeId = "e1243a76-a9bd-432f-bde3-ad8f317ee786";
ResourceIdentifier dscNodeResourceId = DscNodeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, nodeId);
DscNodeResource dscNode = client.GetDscNodeResource(dscNodeResourceId);

// invoke the operation
await dscNode.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
