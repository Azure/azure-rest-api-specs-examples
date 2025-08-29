using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-06-01/examples/SnapshotPolicies_ListVolumes.json
// this example is just showing the usage of "SnapshotPolicies_ListVolumes" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SnapshotPolicyResource created on azure
// for more information of creating SnapshotPolicyResource, please refer to the document of SnapshotPolicyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
string snapshotPolicyName = "snapshotPolicyName";
ResourceIdentifier snapshotPolicyResourceId = SnapshotPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, snapshotPolicyName);
SnapshotPolicyResource snapshotPolicy = client.GetSnapshotPolicyResource(snapshotPolicyResourceId);

// invoke the operation and iterate over the result
await foreach (NetAppVolumeResource item in snapshotPolicy.GetVolumesAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetAppVolumeData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
