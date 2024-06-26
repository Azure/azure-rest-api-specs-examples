const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the instructions by billing profile id.
 *
 * @summary Lists the instructions by billing profile id.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InstructionsListByBillingProfile.json
 */
async function instructionsListByBillingProfile() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.instructions.listByBillingProfile(
    billingAccountName,
    billingProfileName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

instructionsListByBillingProfile().catch(console.error);
