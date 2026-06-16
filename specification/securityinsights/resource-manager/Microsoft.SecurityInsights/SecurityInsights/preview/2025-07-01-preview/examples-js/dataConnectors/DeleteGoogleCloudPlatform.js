const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete the data connector.
 *
 * @summary delete the data connector.
 * x-ms-original-file: 2025-07-01-preview/dataConnectors/DeleteGoogleCloudPlatform.json
 */
async function deleteAGCPDataConnector() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  await client.dataConnectors.delete(
    "myRg",
    "myWorkspace",
    "GCP_fce27b90-d6f5-4d30-991a-af509a2b50a1",
  );
}
