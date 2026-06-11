const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists the alerts for scope defined.
 *
 * @summary lists the alerts for scope defined.
 * x-ms-original-file: 2025-03-01/InvoiceSectionAlerts.json
 */
async function invoiceSectionAlerts() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.alerts.list(
    "providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579/invoiceSections/9876",
  );
  console.log(result);
}
