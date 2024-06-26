const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List savings plans in an order.
 *
 * @summary List savings plans in an order.
 * x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlansInOrderList.json
 */
async function savingsPlansInOrderList() {
  const savingsPlanOrderId = "20000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const resArray = new Array();
  for await (let item of client.savingsPlan.list(savingsPlanOrderId)) {
    resArray.push(item);
  }
  console.log(resArray);
}

savingsPlansInOrderList().catch(console.error);
