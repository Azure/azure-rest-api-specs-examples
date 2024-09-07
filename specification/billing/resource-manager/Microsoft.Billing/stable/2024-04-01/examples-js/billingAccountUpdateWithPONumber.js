const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the properties of a billing account. Currently, displayName and address can be updated for billing accounts with agreement type Microsoft Customer Agreement. Currently address and notification email address can be updated for billing accounts with agreement type Microsoft Online Services Agreement. Currently, purchase order number can be edited for billing accounts with agreement type Enterprise Agreement.
 *
 * @summary Updates the properties of a billing account. Currently, displayName and address can be updated for billing accounts with agreement type Microsoft Customer Agreement. Currently address and notification email address can be updated for billing accounts with agreement type Microsoft Online Services Agreement. Currently, purchase order number can be edited for billing accounts with agreement type Enterprise Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingAccountUpdateWithPONumber.json
 */
async function billingAccountUpdateWithPoNumber() {
  const billingAccountName = "6575495";
  const parameters = {
    properties: { enrollmentDetails: { poNumber: "poNumber123" } },
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingAccounts.beginUpdateAndWait(billingAccountName, parameters);
  console.log(result);
}
