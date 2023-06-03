using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Hci;
using Azure.ResourceManager.Hci.Models;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2023-02-01/examples/DeleteUpdateSummaries.json
// this example is just showing the usage of "UpdateSummaries_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this UpdateSummaryResource created on azure
// for more information of creating UpdateSummaryResource, please refer to the document of UpdateSummaryResource
string subscriptionId = "b8d594e5-51f3-4c11-9c54-a7771b81c712";
string resourceGroupName = "testrg";
string clusterName = "testcluster";
ResourceIdentifier updateSummaryResourceId = UpdateSummaryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
UpdateSummaryResource updateSummary = client.GetUpdateSummaryResource(updateSummaryResourceId);

// invoke the operation
await updateSummary.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
