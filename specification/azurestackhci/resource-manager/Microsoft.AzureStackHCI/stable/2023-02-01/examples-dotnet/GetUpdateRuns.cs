using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Hci;
using Azure.ResourceManager.Hci.Models;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2023-02-01/examples/GetUpdateRuns.json
// this example is just showing the usage of "UpdateRuns_Get" operation, for the dependent resources, they will have to be created separately.

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
NullableResponse<UpdateRunResource> response = await collection.GetIfExistsAsync(updateRunName);
UpdateRunResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    UpdateRunData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
