const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Execute Insights for an entity.
 *
 * @summary Execute Insights for an entity.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/entities/insights/PostGetInsights.json
 */
async function entityInsight() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const entityId = "e1d3d618-e11f-478b-98e3-bb381539a8e1";
  const parameters = {
    addDefaultExtendedTimeRange: false,
    endTime: new Date("2021-10-01T00:00:00.000Z"),
    insightQueryIds: ["cae8d0aa-aa45-4d53-8d88-17dd64ffd4e4"],
    startTime: new Date("2021-09-01T00:00:00.000Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.entities.getInsights(
    resourceGroupName,
    workspaceName,
    entityId,
    parameters
  );
  console.log(result);
}

entityInsight().catch(console.error);
