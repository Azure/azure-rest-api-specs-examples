using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-03-02/examples/snapshotExamples/Snapshot_Update_WithAcceleratedNetwork.json
// this example is just showing the usage of "Snapshots_Update" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this SnapshotResource created on azure
// for more information of creating SnapshotResource, please refer to the document of SnapshotResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string snapshotName = "mySnapshot";
ResourceIdentifier snapshotResourceId = SnapshotResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, snapshotName);
SnapshotResource snapshot = client.GetSnapshotResource(snapshotResourceId);

// invoke the operation
SnapshotPatch patch = new SnapshotPatch()
{
    Tags =
    {
    ["department"] = "Development",
    ["project"] = "UpdateSnapshots",
    },
    DiskSizeGB = 20,
    SupportedCapabilities = new SupportedCapabilities()
    {
        AcceleratedNetwork = false,
    },
};
ArmOperation<SnapshotResource> lro = await snapshot.UpdateAsync(WaitUntil.Completed, patch);
SnapshotResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SnapshotData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
