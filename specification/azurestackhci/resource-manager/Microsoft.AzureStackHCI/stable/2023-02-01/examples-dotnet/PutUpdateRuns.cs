using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Hci;
using Azure.ResourceManager.Hci.Models;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2023-02-01/examples/PutUpdateRuns.json
// this example is just showing the usage of "UpdateRuns_Put" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this UpdateResource created on azure
// for more information of creating UpdateResource, please refer to the document of UpdateResource
string subscriptionId = "b8d594e5-51f3-4c11-9c54-a7771b81c712";
string resourceGroupName = "testrg";
string clusterName = "testcluster";
string updateName = "Microsoft4.2203.2.32";
ResourceIdentifier updateResourceId = UpdateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, updateName);
UpdateResource update = client.GetUpdateResource(updateResourceId);

// get the collection of this UpdateRunResource
UpdateRunCollection collection = update.GetUpdateRuns();

// invoke the operation
string updateRunName = "23b779ba-0d52-4a80-8571-45ca74664ec3";
UpdateRunData data = new UpdateRunData()
{
    NamePropertiesProgressName = "Unnamed step",
    Description = "Update Azure Stack.",
    ErrorMessage = "",
    Status = "Success",
    StartTimeUtc = DateTimeOffset.Parse("2022-04-06T01:36:33.3876751+00:00"),
    EndTimeUtc = DateTimeOffset.Parse("2022-04-06T13:58:42.969006+00:00"),
    LastUpdatedTimeUtc = DateTimeOffset.Parse("2022-04-06T13:58:42.969006+00:00"),
    Steps =
    {
    new HciUpdateStep()
    {
    Name = "PreUpdate Cloud",
    Description = "Prepare for SSU update",
    ErrorMessage = "",
    Status = "Success",
    StartTimeUtc = DateTimeOffset.Parse("2022-04-06T01:36:33.3876751+00:00"),
    EndTimeUtc = DateTimeOffset.Parse("2022-04-06T01:37:16.8728314+00:00"),
    LastUpdatedTimeUtc = DateTimeOffset.Parse("2022-04-06T01:37:16.8728314+00:00"),
    Steps =
    {
    },
    }
    },
};
ArmOperation<UpdateRunResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, updateRunName, data);
UpdateRunResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
UpdateRunData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
