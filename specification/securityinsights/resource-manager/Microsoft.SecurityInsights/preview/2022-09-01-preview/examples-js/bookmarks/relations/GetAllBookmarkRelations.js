const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all bookmark relations.
 *
 * @summary Gets all bookmark relations.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/bookmarks/relations/GetAllBookmarkRelations.json
 */
async function getAllBookmarkRelations() {
  const subscriptionId =
    process.env["SECURITYINSIGHT_SUBSCRIPTION_ID"] || "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = process.env["SECURITYINSIGHT_RESOURCE_GROUP"] || "myRg";
  const workspaceName = "myWorkspace";
  const bookmarkId = "2216d0e1-91e3-4902-89fd-d2df8c535096";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.bookmarkRelations.list(
    resourceGroupName,
    workspaceName,
    bookmarkId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
