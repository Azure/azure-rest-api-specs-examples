using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/restorePointExamples/RestorePointCollection_Update_MinimumSet_Gen.json
// this example is just showing the usage of "RestorePointCollections_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RestorePointGroupResource created on azure
// for more information of creating RestorePointGroupResource, please refer to the document of RestorePointGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string restorePointGroupName = "aaaaaaaaaaaaaaaaaa";
ResourceIdentifier restorePointGroupResourceId = RestorePointGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, restorePointGroupName);
RestorePointGroupResource restorePointGroup = client.GetRestorePointGroupResource(restorePointGroupResourceId);

// invoke the operation
RestorePointGroupPatch patch = new RestorePointGroupPatch();
RestorePointGroupResource result = await restorePointGroup.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RestorePointGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
