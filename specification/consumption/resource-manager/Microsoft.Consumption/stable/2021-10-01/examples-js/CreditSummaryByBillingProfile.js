const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The credit summary by billingAccountId and billingProfileId.
 *
 * @summary The credit summary by billingAccountId and billingProfileId.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/CreditSummaryByBillingProfile.json
 */
async function creditSummaryByBillingProfile() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountId = "1234:5678";
  const billingProfileId = "2468";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const result = await client.credits.get(billingAccountId, billingProfileId);
  console.log(result);
}

creditSummaryByBillingProfile().catch(console.error);
