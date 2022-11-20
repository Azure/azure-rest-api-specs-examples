const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Validates if a subscription's charges can be moved to a new invoice section. This operation is supported for billing accounts with agreement type Microsoft Customer Agreement.
 *
 * @summary Validates if a subscription's charges can be moved to a new invoice section. This operation is supported for billing accounts with agreement type Microsoft Customer Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ValidateSubscriptionMoveSuccess.json
 */
async function subscriptionMoveValidateSuccess() {
  const subscriptionId = "{subscriptionId}";
  const billingAccountName = "{billingAccountName}";
  const parameters = {
    destinationInvoiceSectionId:
      "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{newInvoiceSectionName}",
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.billingSubscriptions.validateMove(billingAccountName, parameters);
  console.log(result);
}

subscriptionMoveValidateSuccess().catch(console.error);
