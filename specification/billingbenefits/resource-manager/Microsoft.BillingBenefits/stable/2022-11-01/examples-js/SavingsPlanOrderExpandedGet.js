const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a savings plan order.
 *
 * @summary Get a savings plan order.
 * x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanOrderExpandedGet.json
 */
async function savingsPlanOrderWithExpandedPaymentsGet() {
  const savingsPlanOrderId = "20000000-0000-0000-0000-000000000000";
  const expand = "schedule";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.savingsPlanOrder.get(savingsPlanOrderId, options);
  console.log(result);
}

savingsPlanOrderWithExpandedPaymentsGet().catch(console.error);
