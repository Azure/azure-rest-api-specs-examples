const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Kusto cluster.
 *
 * @summary create or update a Kusto cluster.
 * x-ms-original-file: 2025-02-14/KustoClusterCreateWithCMKFederatedIdentity.json
 */
async function kustoClusterCreateWithCMKFederatedIdentity() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.clusters.createOrUpdate("kustoRgTest", "kustoClusterCMK", {
    location: "westus",
    sku: { name: "Standard_L16as_v3", capacity: 2, tier: "Standard" },
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustoRgTest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/kustoClusterIdentity":
          {},
      },
    },
    enableStreamingIngest: true,
    enablePurge: true,
    enableDoubleEncryption: false,
    enableAutoStop: true,
    publicIPType: "IPv4",
    publicNetworkAccess: "Enabled",
    restrictOutboundNetworkAccess: "Disabled",
    keyVaultProperties: {
      keyVaultUri: "https://myvault.vault.azure.net",
      keyName: "myClusterCMKKey",
      keyVersion: "12345678-1234-1234-1234-123456789098",
      userIdentity:
        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustoRgTest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/kustoClusterIdentity",
      federatedIdentityClientId: "11111111-2222-3333-4444-555555555555",
    },
    engineType: "V3",
  });
  console.log(result);
}
