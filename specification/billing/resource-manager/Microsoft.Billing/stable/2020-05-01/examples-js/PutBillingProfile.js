const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a billing profile. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement or Microsoft Partner Agreement.
 *
 * @summary Creates or updates a billing profile. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement or Microsoft Partner Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/PutBillingProfile.json
 */
async function createBillingProfile() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const parameters = {
    billTo: {
      addressLine1: "Test Address 1",
      city: "Redmond",
      country: "US",
      firstName: "Test",
      lastName: "User",
      postalCode: "12345",
      region: "WA",
    },
    displayName: "Finance",
    enabledAzurePlans: [{ skuId: "0001" }, { skuId: "0002" }],
    invoiceEmailOptIn: true,
    poNumber: "ABC12345",
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.billingProfiles.beginCreateOrUpdateAndWait(
    billingAccountName,
    billingProfileName,
    parameters
  );
  console.log(result);
}

createBillingProfile().catch(console.error);
