const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a watchlist, without its watchlist items.
 *
 * @summary Gets a watchlist, without its watchlist items.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/watchlists/GetWatchlistByAlias.json
 */
async function getAWatchlist() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const watchlistAlias = "highValueAsset";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.watchlists.get(resourceGroupName, workspaceName, watchlistAlias);
  console.log(result);
}

getAWatchlist().catch(console.error);
