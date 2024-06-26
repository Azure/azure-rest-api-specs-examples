const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a watchlist item.
 *
 * @summary Delete a watchlist item.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/watchlists/DeleteWatchlistItem.json
 */
async function deleteAWatchlistItem() {
  const subscriptionId =
    process.env["SECURITYINSIGHT_SUBSCRIPTION_ID"] || "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = process.env["SECURITYINSIGHT_RESOURCE_GROUP"] || "myRg";
  const workspaceName = "myWorkspace";
  const watchlistAlias = "highValueAsset";
  const watchlistItemId = "4008512e-1d30-48b2-9ee2-d3612ed9d3ea";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.watchlistItems.delete(
    resourceGroupName,
    workspaceName,
    watchlistAlias,
    watchlistItemId
  );
  console.log(result);
}
