const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a data connector.
 *
 * @summary Gets a data connector.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/dataConnectors/GetThreatIntelligenceTaxiiById.json
 */
async function getATiTaxiiDataConnector() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const dataConnectorId = "c39bb458-02a7-4b3f-b0c8-71a1d2692652";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.dataConnectors.get(resourceGroupName, workspaceName, dataConnectorId);
  console.log(result);
}

getATiTaxiiDataConnector().catch(console.error);
