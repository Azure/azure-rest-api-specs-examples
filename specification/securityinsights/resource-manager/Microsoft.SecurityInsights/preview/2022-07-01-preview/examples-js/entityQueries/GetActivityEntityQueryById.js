const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an entity query.
 *
 * @summary Gets an entity query.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/entityQueries/GetActivityEntityQueryById.json
 */
async function getAnActivityEntityQuery() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const entityQueryId = "07da3cc8-c8ad-4710-a44e-334cdcb7882b";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.entityQueries.get(resourceGroupName, workspaceName, entityQueryId);
  console.log(result);
}

getAnActivityEntityQuery().catch(console.error);
