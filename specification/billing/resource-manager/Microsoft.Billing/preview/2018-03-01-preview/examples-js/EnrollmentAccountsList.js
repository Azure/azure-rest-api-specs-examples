const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the enrollment accounts the caller has access to.
 *
 * @summary Lists the enrollment accounts the caller has access to.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/preview/2018-03-01-preview/examples/EnrollmentAccountsList.json
 */
async function enrollmentAccountsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.enrollmentAccounts.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

enrollmentAccountsList().catch(console.error);
