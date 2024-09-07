const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides a list of check access response objects for a billing account.
 *
 * @summary Provides a list of check access response objects for a billing account.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/checkAccessByBillingAccount.json
 */
async function checkAccessByBillingAccount() {
  const billingAccountName =
    "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const parameters = {
    actions: [
      "Microsoft.Billing/billingAccounts/read",
      "Microsoft.Subscription/subscriptions/write",
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingPermissions.checkAccessByBillingAccount(
    billingAccountName,
    parameters,
  );
  console.log(result);
}
