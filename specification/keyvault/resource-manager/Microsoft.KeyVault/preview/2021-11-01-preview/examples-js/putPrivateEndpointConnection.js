const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the specified private endpoint connection associated with the key vault.
 *
 * @summary Updates the specified private endpoint connection associated with the key vault.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/putPrivateEndpointConnection.json
 */
async function keyVaultPutPrivateEndpointConnection() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "sample-group";
  const vaultName = "sample-vault";
  const privateEndpointConnectionName = "sample-pec";
  const properties = {
    etag: "",
    privateLinkServiceConnectionState: {
      description: "My name is Joe and I'm approving this.",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.put(
    resourceGroupName,
    vaultName,
    privateEndpointConnectionName,
    properties
  );
  console.log(result);
}

keyVaultPutPrivateEndpointConnection().catch(console.error);
