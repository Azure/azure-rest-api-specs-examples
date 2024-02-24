using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/restorePointExamples/RestorePoint_Copy_BetweenRegions.json
// this example is just showing the usage of "RestorePoints_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RestorePointResource created on azure
// for more information of creating RestorePointResource, please refer to the document of RestorePointResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string restorePointGroupName = "rpcName";
string restorePointName = "rpName";
ResourceIdentifier restorePointResourceId = RestorePointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, restorePointGroupName, restorePointName);
RestorePointResource restorePoint = client.GetRestorePointResource(restorePointResourceId);

// invoke the operation
RestorePointData data = new RestorePointData()
{
    SourceRestorePointId = new ResourceIdentifier("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/sourceRpcName/restorePoints/sourceRpName"),
};
ArmOperation<RestorePointResource> lro = await restorePoint.UpdateAsync(WaitUntil.Completed, data);
RestorePointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RestorePointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
