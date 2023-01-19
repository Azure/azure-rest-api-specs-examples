using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automation;
using Azure.ResourceManager.Automation.Models;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/getHybridRunbookWorker.json
// this example is just showing the usage of "HybridRunbookWorkers_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridRunbookWorkerGroupResource created on azure
// for more information of creating HybridRunbookWorkerGroupResource, please refer to the document of HybridRunbookWorkerGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "testaccount";
string hybridRunbookWorkerGroupName = "TestHybridGroup";
ResourceIdentifier hybridRunbookWorkerGroupResourceId = HybridRunbookWorkerGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName);
HybridRunbookWorkerGroupResource hybridRunbookWorkerGroup = client.GetHybridRunbookWorkerGroupResource(hybridRunbookWorkerGroupResourceId);

// get the collection of this HybridRunbookWorkerResource
HybridRunbookWorkerCollection collection = hybridRunbookWorkerGroup.GetHybridRunbookWorkers();

// invoke the operation
string hybridRunbookWorkerId = "c010ad12-ef14-4a2a-aa9e-ef22c4745ddd";
bool result = await collection.ExistsAsync(hybridRunbookWorkerId);

Console.WriteLine($"Succeeded: {result}");
