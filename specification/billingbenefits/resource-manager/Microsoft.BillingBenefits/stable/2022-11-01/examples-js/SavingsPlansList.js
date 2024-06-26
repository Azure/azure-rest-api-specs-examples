const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List savings plans.
 *
 * @summary List savings plans.
 * x-ms-original-file: specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlansList.json
 */
async function savingsPlansList() {
  const filter = "(properties%2farchived+eq+false)";
  const orderby = "properties/displayName asc";
  const skiptoken = 50;
  const take = 1;
  const options = {
    filter,
    orderby,
    skiptoken,
    take,
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const resArray = new Array();
  for await (let item of client.savingsPlan.listAll(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

savingsPlansList().catch(console.error);
