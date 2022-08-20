const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Network Connections resource
 *
 * @summary Deletes a Network Connections resource
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/NetworkConnections_Delete.json
 */
async function networkConnectionsDelete() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const networkConnectionName = "{networkConnectionName}";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.networkConnections.beginDeleteAndWait(
    resourceGroupName,
    networkConnectionName
  );
  console.log(result);
}

networkConnectionsDelete().catch(console.error);
