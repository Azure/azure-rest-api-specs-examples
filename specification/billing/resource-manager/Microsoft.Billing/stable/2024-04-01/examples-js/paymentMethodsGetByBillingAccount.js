const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a payment method available for a billing account. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
 *
 * @summary Gets a payment method available for a billing account. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsGetByBillingAccount.json
 */
async function paymentMethodGetAtBillingProfile() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000032:00000000-0000-0000-0000-000000000099_2019-05-31";
  const paymentMethodName = "21dd9edc-af71-4d62-80ce-37151d475326";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.paymentMethods.getByBillingAccount(
    billingAccountName,
    paymentMethodName,
  );
  console.log(result);
}
