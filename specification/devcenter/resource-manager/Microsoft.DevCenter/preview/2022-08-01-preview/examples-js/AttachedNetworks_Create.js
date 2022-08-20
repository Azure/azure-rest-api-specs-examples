const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an attached NetworkConnection.
 *
 * @summary Creates or updates an attached NetworkConnection.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/AttachedNetworks_Create.json
 */
async function attachedNetworksCreate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const attachedNetworkConnectionName = "{attachedNetworkConnectionName}";
  const body = {
    networkConnectionId:
      "/subscriptions/{subscriptionId}/resourceGroups/rg1/providers/Microsoft.DevCenter/NetworkConnections/network-uswest3",
  };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.attachedNetworks.beginCreateOrUpdateAndWait(
    resourceGroupName,
    devCenterName,
    attachedNetworkConnectionName,
    body
  );
  console.log(result);
}

attachedNetworksCreate().catch(console.error);
