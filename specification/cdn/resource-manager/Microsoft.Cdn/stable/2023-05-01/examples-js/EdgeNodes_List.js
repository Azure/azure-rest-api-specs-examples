const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Edgenodes are the global Point of Presence (POP) locations used to deliver CDN content to end users.
 *
 * @summary Edgenodes are the global Point of Presence (POP) locations used to deliver CDN content to end users.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/EdgeNodes_List.json
 */
async function edgeNodesList() {
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.edgeNodes.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
