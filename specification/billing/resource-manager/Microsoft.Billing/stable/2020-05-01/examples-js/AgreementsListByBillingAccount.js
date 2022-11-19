const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the agreements for a billing account.
 *
 * @summary Lists the agreements for a billing account.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/AgreementsListByBillingAccount.json
 */
async function agreementsListByBillingAccount() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.agreements.listByBillingAccount(billingAccountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

agreementsListByBillingAccount().catch(console.error);
