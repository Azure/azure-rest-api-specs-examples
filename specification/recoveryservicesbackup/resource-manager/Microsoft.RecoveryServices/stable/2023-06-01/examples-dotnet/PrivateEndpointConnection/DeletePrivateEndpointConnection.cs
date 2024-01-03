using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/PrivateEndpointConnection/DeletePrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnection_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackupPrivateEndpointConnectionResource created on azure
// for more information of creating BackupPrivateEndpointConnectionResource, please refer to the document of BackupPrivateEndpointConnectionResource
string subscriptionId = "04cf684a-d41f-4550-9f70-7708a3a2283b";
string resourceGroupName = "gaallaRG";
string vaultName = "gaallavaultbvtd2msi";
string privateEndpointConnectionName = "gaallatestpe2.5704c932-249a-490b-a142-1396838cd3b";
ResourceIdentifier backupPrivateEndpointConnectionResourceId = BackupPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, privateEndpointConnectionName);
BackupPrivateEndpointConnectionResource backupPrivateEndpointConnection = client.GetBackupPrivateEndpointConnectionResource(backupPrivateEndpointConnectionResourceId);

// invoke the operation
await backupPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
