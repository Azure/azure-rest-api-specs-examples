using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesDataReplication.Models;
using Azure.ResourceManager.RecoveryServicesDataReplication;

// Generated from example definition: 2024-09-01/PrivateEndpointConnection_Update.json
// this example is just showing the usage of "PrivateEndpointConnection_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataReplicationVaultResource created on azure
// for more information of creating DataReplicationVaultResource, please refer to the document of DataReplicationVaultResource
string subscriptionId = "930CEC23-4430-4513-B855-DBA237E2F3BF";
string resourceGroupName = "rgswagger_2024-09-01";
string vaultName = "4";
ResourceIdentifier dataReplicationVaultResourceId = DataReplicationVaultResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName);
DataReplicationVaultResource dataReplicationVault = client.GetDataReplicationVaultResource(dataReplicationVaultResourceId);

// get the collection of this DataReplicationPrivateEndpointConnectionResource
DataReplicationPrivateEndpointConnectionCollection collection = dataReplicationVault.GetDataReplicationPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "jitf";
DataReplicationPrivateEndpointConnectionData data = new DataReplicationPrivateEndpointConnectionData
{
    Properties = new DataReplicationPrivateEndpointConnectionProperties
    {
        PrivateEndpointId = new ResourceIdentifier("cwcdqoynostmqwdwy"),
        PrivateLinkServiceConnectionState = new DataReplicationPrivateLinkServiceConnectionState
        {
            Status = DataReplicationPrivateEndpointConnectionStatus.Approved,
            Description = "y",
            ActionsRequired = "afwbq",
        },
    },
};
ArmOperation<DataReplicationPrivateEndpointConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, privateEndpointConnectionName, data);
DataReplicationPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataReplicationPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
