const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Timeline for an entity.
 *
 * @summary Timeline for an entity.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/timeline/PostTimelineEntity.json
 */
async function entityTimeline() {
  const subscriptionId =
    process.env["SECURITYINSIGHT_SUBSCRIPTION_ID"] || "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = process.env["SECURITYINSIGHT_RESOURCE_GROUP"] || "myRg";
  const workspaceName = "myWorkspace";
  const entityId = "e1d3d618-e11f-478b-98e3-bb381539a8e1";
  const parameters = {
    endTime: new Date("2021-10-01T00:00:00.000Z"),
    numberOfBucket: 4,
    startTime: new Date("2021-09-01T00:00:00.000Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.entitiesGetTimeline.list(
    resourceGroupName,
    workspaceName,
    entityId,
    parameters
  );
  console.log(result);
}
