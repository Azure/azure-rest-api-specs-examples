const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists the Alerts for external cloud provider type defined.
 *
 * @summary lists the Alerts for external cloud provider type defined.
 * x-ms-original-file: 2025-03-01/ExternalBillingAccountAlerts.json
 */
async function externalBillingAccountAlerts() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.alerts.listExternal("externalBillingAccounts", "100");
  console.log(result);
}
