const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Query the usage data for scope defined.
 *
 * @summary Query the usage data for scope defined.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/EnrollmentAccountQueryGrouping.json
 */
async function enrollmentAccountQueryGroupingLegacy() {
  const scope = "providers/Microsoft.Billing/billingAccounts/100/enrollmentAccounts/456";
  const parameters = {
    type: "Usage",
    dataset: {
      aggregation: { totalCost: { name: "PreTaxCost", function: "Sum" } },
      granularity: "Daily",
      grouping: [{ name: "ResourceGroup", type: "Dimension" }],
    },
    timeframe: "TheLastMonth",
  };
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.query.usage(scope, parameters);
  console.log(result);
}
