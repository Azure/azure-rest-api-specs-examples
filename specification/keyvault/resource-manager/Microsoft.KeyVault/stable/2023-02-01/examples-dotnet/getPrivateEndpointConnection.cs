using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.KeyVault;
using Azure.ResourceManager.KeyVault.Models;

// Generated from example definition: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-02-01/examples/getPrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

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
bool result = await collection.ExistsAsync(privateEndpointConnectionName);

Console.WriteLine($"Succeeded: {result}");
