using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.KeyVault.Models;
using Azure.ResourceManager.KeyVault;

// Generated from example definition: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/putPrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Put" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KeyVaultResource created on azure
// for more information of creating KeyVaultResource, please refer to the document of KeyVaultResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "sample-group";
string vaultName = "sample-vault";
ResourceIdentifier keyVaultResourceId = KeyVaultResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName);
KeyVaultResource keyVault = client.GetKeyVaultResource(keyVaultResourceId);

// get the collection of this KeyVaultPrivateEndpointConnectionResource
KeyVaultPrivateEndpointConnectionCollection collection = keyVault.GetKeyVaultPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "sample-pec";
KeyVaultPrivateEndpointConnectionData data = new KeyVaultPrivateEndpointConnectionData
{
    ETag = new ETag(""),
    ConnectionState = new KeyVaultPrivateLinkServiceConnectionState
    {
        Status = KeyVaultPrivateEndpointServiceConnectionStatus.Approved,
        Description = "My name is Joe and I'm approving this.",
    },
};
ArmOperation<KeyVaultPrivateEndpointConnectionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, privateEndpointConnectionName, data);
KeyVaultPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KeyVaultPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
