using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesDataReplication.Models;
using Azure.ResourceManager.RecoveryServicesDataReplication;

// Generated from example definition: 2024-09-01/ProtectedItem_Update.json
// this example is just showing the usage of "ProtectedItemModel_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataReplicationProtectedItemResource created on azure
// for more information of creating DataReplicationProtectedItemResource, please refer to the document of DataReplicationProtectedItemResource
string subscriptionId = "930CEC23-4430-4513-B855-DBA237E2F3BF";
string resourceGroupName = "rgswagger_2024-09-01";
string vaultName = "4";
string protectedItemName = "d";
ResourceIdentifier dataReplicationProtectedItemResourceId = DataReplicationProtectedItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, protectedItemName);
DataReplicationProtectedItemResource dataReplicationProtectedItem = client.GetDataReplicationProtectedItemResource(dataReplicationProtectedItemResourceId);

// invoke the operation
DataReplicationProtectedItemPatch patch = new DataReplicationProtectedItemPatch();
ArmOperation<DataReplicationProtectedItemResource> lro = await dataReplicationProtectedItem.UpdateAsync(WaitUntil.Completed, patch);
DataReplicationProtectedItemResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataReplicationProtectedItemData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
