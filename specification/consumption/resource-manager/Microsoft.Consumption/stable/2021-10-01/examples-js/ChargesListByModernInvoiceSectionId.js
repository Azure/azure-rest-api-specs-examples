const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the charges based for the defined scope.
 *
 * @summary Lists the charges based for the defined scope.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ChargesListByModernInvoiceSectionId.json
 */
async function chargesListByInvoiceSectionIdModern() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "providers/Microsoft.Billing/BillingAccounts/1234:56789/invoiceSections/97531";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const result = await client.charges.list(scope);
  console.log(result);
}

chargesListByInvoiceSectionIdModern().catch(console.error);
