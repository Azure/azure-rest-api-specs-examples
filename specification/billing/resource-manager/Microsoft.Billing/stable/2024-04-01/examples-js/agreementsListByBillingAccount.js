const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the agreements for a billing account.
 *
 * @summary Lists the agreements for a billing account.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/agreementsListByBillingAccount.json
 */
async function agreementsListByBillingAccount() {
  const billingAccountName =
    "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const expand = "Participants";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.agreements.listByBillingAccount(billingAccountName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
