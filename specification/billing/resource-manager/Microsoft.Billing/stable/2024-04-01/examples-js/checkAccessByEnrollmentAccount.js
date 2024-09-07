const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides a list of check access response objects for an enrollment account.
 *
 * @summary Provides a list of check access response objects for an enrollment account.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/checkAccessByEnrollmentAccount.json
 */
async function checkAccessByEnrollmentAccount() {
  const billingAccountName =
    "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const enrollmentAccountName = "123456";
  const parameters = {
    actions: [
      "Microsoft.Billing/billingAccounts/read",
      "Microsoft.Subscription/subscriptions/write",
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingPermissions.checkAccessByEnrollmentAccount(
    billingAccountName,
    enrollmentAccountName,
    parameters,
  );
  console.log(result);
}
