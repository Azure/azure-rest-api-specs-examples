using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/restorePointExamples/RestorePoint_Get.json
// this example is just showing the usage of "RestorePoints_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RestorePointGroupResource created on azure
// for more information of creating RestorePointGroupResource, please refer to the document of RestorePointGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string restorePointGroupName = "rpcName";
ResourceIdentifier restorePointGroupResourceId = RestorePointGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, restorePointGroupName);
RestorePointGroupResource restorePointGroup = client.GetRestorePointGroupResource(restorePointGroupResourceId);

// get the collection of this RestorePointResource
RestorePointCollection collection = restorePointGroup.GetRestorePoints();

// invoke the operation
string restorePointName = "rpName";
NullableResponse<RestorePointResource> response = await collection.GetIfExistsAsync(restorePointName);
RestorePointResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    RestorePointData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
