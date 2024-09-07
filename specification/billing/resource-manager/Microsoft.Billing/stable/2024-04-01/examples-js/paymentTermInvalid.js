const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Validates payment terms on a billing account with agreement type 'Microsoft Customer Agreement' and account type 'Enterprise'.
 *
 * @summary Validates payment terms on a billing account with agreement type 'Microsoft Customer Agreement' and account type 'Enterprise'.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentTermInvalid.json
 */
async function paymentTermInvalid() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const parameters = [
    {
      endDate: new Date("2023-01-25T22:39:34.2606750Z"),
      startDate: new Date("2023-02-05T22:39:34.2606750Z"),
      term: "net10",
    },
  ];
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingAccounts.validatePaymentTerms(billingAccountName, parameters);
  console.log(result);
}
