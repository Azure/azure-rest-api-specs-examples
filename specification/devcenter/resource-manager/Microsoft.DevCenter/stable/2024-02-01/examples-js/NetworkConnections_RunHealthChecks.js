const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Triggers a new health check run. The execution and health check result can be tracked via the network Connection health check details
 *
 * @summary Triggers a new health check run. The execution and health check result can be tracked via the network Connection health check details
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/NetworkConnections_RunHealthChecks.json
 */
async function networkConnectionsRunHealthChecks() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = process.env["DEVCENTER_RESOURCE_GROUP"] || "rg1";
  const networkConnectionName = "uswest3network";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.networkConnections.beginRunHealthChecksAndWait(
    resourceGroupName,
    networkConnectionName,
  );
  console.log(result);
}
