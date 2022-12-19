const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets health check status details.
 *
 * @summary Gets health check status details.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/NetworkConnections_GetHealthDetails.json
 */
async function networkConnectionsGetHealthDetails() {
  const subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = "rg1";
  const networkConnectionName = "eastusnetwork";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.networkConnections.getHealthDetails(
    resourceGroupName,
    networkConnectionName
  );
  console.log(result);
}

networkConnectionsGetHealthDetails().catch(console.error);
