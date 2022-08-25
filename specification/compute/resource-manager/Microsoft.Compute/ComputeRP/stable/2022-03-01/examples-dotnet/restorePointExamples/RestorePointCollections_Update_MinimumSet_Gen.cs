using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/restorePointExamples/RestorePointCollections_Update_MinimumSet_Gen.json
// this example is just showing the usage of "RestorePointCollections_Update" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
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
