const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the properties of a billing account. Currently, displayName and address can be updated. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
 *
 * @summary Updates the properties of a billing account. Currently, displayName and address can be updated. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdateBillingAccount.json
 */
async function updateBillingAccount() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountName = "{billingAccountName}";
  const parameters = {
    displayName: "Test Account",
    soldTo: {
      addressLine1: "Test Address 1",
      city: "Redmond",
      companyName: "Contoso",
      country: "US",
      firstName: "Test",
      lastName: "User",
      postalCode: "12345",
      region: "WA",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.billingAccounts.beginUpdateAndWait(billingAccountName, parameters);
  console.log(result);
}

updateBillingAccount().catch(console.error);
