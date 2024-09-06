const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the policies for a billing account of Enterprise Agreement type.
 *
 * @summary Update the policies for a billing account of Enterprise Agreement type.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/policiesPutByBillingAccount.json
 */
async function policiesPutByBillingAccount() {
  const billingAccountName = "1234567";
  const parameters = {
    properties: {
      enterpriseAgreementPolicies: {
        authenticationType: "OrganizationalAccountOnly",
      },
      marketplacePurchases: "AllAllowed",
      reservationPurchases: "Allowed",
      savingsPlanPurchases: "NotAllowed",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.policies.beginCreateOrUpdateByBillingAccountAndWait(
    billingAccountName,
    parameters,
  );
  console.log(result);
}
