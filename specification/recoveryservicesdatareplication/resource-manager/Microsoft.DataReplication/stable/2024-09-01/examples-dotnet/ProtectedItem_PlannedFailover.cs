using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesDataReplication.Models;
using Azure.ResourceManager.RecoveryServicesDataReplication;

// Generated from example definition: 2024-09-01/ProtectedItem_PlannedFailover.json
// this example is just showing the usage of "ProtectedItem_PlannedFailover" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataReplicationProtectedItemResource created on azure
// for more information of creating DataReplicationProtectedItemResource, please refer to the document of DataReplicationProtectedItemResource
string subscriptionId = "930CEC23-4430-4513-B855-DBA237E2F3BF";
string resourceGroupName = "rgrecoveryservicesdatareplication";
string vaultName = "4";
string protectedItemName = "d";
ResourceIdentifier dataReplicationProtectedItemResourceId = DataReplicationProtectedItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, protectedItemName);
DataReplicationProtectedItemResource dataReplicationProtectedItem = client.GetDataReplicationProtectedItemResource(dataReplicationProtectedItemResourceId);

// invoke the operation
PlannedFailover body = new PlannedFailover(new PlannedFailoverProperties(null));
ArmOperation<PlannedFailover> lro = await dataReplicationProtectedItem.PlannedFailoverAsync(WaitUntil.Completed, body);
PlannedFailover result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
