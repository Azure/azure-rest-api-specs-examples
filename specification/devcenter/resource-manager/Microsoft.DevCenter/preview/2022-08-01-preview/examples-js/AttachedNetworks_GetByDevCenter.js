const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an attached NetworkConnection.
 *
 * @summary Gets an attached NetworkConnection.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/AttachedNetworks_GetByDevCenter.json
 */
async function attachedNetworksGetByDevCenter() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const attachedNetworkConnectionName = "network-uswest3";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.attachedNetworks.getByDevCenter(
    resourceGroupName,
    devCenterName,
    attachedNetworkConnectionName
  );
  console.log(result);
}

attachedNetworksGetByDevCenter().catch(console.error);
