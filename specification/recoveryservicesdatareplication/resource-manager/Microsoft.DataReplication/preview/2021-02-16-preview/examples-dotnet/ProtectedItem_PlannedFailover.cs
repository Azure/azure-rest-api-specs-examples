using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesDataReplication;
using Azure.ResourceManager.RecoveryServicesDataReplication.Models;

// Generated from example definition: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/ProtectedItem_PlannedFailover.json
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
PlannedFailoverModel body = new PlannedFailoverModel(new PlannedFailoverModelProperties(new GeneralPlannedFailoverModelCustomProperties()));
ArmOperation<PlannedFailoverModel> lro = await dataReplicationProtectedItem.PlannedFailoverAsync(WaitUntil.Completed, body: body);
PlannedFailoverModel result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
