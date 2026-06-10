const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to query the usage data for scope defined.
 *
 * @summary query the usage data for scope defined.
 * x-ms-original-file: 2025-03-01/MCAInvoiceSectionQueryGrouping.json
 */
async function invoiceSectionQueryGroupingMCA() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.query.usage(
    "providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579/invoiceSections/9876",
    {
      type: "Usage",
      dataset: {
        aggregation: { totalCost: { name: "PreTaxCost", function: "Sum" } },
        granularity: "None",
        grouping: [{ name: "ResourceGroup", type: "Dimension" }],
      },
      timeframe: "TheLastMonth",
    },
  );
  console.log(result);
}
