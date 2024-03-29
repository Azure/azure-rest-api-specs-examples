const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all watchlist Items.
 *
 * @summary Gets all watchlist Items.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/watchlists/GetWatchlistItems.json
 */
async function getAllWatchlistItems() {
  const subscriptionId =
    process.env["SECURITYINSIGHT_SUBSCRIPTION_ID"] || "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = process.env["SECURITYINSIGHT_RESOURCE_GROUP"] || "myRg";
  const workspaceName = "myWorkspace";
  const watchlistAlias = "highValueAsset";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.watchlistItems.list(
    resourceGroupName,
    workspaceName,
    watchlistAlias
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
