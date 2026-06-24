const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a Kusto cluster.
 *
 * @summary update a Kusto cluster.
 * x-ms-original-file: 2025-02-14/KustoClusterUpdateEnableCMK.json
 */
async function kustoClusterUpdateEnableCMK() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusters.update(
    "kustoRgTest",
    "kustoCluster",
    {
      location: "westus",
      identity: {
        type: "UserAssigned",
        userAssignedIdentities: {
          "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustoRgTest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/kustoClusterIdentity":
            {},
        },
      },
      keyVaultProperties: {
        keyVaultUri: "https://myvault.vault.azure.net",
        keyName: "myClusterCMKKey",
        keyVersion: "12345678-1234-1234-1234-123456789098",
        userIdentity:
          "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustoRgTest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/kustoClusterIdentity",
      },
    },
    { ifMatch: "*" },
  );
  console.log(result);
}
