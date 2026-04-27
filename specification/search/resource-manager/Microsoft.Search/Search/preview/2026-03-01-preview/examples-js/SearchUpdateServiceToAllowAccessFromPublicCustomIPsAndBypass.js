const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing search service in the given resource group.
 *
 * @summary updates an existing search service in the given resource group.
 * x-ms-original-file: 2026-03-01-preview/SearchUpdateServiceToAllowAccessFromPublicCustomIPsAndBypass.json
 */
async function searchUpdateServiceToAllowAccessFromPublicCustomIPsAndBypass() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.services.update("rg1", "mysearchservice", {
    replicaCount: 3,
    partitionCount: 1,
    publicNetworkAccess: "Enabled",
    networkRuleSet: {
      ipRules: [{ value: "123.4.5.6" }, { value: "123.4.6.0/18" }],
      bypass: "AzurePortal",
    },
  });
  console.log(result);
}
