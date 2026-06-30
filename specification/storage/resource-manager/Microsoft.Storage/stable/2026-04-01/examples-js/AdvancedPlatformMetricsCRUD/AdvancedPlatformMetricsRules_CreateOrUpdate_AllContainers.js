const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update the advanced platform metrics rule for the storage account.
 *
 * @summary create or update the advanced platform metrics rule for the storage account.
 * x-ms-original-file: 2026-04-01/AdvancedPlatformMetricsCRUD/AdvancedPlatformMetricsRules_CreateOrUpdate_AllContainers.json
 */
async function advancedPlatformMetricsRulesCreateOrUpdateAllContainersCreateAdvancedPlatformMetricsRuleWithAllContainersFilter() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.advancedPlatformMetrics.createOrUpdate(
    "res6977",
    "sto2527",
    "ContainerLevelCapacityMetrics",
    { properties: { enabled: true, ruleConfig: { filterType: "AllContainersFilter" } } },
  );
  console.log(result);
}
