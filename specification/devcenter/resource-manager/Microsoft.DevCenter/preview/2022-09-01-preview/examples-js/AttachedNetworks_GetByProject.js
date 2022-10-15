const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an attached NetworkConnection.
 *
 * @summary Gets an attached NetworkConnection.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/AttachedNetworks_GetByProject.json
 */
async function attachedNetworksGetByProject() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const projectName = "{projectName}";
  const attachedNetworkConnectionName = "network-uswest3";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.attachedNetworks.getByProject(
    resourceGroupName,
    projectName,
    attachedNetworkConnectionName
  );
  console.log(result);
}

attachedNetworksGetByProject().catch(console.error);
