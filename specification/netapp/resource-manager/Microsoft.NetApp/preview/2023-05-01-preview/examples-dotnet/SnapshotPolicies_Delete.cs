using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetApp;
using Azure.ResourceManager.NetApp.Models;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/preview/2023-05-01-preview/examples/SnapshotPolicies_Delete.json
// this example is just showing the usage of "SnapshotPolicies_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SnapshotPolicyResource created on azure
// for more information of creating SnapshotPolicyResource, please refer to the document of SnapshotPolicyResource
string subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
string resourceGroupName = "resourceGroup";
string accountName = "accountName";
string snapshotPolicyName = "snapshotPolicyName";
ResourceIdentifier snapshotPolicyResourceId = SnapshotPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, snapshotPolicyName);
SnapshotPolicyResource snapshotPolicy = client.GetSnapshotPolicyResource(snapshotPolicyResourceId);

// invoke the operation
await snapshotPolicy.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
