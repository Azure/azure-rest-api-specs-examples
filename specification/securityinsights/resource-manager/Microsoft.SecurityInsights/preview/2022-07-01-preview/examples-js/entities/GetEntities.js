const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all entities.
 *
 * @summary Gets all entities.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/entities/GetEntities.json
 */
async function getAllEntities() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.entities.list(resourceGroupName, workspaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAllEntities().catch(console.error);
