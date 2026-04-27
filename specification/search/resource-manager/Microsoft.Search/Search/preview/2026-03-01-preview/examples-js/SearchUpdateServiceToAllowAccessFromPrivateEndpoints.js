const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing search service in the given resource group.
 *
 * @summary updates an existing search service in the given resource group.
 * x-ms-original-file: 2026-03-01-preview/SearchUpdateServiceToAllowAccessFromPrivateEndpoints.json
 */
async function searchUpdateServiceToAllowAccessFromPrivateEndpoints() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.services.update("rg1", "mysearchservice", {
    replicaCount: 1,
    partitionCount: 1,
    publicNetworkAccess: "Disabled",
  });
  console.log(result);
}
