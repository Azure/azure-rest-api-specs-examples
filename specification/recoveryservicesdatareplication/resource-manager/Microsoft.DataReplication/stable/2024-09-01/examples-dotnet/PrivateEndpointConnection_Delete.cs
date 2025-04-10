using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesDataReplication.Models;
using Azure.ResourceManager.RecoveryServicesDataReplication;

// Generated from example definition: 2024-09-01/PrivateEndpointConnection_Delete.json
// this example is just showing the usage of "PrivateEndpointConnection_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataReplicationPrivateEndpointConnectionResource created on azure
// for more information of creating DataReplicationPrivateEndpointConnectionResource, please refer to the document of DataReplicationPrivateEndpointConnectionResource
string subscriptionId = "930CEC23-4430-4513-B855-DBA237E2F3BF";
string resourceGroupName = "rgswagger_2024-09-01";
string vaultName = "4";
string privateEndpointConnectionName = "sdwqtfhigjirrzhpbmqtzgs";
ResourceIdentifier dataReplicationPrivateEndpointConnectionResourceId = DataReplicationPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, privateEndpointConnectionName);
DataReplicationPrivateEndpointConnectionResource dataReplicationPrivateEndpointConnection = client.GetDataReplicationPrivateEndpointConnectionResource(dataReplicationPrivateEndpointConnectionResourceId);

// invoke the operation
await dataReplicationPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
