const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the Alerts for external cloud provider type defined.
 *
 * @summary Lists the Alerts for external cloud provider type defined.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExternalSubscriptionAlerts.json
 */
async function externalSubscriptionAlerts() {
  const externalCloudProviderType = "externalSubscriptions";
  const externalCloudProviderId = "100";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.alerts.listExternal(
    externalCloudProviderType,
    externalCloudProviderId
  );
  console.log(result);
}
