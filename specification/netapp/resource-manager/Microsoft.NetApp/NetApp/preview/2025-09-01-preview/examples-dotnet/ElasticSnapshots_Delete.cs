using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/preview/2025-09-01-preview/examples/ElasticSnapshots_Delete.json
// this example is just showing the usage of "ElasticSnapshots_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticSnapshotResource created on azure
// for more information of creating ElasticSnapshotResource, please refer to the document of ElasticSnapshotResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
string poolName = "pool1";
string volumeName = "volume1";
string snapshotName = "snapshot1";
ResourceIdentifier elasticSnapshotResourceId = ElasticSnapshotResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName, volumeName, snapshotName);
ElasticSnapshotResource elasticSnapshot = client.GetElasticSnapshotResource(elasticSnapshotResourceId);

// invoke the operation
await elasticSnapshot.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
