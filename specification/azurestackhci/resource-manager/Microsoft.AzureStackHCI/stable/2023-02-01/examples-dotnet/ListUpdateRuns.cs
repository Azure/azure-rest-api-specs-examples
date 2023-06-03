using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Hci;
using Azure.ResourceManager.Hci.Models;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2023-02-01/examples/ListUpdateRuns.json
// this example is just showing the usage of "UpdateRuns_List" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
await foreach (UpdateRunResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    UpdateRunData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
