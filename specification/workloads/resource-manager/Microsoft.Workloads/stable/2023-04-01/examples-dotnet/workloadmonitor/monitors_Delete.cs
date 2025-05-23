using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Workloads;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/monitors_Delete.json
// this example is just showing the usage of "monitors_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SapMonitorResource created on azure
// for more information of creating SapMonitorResource, please refer to the document of SapMonitorResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "mySapMonitor";
ResourceIdentifier sapMonitorResourceId = SapMonitorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName);
SapMonitorResource sapMonitor = client.GetSapMonitorResource(sapMonitorResourceId);

// invoke the operation
ArmOperation<OperationStatusResult> lro = await sapMonitor.DeleteAsync(WaitUntil.Completed);
OperationStatusResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
