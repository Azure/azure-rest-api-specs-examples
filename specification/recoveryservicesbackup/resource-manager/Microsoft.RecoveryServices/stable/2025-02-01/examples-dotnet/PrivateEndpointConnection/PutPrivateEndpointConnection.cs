using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesBackup.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.RecoveryServicesBackup;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/PrivateEndpointConnection/PutPrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnection_Put" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "04cf684a-d41f-4550-9f70-7708a3a2283b";
string resourceGroupName = "gaallaRG";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this BackupPrivateEndpointConnectionResource
BackupPrivateEndpointConnectionCollection collection = resourceGroupResource.GetBackupPrivateEndpointConnections();

// invoke the operation
string vaultName = "gaallavaultbvtd2msi";
string privateEndpointConnectionName = "gaallatestpe2.5704c932-249a-490b-a142-1396838cd3b";
BackupPrivateEndpointConnectionData data = new BackupPrivateEndpointConnectionData(default)
{
    Properties = new BackupPrivateEndpointConnectionProperties
    {
        ProvisioningState = BackupPrivateEndpointConnectionProvisioningState.Succeeded,
        PrivateEndpointId = new ResourceIdentifier("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/gaallaRG/providers/Microsoft.Network/privateEndpoints/gaallatestpe3"),
        GroupIds = { VaultSubResourceType.AzureBackupSecondary },
        PrivateLinkServiceConnectionState = new RecoveryServicesBackupPrivateLinkServiceConnectionState
        {
            Status = PrivateEndpointConnectionStatus.Approved,
            Description = "Approved by johndoe@company.com",
        },
    },
};
ArmOperation<BackupPrivateEndpointConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, vaultName, privateEndpointConnectionName, data);
BackupPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BackupPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
