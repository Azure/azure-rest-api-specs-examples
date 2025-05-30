const { KeyVaultManagementClient } = require("@azure/arm-keyvault");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates the specified private endpoint connection associated with the managed hsm pool.
 *
 * @summary Updates the specified private endpoint connection associated with the managed hsm pool.
 * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2024-11-01/examples/ManagedHsm_putPrivateEndpointConnection.json
 */
async function managedHsmPutPrivateEndpointConnection() {
  const subscriptionId =
    process.env["KEYVAULT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["KEYVAULT_RESOURCE_GROUP"] || "sample-group";
  const name = "sample-mhsm";
  const privateEndpointConnectionName = "sample-pec";
  const properties = {
    privateLinkServiceConnectionState: {
      description: "My name is Joe and I'm approving this.",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new KeyVaultManagementClient(credential, subscriptionId);
  const result = await client.mhsmPrivateEndpointConnections.put(
    resourceGroupName,
    name,
    privateEndpointConnectionName,
    properties,
  );
  console.log(result);
}
