const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the billing property of a subscription. Currently, cost center can be updated for billing accounts with agreement type Microsoft Customer Agreement and subscription service usage address can be updated for billing accounts with agreement type Microsoft Online Service Program.
 *
 * @summary Updates the billing property of a subscription. Currently, cost center can be updated for billing accounts with agreement type Microsoft Customer Agreement and subscription service usage address can be updated for billing accounts with agreement type Microsoft Online Service Program.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingPropertyPatchSubscriptionServiceUsageAddress.json
 */
async function billingPropertyPatchSubscriptionServiceUsageAddress() {
  const subscriptionId =
    process.env["BILLING_SUBSCRIPTION_ID"] || "11111111-1111-1111-1111-111111111111";
  const parameters = {
    properties: {
      subscriptionServiceUsageAddress: {
        addressLine1: "Address line 1",
        addressLine2: "Address line 2",
        city: "City",
        country: "US",
        firstName: "Jenny",
        lastName: "Doe",
        middleName: "Ann",
        postalCode: "12345",
        region: "State",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.billingPropertyOperations.update(parameters);
  console.log(result);
}
