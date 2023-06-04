using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Hci;
using Azure.ResourceManager.Hci.Models;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2023-02-01/examples/GetUpdateSummaries.json
// this example is just showing the usage of "UpdateSummaries_Get" operation, for the dependent resources, they will have to be created separately.

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
UpdateSummaryResource result = await updateSummary.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
UpdateSummaryData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
