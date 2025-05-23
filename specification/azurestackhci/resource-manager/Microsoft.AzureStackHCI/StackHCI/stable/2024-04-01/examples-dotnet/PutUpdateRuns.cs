using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Hci.Models;
using Azure.ResourceManager.Hci;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/PutUpdateRuns.json
// this example is just showing the usage of "UpdateRuns_Put" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HciClusterUpdateResource created on azure
// for more information of creating HciClusterUpdateResource, please refer to the document of HciClusterUpdateResource
string subscriptionId = "b8d594e5-51f3-4c11-9c54-a7771b81c712";
string resourceGroupName = "testrg";
string clusterName = "testcluster";
string updateName = "Microsoft4.2203.2.32";
ResourceIdentifier hciClusterUpdateResourceId = HciClusterUpdateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, updateName);
HciClusterUpdateResource hciClusterUpdate = client.GetHciClusterUpdateResource(hciClusterUpdateResourceId);

// get the collection of this HciClusterUpdateRunResource
HciClusterUpdateRunCollection collection = hciClusterUpdate.GetHciClusterUpdateRuns();

// invoke the operation
string updateRunName = "23b779ba-0d52-4a80-8571-45ca74664ec3";
HciClusterUpdateRunData data = new HciClusterUpdateRunData
{
    NamePropertiesProgressName = "Unnamed step",
    Description = "Update Azure Stack.",
    ErrorMessage = "",
    Status = "Success",
    StartOn = DateTimeOffset.Parse("2022-04-06T01:36:33.3876751+00:00"),
    EndOn = DateTimeOffset.Parse("2022-04-06T13:58:42.969006+00:00"),
    LastCompletedOn = DateTimeOffset.Parse("2022-04-06T13:58:42.969006+00:00"),
    Steps = {new HciUpdateStep
    {
    Name = "PreUpdate Cloud",
    Description = "Prepare for SSU update",
    ErrorMessage = "",
    Status = "Success",
    StartOn = DateTimeOffset.Parse("2022-04-06T01:36:33.3876751+00:00"),
    EndOn = DateTimeOffset.Parse("2022-04-06T01:37:16.8728314+00:00"),
    LastUpdatedOn = DateTimeOffset.Parse("2022-04-06T01:37:16.8728314+00:00"),
    Steps = {},
    }},
};
ArmOperation<HciClusterUpdateRunResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, updateRunName, data);
HciClusterUpdateRunResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HciClusterUpdateRunData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
