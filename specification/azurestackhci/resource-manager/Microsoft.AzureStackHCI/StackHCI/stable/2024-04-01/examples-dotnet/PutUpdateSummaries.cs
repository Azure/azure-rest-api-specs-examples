using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Hci.Models;
using Azure.ResourceManager.Hci;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/PutUpdateSummaries.json
// this example is just showing the usage of "UpdateSummaries_Put" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HciClusterUpdateSummaryResource created on azure
// for more information of creating HciClusterUpdateSummaryResource, please refer to the document of HciClusterUpdateSummaryResource
string subscriptionId = "b8d594e5-51f3-4c11-9c54-a7771b81c712";
string resourceGroupName = "testrg";
string clusterName = "testcluster";
ResourceIdentifier hciClusterUpdateSummaryResourceId = HciClusterUpdateSummaryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
HciClusterUpdateSummaryResource hciClusterUpdateSummary = client.GetHciClusterUpdateSummaryResource(hciClusterUpdateSummaryResourceId);

// invoke the operation
HciClusterUpdateSummaryData data = new HciClusterUpdateSummaryData()
{
    OemFamily = "DellEMC",
    HardwareModel = "PowerEdge R730xd",
    CurrentVersion = "4.2203.2.32",
    LastUpdatedOn = DateTimeOffset.Parse("2022-04-06T14:08:18.254Z"),
    LastCheckedOn = DateTimeOffset.Parse("2022-04-07T18:04:07Z"),
    State = HciClusterUpdateState.AppliedSuccessfully,
};
ArmOperation<HciClusterUpdateSummaryResource> lro = await hciClusterUpdateSummary.CreateOrUpdateAsync(WaitUntil.Completed, data);
HciClusterUpdateSummaryResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HciClusterUpdateSummaryData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
