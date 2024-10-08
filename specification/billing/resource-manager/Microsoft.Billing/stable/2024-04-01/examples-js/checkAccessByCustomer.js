const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides a list of check access response objects for a customer.
 *
 * @summary Provides a list of check access response objects for a customer.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/checkAccessByCustomer.json
 */
async function checkAccessByCustomer() {
  const billingAccountName =
    "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const billingProfileName = "xxxx-xxxx-xxx-xxx";
  const customerName = "703ab484-dda2-4402-827b-a74513b61e2d";
  const parameters = {
    actions: [
      "Microsoft.Billing/billingAccounts/read",
      "Microsoft.Subscription/subscriptions/write",
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingPermissions.checkAccessByCustomer(
    billingAccountName,
    billingProfileName,
    customerName,
    parameters,
  );
  console.log(result);
}
