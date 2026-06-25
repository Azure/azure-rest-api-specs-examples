const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a Kusto cluster.
 *
 * @summary update a Kusto cluster.
 * x-ms-original-file: 2025-02-14/KustoClusterUpdateCMKKeyRotation.json
 */
async function kustoClusterUpdateCMKKeyRotation() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusters.update(
    "kustoRgTest",
    "kustoClusterCMK",
    {
      location: "westus",
      keyVaultProperties: {
        keyVaultUri: "https://myvault.vault.azure.net",
        keyName: "myClusterCMKKey",
        keyVersion: "87654321-4321-4321-4321-210987654321",
        userIdentity:
          "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustoRgTest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/kustoClusterIdentity",
      },
    },
    { ifMatch: "*" },
  );
  console.log(result);
}
