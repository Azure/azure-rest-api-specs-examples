const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Validates if the invoice section can be deleted. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement.
 *
 * @summary Validates if the invoice section can be deleted. The operation is supported for billing accounts with agreement type Microsoft Customer Agreement.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/invoiceSectionsValidateDeleteEligibilityFailure.json
 */
async function invoiceSectionsValidateDeleteEligibilityFailure() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const billingProfileName = "xxxx-xxxx-xxx-xxx";
  const invoiceSectionName = "yyyy-yyyy-yyy-yyy";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.invoiceSections.validateDeleteEligibility(
    billingAccountName,
    billingProfileName,
    invoiceSectionName,
  );
  console.log(result);
}
