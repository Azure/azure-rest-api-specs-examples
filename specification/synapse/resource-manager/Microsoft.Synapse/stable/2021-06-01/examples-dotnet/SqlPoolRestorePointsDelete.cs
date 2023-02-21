using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolRestorePointsDelete.json
// this example is just showing the usage of "SqlPoolRestorePoints_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseRestorePointResource created on azure
// for more information of creating SynapseRestorePointResource, please refer to the document of SynapseRestorePointResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default-SQL-SouthEastAsia";
string workspaceName = "testws";
string sqlPoolName = "testpool";
string restorePointName = "131546477590000000";
ResourceIdentifier synapseRestorePointResourceId = SynapseRestorePointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, sqlPoolName, restorePointName);
SynapseRestorePointResource synapseRestorePoint = client.GetSynapseRestorePointResource(synapseRestorePointResourceId);

// invoke the operation
await synapseRestorePoint.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
