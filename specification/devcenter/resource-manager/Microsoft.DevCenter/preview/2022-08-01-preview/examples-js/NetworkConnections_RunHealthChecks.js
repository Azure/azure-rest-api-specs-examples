const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Triggers a new health check run. The execution and health check result can be tracked via the network Connection health check details
 *
 * @summary Triggers a new health check run. The execution and health check result can be tracked via the network Connection health check details
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/NetworkConnections_RunHealthChecks.json
 */
async function networkConnectionsRunHealthChecks() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const networkConnectionName = "uswest3network";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.networkConnections.runHealthChecks(
    resourceGroupName,
    networkConnectionName
  );
  console.log(result);
}

networkConnectionsRunHealthChecks().catch(console.error);
