const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a Log Analytics cluster.
 *
 * @summary Updates a Log Analytics cluster.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersUpdate.json
 */
async function clustersPatch() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "oiautorest6685";
  const clusterName = "oiautorest6685";
  const parameters = {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/0000000000000000000000000000000/resourcegroups/oiautorest6685/providers/MicrosoftManagedIdentity/userAssignedIdentities/myidentity":
          {},
      },
    },
    keyVaultProperties: {
      keyName: "aztest2170cert",
      keyRsaSize: 1024,
      keyVaultUri: "https://aztest2170.vault.azure.net",
      keyVersion: "654ft6c4e63845cbb50fd6fg51540429",
    },
    sku: { name: "CapacityReservation", capacity: 1000 },
    tags: { tag1: "val1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginUpdateAndWait(
    resourceGroupName,
    clusterName,
    parameters
  );
  console.log(result);
}

clustersPatch().catch(console.error);
