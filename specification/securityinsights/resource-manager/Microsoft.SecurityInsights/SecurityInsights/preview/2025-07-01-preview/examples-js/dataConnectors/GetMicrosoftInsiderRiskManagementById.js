const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a data connector.
 *
 * @summary gets a data connector.
 * x-ms-original-file: 2025-07-01-preview/dataConnectors/GetMicrosoftInsiderRiskManagementById.json
 */
async function getAnOfficeIRMDataConnector() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.dataConnectors.get(
    "myRg",
    "myWorkspace",
    "3d3e955e-33eb-401d-89a7-251c81ddd660",
  );
  console.log(result);
}
