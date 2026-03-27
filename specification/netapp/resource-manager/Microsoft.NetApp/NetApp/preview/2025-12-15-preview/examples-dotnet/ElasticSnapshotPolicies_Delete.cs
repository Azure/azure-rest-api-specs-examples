using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/preview/2025-12-15-preview/examples/ElasticSnapshotPolicies_Delete.json
// this example is just showing the usage of "ElasticSnapshotPolicies_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppElasticSnapshotPolicyResource created on azure
// for more information of creating NetAppElasticSnapshotPolicyResource, please refer to the document of NetAppElasticSnapshotPolicyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resourceGroup";
string accountName = "accountName";
string snapshotPolicyName = "snapshotPolicyName";
ResourceIdentifier netAppElasticSnapshotPolicyResourceId = NetAppElasticSnapshotPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, snapshotPolicyName);
NetAppElasticSnapshotPolicyResource netAppElasticSnapshotPolicy = client.GetNetAppElasticSnapshotPolicyResource(netAppElasticSnapshotPolicyResourceId);

// invoke the operation
await netAppElasticSnapshotPolicy.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
