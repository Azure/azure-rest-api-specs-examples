using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/deleteHybridRunbookWorker.json
// this example is just showing the usage of "HybridRunbookWorkers_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridRunbookWorkerResource created on azure
// for more information of creating HybridRunbookWorkerResource, please refer to the document of HybridRunbookWorkerResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "myAutomationAccount20";
string hybridRunbookWorkerGroupName = "myGroup";
string hybridRunbookWorkerId = "c010ad12-ef14-4a2a-aa9e-ef22c4745ddd";
ResourceIdentifier hybridRunbookWorkerResourceId = HybridRunbookWorkerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, hybridRunbookWorkerId);
HybridRunbookWorkerResource hybridRunbookWorker = client.GetHybridRunbookWorkerResource(hybridRunbookWorkerResourceId);

// invoke the operation
await hybridRunbookWorker.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
