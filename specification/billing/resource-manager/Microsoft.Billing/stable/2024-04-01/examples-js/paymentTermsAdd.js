const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Adds payment terms to all the billing profiles under the billing account. Currently, payment terms can be added only on billing accounts that have Agreement Type as 'Microsoft Customer Agreement' and AccountType as 'Enterprise'. This action needs pre-authorization and only Field Sellers are authorized to add the payment terms and is not a self-serve action.
 *
 * @summary Adds payment terms to all the billing profiles under the billing account. Currently, payment terms can be added only on billing accounts that have Agreement Type as 'Microsoft Customer Agreement' and AccountType as 'Enterprise'. This action needs pre-authorization and only Field Sellers are authorized to add the payment terms and is not a self-serve action.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentTermsAdd.json
 */
async function paymentTermsAdd() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const parameters = [
    {
      endDate: new Date("2023-01-25T22:39:34.2606750Z"),
      startDate: new Date("2023-01-05T22:39:34.2606750Z"),
      term: "net10",
    },
  ];
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingAccounts.beginAddPaymentTermsAndWait(
    billingAccountName,
    parameters,
  );
  console.log(result);
}
