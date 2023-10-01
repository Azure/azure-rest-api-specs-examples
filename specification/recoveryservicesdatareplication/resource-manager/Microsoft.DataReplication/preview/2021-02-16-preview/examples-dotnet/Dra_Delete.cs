using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesDataReplication;
using Azure.ResourceManager.RecoveryServicesDataReplication.Models;

// Generated from example definition: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Dra_Delete.json
// this example is just showing the usage of "Dra_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataReplicationDraResource created on azure
// for more information of creating DataReplicationDraResource, please refer to the document of DataReplicationDraResource
string subscriptionId = "930CEC23-4430-4513-B855-DBA237E2F3BF";
string resourceGroupName = "rgrecoveryservicesdatareplication";
string fabricName = "wPR";
string fabricAgentName = "M";
ResourceIdentifier dataReplicationDraResourceId = DataReplicationDraResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, fabricName, fabricAgentName);
DataReplicationDraResource dataReplicationDra = client.GetDataReplicationDraResource(dataReplicationDraResourceId);

// invoke the operation
await dataReplicationDra.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
