using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automation.Models;
using Azure.ResourceManager.Automation;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getDscNodeReport.json
// this example is just showing the usage of "NodeReports_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DscNodeResource created on azure
// for more information of creating DscNodeResource, please refer to the document of DscNodeResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "myAutomationAccount33";
string nodeId = "nodeId";
ResourceIdentifier dscNodeResourceId = DscNodeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, nodeId);
DscNodeResource dscNode = client.GetDscNodeResource(dscNodeResourceId);

// invoke the operation
string reportId = "903a5ead-140c-11e7-a943-000d3a6140c9";
DscNodeReport result = await dscNode.GetNodeReportAsync(reportId);

Console.WriteLine($"Succeeded: {result}");
