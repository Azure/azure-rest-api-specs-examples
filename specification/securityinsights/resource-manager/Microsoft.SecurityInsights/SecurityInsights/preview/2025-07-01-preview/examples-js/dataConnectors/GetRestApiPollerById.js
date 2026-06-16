const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a data connector.
 *
 * @summary gets a data connector.
 * x-ms-original-file: 2025-07-01-preview/dataConnectors/GetRestApiPollerById.json
 */
async function getARestApiPollerDataConnector() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.dataConnectors.get(
    "myRg",
    "myWorkspace",
    "RestApiPoller_fce27b90-d6f5-4d30-991a-af509a2b50a1",
  );
  console.log(result);
}
