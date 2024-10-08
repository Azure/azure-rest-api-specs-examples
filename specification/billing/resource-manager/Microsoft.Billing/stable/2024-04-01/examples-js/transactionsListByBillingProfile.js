const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the billed or unbilled transactions by billing profile name for given start and end date. Transactions include purchases, refunds and Azure usage charges. Unbilled transactions are listed under pending invoice Id and do not include tax. Tax is added to the amount once an invoice is generated.
 *
 * @summary Lists the billed or unbilled transactions by billing profile name for given start and end date. Transactions include purchases, refunds and Azure usage charges. Unbilled transactions are listed under pending invoice Id and do not include tax. Tax is added to the amount once an invoice is generated.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/transactionsListByBillingProfile.json
 */
async function transactionsListByBillingProfile() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const billingProfileName = "xxxx-xxxx-xxx-xxx";
  const periodStartDate = new Date("2024-04-01");
  const periodEndDate = new Date("2023-05-30");
  const typeParam = "Billed";
  const filter = "properties/date gt '2020-10-01'";
  const search = "storage";
  const options = {
    filter,
    search,
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.transactions.listByBillingProfile(
    billingAccountName,
    billingProfileName,
    periodStartDate,
    periodEndDate,
    typeParam,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
