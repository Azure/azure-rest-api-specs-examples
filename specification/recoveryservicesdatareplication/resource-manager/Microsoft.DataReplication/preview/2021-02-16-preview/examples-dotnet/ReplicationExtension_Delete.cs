using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesDataReplication.Models;
using Azure.ResourceManager.RecoveryServicesDataReplication;

// Generated from example definition: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/ReplicationExtension_Delete.json
// this example is just showing the usage of "ReplicationExtension_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataReplicationReplicationExtensionResource created on azure
// for more information of creating DataReplicationReplicationExtensionResource, please refer to the document of DataReplicationReplicationExtensionResource
string subscriptionId = "930CEC23-4430-4513-B855-DBA237E2F3BF";
string resourceGroupName = "rgrecoveryservicesdatareplication";
string vaultName = "4";
string replicationExtensionName = "g16yjJ";
ResourceIdentifier dataReplicationReplicationExtensionResourceId = DataReplicationReplicationExtensionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, replicationExtensionName);
DataReplicationReplicationExtensionResource dataReplicationReplicationExtension = client.GetDataReplicationReplicationExtensionResource(dataReplicationReplicationExtensionResourceId);

// invoke the operation
await dataReplicationReplicationExtension.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
