const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to query the usage data for scope defined.
 *
 * @summary query the usage data for scope defined.
 * x-ms-original-file: 2025-03-01/ResourceGroupQueryGrouping.json
 */
async function resourceGroupQueryGroupingLegacy() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.query.usage(
    "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ScreenSharingTest-peer",
    {
      type: "Usage",
      dataset: {
        aggregation: { totalCost: { name: "PreTaxCost", function: "Sum" } },
        granularity: "Daily",
        grouping: [{ name: "ResourceType", type: "Dimension" }],
      },
      timeframe: "TheLastMonth",
    },
  );
  console.log(result);
}
