using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/examples/snapshotExamples/Snapshot_EndGetAccess.json
// this example is just showing the usage of "Snapshots_RevokeAccess" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SnapshotResource created on azure
// for more information of creating SnapshotResource, please refer to the document of SnapshotResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string snapshotName = "mySnapshot";
ResourceIdentifier snapshotResourceId = SnapshotResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, snapshotName);
SnapshotResource snapshot = client.GetSnapshotResource(snapshotResourceId);

// invoke the operation
await snapshot.RevokeAccessAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
