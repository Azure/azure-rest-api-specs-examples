using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesDataReplication.Models;
using Azure.ResourceManager.RecoveryServicesDataReplication;

// Generated from example definition: 2024-09-01/ReplicationExtension_Delete.json
// this example is just showing the usage of "ReplicationExtensionModel_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataReplicationExtensionResource created on azure
// for more information of creating DataReplicationExtensionResource, please refer to the document of DataReplicationExtensionResource
string subscriptionId = "930CEC23-4430-4513-B855-DBA237E2F3BF";
string resourceGroupName = "rgrecoveryservicesdatareplication";
string vaultName = "4";
string replicationExtensionName = "g16yjJ";
ResourceIdentifier dataReplicationExtensionResourceId = DataReplicationExtensionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, replicationExtensionName);
DataReplicationExtensionResource dataReplicationExtension = client.GetDataReplicationExtensionResource(dataReplicationExtensionResourceId);

// invoke the operation
await dataReplicationExtension.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
