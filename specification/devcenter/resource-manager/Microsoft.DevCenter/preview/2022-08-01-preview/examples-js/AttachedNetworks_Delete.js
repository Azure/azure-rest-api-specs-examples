const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Un-attach a NetworkConnection.
 *
 * @summary Un-attach a NetworkConnection.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/AttachedNetworks_Delete.json
 */
async function attachedNetworksDelete() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const attachedNetworkConnectionName = "{attachedNetworkConnectionName}";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.attachedNetworks.beginDeleteAndWait(
    resourceGroupName,
    devCenterName,
    attachedNetworkConnectionName
  );
  console.log(result);
}

attachedNetworksDelete().catch(console.error);
