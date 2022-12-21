const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all Savings plan orders.
 *
 * @summary List all Savings plan orders.
 * x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanOrderList.json
 */
async function savingsPlanOrderList() {
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const resArray = new Array();
  for await (let item of client.savingsPlanOrder.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

savingsPlanOrderList().catch(console.error);
