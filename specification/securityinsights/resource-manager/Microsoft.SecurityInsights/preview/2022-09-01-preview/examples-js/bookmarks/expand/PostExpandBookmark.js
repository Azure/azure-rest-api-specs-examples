const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Expand an bookmark
 *
 * @summary Expand an bookmark
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/bookmarks/expand/PostExpandBookmark.json
 */
async function expandAnBookmark() {
  const subscriptionId =
    process.env["SECURITYINSIGHT_SUBSCRIPTION_ID"] || "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = process.env["SECURITYINSIGHT_RESOURCE_GROUP"] || "myRg";
  const workspaceName = "myWorkspace";
  const bookmarkId = "73e01a99-5cd7-4139-a149-9f2736ff2ab5";
  const parameters = {
    endTime: new Date("2020-01-24T17:21:00.000Z"),
    expansionId: "27f76e63-c41b-480f-bb18-12ad2e011d49",
    startTime: new Date("2019-12-25T17:21:00.000Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.bookmarkOperations.expand(
    resourceGroupName,
    workspaceName,
    bookmarkId,
    parameters
  );
  console.log(result);
}
