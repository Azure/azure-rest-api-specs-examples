const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an entity.
 *
 * @summary Gets an entity.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/entities/GetMailboxEntityById.json
 */
async function getAMailboxEntity() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const entityId = "e1d3d618-e11f-478b-98e3-bb381539a8e1";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.entities.get(resourceGroupName, workspaceName, entityId);
  console.log(result);
}

getAMailboxEntity().catch(console.error);
