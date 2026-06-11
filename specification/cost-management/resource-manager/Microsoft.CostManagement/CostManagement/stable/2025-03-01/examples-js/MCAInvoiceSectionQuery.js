const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to query the usage data for scope defined.
 *
 * @summary query the usage data for scope defined.
 * x-ms-original-file: 2025-03-01/MCAInvoiceSectionQuery.json
 */
async function invoiceSectionQueryMCA() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.query.usage(
    "providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579/invoiceSections/9876",
    {
      type: "Usage",
      dataset: {
        filter: {
          and: [
            {
              or: [
                {
                  dimensions: {
                    name: "ResourceLocation",
                    operator: "In",
                    values: ["East US", "West Europe"],
                  },
                },
                { tags: { name: "Environment", operator: "In", values: ["UAT", "Prod"] } },
              ],
            },
            { dimensions: { name: "ResourceGroup", operator: "In", values: ["API"] } },
          ],
        },
        granularity: "Daily",
      },
      timeframe: "MonthToDate",
    },
  );
  console.log(result);
}
